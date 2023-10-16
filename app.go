package main

import (
	"context"
	"fmt"
	"image/png"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// binds TPC port 45456 to prevent multiple instances running concurrently
	listener, err := net.Listen("tcp4", "0.0.0.0:45456")
	if err != nil {
		fmt.Println(err)
		//todo: add err checker
		fmt.Println("Cannot run multiple instances concurrently.")
		runtime.Quit(a.ctx)
	}
	fmt.Println(listener.Addr())
	fmt.Println("Port check passed")

	//todo: keygen.sh
}

// cleanup on shutdown
func (a *App) shutdown(ctx context.Context) {

}

func (a *App) QRpopup() {
	fmt.Println("QR popped")

	//GBPrimePay URL and monetary amount to charge
	gburl := "https://api.globalprimepay.com/v3/qrcode"
	amount := "150.00"

	client := &http.Client{
		Timeout: time.Second * 1,
	}

	// url values formatting according to gbpay documentation
	body := url.Values{}
	body.Set("token", gbtoken)
	body.Set("amount", amount)
	body.Set("referenceNo", strconv.FormatInt(time.Now().Unix(), 10))

	encodedBody := body.Encode()

	req, err := http.NewRequest("POST", gburl, strings.NewReader(encodedBody))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(encodedBody)))

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	qrimage, err := png.Decode(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	// write received image/png to file
	f, err := os.Create("qr.png")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	if err = png.Encode(f, qrimage); err != nil {
		fmt.Println(err)
	}
}

func (a *App) VoucherPopup() {
	//todo: screen to enter voucher code for free sessions
	fmt.Println("Voucher popped")
}
