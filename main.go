package main

import (
	"encoding/base32"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"net/url"
	"os"
	"strings"

	"github.com/dnjp/gauthexport/proto/gauth"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"google.golang.org/protobuf/proto"
)

func usage(name string) {
	fmt.Fprintf(os.Stdout, "%s [QR FILE] [OUT DIR]\n", name)
}

func decodeQR(path string) (*gozxing.Result, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return nil, err
	}
	result, err := qrcode.NewQRCodeReader().Decode(bmp, nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func dataFromURL(migrateURL *url.URL) ([]byte, error) {
	return base64.StdEncoding.DecodeString(migrateURL.Query().Get("data"))
}

func newMigration(b []byte) (*gauth.Migration, error) {
	var mig gauth.Migration
	err := proto.Unmarshal(b, &mig)
	if err != nil {
		return nil, err
	}
	return &mig, nil
}

func migrationFromURL(rawURL string) (*gauth.Migration, error) {
	migrateURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	b, err := dataFromURL(migrateURL)
	if err != nil {
		return nil, err
	}

	mig, err := newMigration(b)
	if err != nil {
		return nil, err
	}

	return mig, nil
}

func writeQRCodes(dir string, mig *gauth.Migration) error {
	for _, param := range mig.OtpParameters {
		if param.Algorithm != gauth.Migration_ALGO_SHA1 {
			return fmt.Errorf("invalid algorithm: %s", param.Algorithm.String())
		}

		issuer, err := url.QueryUnescape(param.Issuer)
		if err != nil {
			return err
		}

		name := param.Name
		if strings.Contains(name, issuer+":") {
			name = strings.ReplaceAll(name, issuer+":", "")
		}

		q := "otpauth://totp/"
		q += url.QueryEscape(param.Issuer) + ":" + name + "?"
		qv := make(url.Values)
		qv.Add("secret", base32.StdEncoding.EncodeToString(param.Secret))
		qv.Add("issuer", issuer)
		qv.Add("algorithm", "SHA1")
		q += qv.Encode()

		qrw := qrcode.NewQRCodeWriter()
		img, err := qrw.Encode(q, gozxing.BarcodeFormat_QR_CODE, 250, 250, nil)
		if err != nil {
			return err
		}

		file, err := os.Create(fmt.Sprintf("%s/%s.png", dir, issuer))
		if err != nil {
			return err
		}
		defer file.Close()

		err = png.Encode(file, img)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "must provide QR filename and output directory\n")
		usage(os.Args[0])
		os.Exit(1)
	}

	qrImg := os.Args[len(os.Args)-2]
	outdir := os.Args[len(os.Args)-1]

	migrateURL, err := decodeQR(qrImg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}

	mig, err := migrationFromURL(migrateURL.String())
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}

	err = writeQRCodes(outdir, mig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}
