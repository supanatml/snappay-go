package main

import (
	"context"
	"fmt"
	"net"

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

func (a *App) shutdown(ctx context.Context) {

}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) QRpopup() {
	//todo: get QR from GB
	fmt.Println("QR popped")
}

func (a *App) VoucherPopup() {
	//todo: get QR from GB
	fmt.Println("Voucher popped")
}

func (a *App) TermsPopup() {
	//todo: get QR from GB
	fmt.Println("Terms popped")
}
