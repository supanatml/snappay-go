package main

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"image/png"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/keygen-sh/keygen-go/v2"
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

func checkSingleInstanceRunning() bool {
	// binds TPC port 45456 to prevent multiple instances running concurrentl
	listener, err := net.Listen("tcp4", "0.0.0.0:45456")
	if err != nil {
		fmt.Println(err)
		//todo: add err checker
		fmt.Println("Cannot run multiple instances concurrently.")
		return false
	}
	fmt.Println(listener.Addr())
	fmt.Println("Port check passed")
	return true
	//todo: debug why http.Listen stops listening when QRpopup is called
}

func checkValidLicense() bool {
	//todo: keygen.sh
	keygen.Account = keygenID
	keygen.Product = keygenProductID
	keygen.LicenseKey = licenseKey

	out, err := exec.Command("powershell", "-command", "wmic csproduct get uuid").Output()
	if err != nil {
		fmt.Println(err)
	}

	uuid := strings.Split(string(out), "  \r\r\n")[1]

	sha256out := sha256.Sum256([]byte(uuid))
	fingerprint := hex.EncodeToString(sha256out[:])

	fmt.Printf("%q\n", uuid)
	fmt.Println(fingerprint)

	license, err := keygen.Validate(fingerprint)
	switch {
	case err == keygen.ErrLicenseNotActivated:
		//Activate the current fingerprint
		fmt.Println("License not activated; attempting to activate.")
		machine, err := license.Activate(fingerprint)
		switch {
		case err == keygen.ErrMachineLimitExceeded:
			fmt.Println("Machine limit exceeded; exiting.")
			return false
		case err != nil:
			fmt.Println("Machine activation failed; exiting.")
			return false
		default:
			fmt.Println("New license activated on this machine.")
			fmt.Println(machine)
			return true
		}
	case err != nil:
		fmt.Println("License is invalid.")
		return false
	default:
		fmt.Println("License is valid; continuing.")
		return true
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	//if !checkSingleInstanceRunning() || !checkValidLicense() {
	//	runtime.Quit(a.ctx)
	//}
}

// cleanup on shutdown
func (a *App) shutdown(ctx context.Context) {

}

func (a *App) QRpopup() string {
	fmt.Println("QR popped")

	//GBPrimePay URL and monetary amount to charge
	//gburl := "https://api.globalprimepay.com/v3/qrcode"
	gburl := "https://api.gbprimepay.com/v3/qrcode"
	amount := "169.00"

	client := &http.Client{
		Timeout: time.Second * 5,
	}

	referenceNo := strconv.FormatInt(time.Now().Unix(), 10)

	// url values formatting according to gbpay documentation
	body := url.Values{}
	body.Set("token", gbtoken)
	body.Set("amount", amount)
	body.Set("referenceNo", referenceNo)

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

	return referenceNo
}

func (a *App) CheckTransaction(ref string) bool {
	//todo: check transsaction with gbpay

	//GBPrimePay URL
	//gburl := "https://api.globalprimepay.com/v1/check_status_txn"
	gburl := "https://api.gbprimepay.com/v1/check_status_txn"
	authkey := "Basic " + base64.StdEncoding.EncodeToString([]byte(gbsecret+":"))

	fmt.Println(authkey)

	client := &http.Client{
		Timeout: time.Second * 5,
	}

	var jsonData = []byte(`{"referenceNo": ` + ref + `}`)

	req, err := http.NewRequest("POST", gburl, strings.NewReader(string(jsonData)))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authkey)
	req.Header.Add("Content-Length", strconv.Itoa(len(jsonData)))

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("response Body:", string(body))

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(body), &jsonMap)

	fmt.Println(jsonMap)

	//todo: check tx status
	txnMap := jsonMap["txn"].(map[string]interface{})
	status := txnMap["status"]

	fmt.Println(status)
	fmt.Println(reflect.TypeOf(status))

	//status code according to gbpay documentation "G = Generate , A = Authorize , S = Settle, V = Void, D = Decline"
	if status == "S" {
		//start session
		fmt.Println("yeahh")
		a.StartSession()
		return true
	} else {
		//todo: let gbpay know to not accept future payments
		return false
	}
}

func (a *App) VoucherPopup() {
	//todo: screen to enter voucher code for free sessions
	fmt.Println("Voucher popped")
}

// for debugging purposes
func (a *App) PrintLog(s string) {
	fmt.Println(s)
}

func (a *App) StartSession() {
	//todo: call dslrbooth's API
	dslrboothURL := "http://localhost:1500/api/start?mode=print" + "&password=" + dslrboothPassword

	response, err := http.Get(dslrboothURL)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("response Body:", string(body))

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(body), &jsonMap)

	fmt.Println(jsonMap)

	//todo: check tx status
	isSuccessful := jsonMap["IsSuccessful"].(bool)

	if isSuccessful {
		//success and quit
		runtime.Quit(a.ctx)
	}
}
