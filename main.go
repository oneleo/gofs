package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

var (
	port   *int64  = new(int64)
	folder *string = new(string)
	browse *bool   = new(bool)
)

func init() {
	const (
		defaultPort = 80
		usagePort   = "Set listening port"

		defaultFolder = "./"
		usageFolder   = "Set shared folder"

		defaultBrowse = false
		usageBrowse   = "Open URL in browser"
	)
	// flag.Int64Var(&port, "port", defaultPort, usagePort)
	flag.Int64Var(port, "p", defaultPort, usagePort+" (shorthand)")
	// flag.StringVar(&folder, "folder", defaultFolder, usageFolder)
	flag.StringVar(folder, "f", defaultFolder, usageFolder+" (shorthand)")
	// flag.BoolVar(browse, "browse", defaultBrowse, usageBrowse)
	flag.BoolVar(browse, "b", defaultBrowse, usageBrowse+" (shorthand)")

	flag.Usage = func() {
		// template: go help doc
		fmt.Fprintf(os.Stderr, "Gofs is a small file server for Browsing specify local folder and files.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\n\tgofs [-p port] [-f folder] [-b]\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  gofs\n\tStart this file server listen on default port \":80\" for browse \"./\" current folder.\n")
		fmt.Fprintf(os.Stderr, "  gofs -b\n\tStart this file server listen on default port \":80\" for browse \"./\" current folder, and \"auto-open\" in browser.\n")
		fmt.Fprintf(os.Stderr, "  gofs -p 8081 -f ./www -b\n\tStart this file server listen on specify port \":8081\" for browse \"./www\" specify folder, and \"auto-open\" in browser.")
	}
}

// openURI will open URI in browser.
// Reconmannd to use this new project: https://github.com/toqueteos/webbrowser
func openURI(uri string) {
	var c *string = new(string)
	var args *[]string = new([]string)

	// Obtain OS type.
	switch runtime.GOOS {
	case "windows":
		// *c = "cmd"
		// *args = []string{"/c", "start"}
		*c = "rundll32"
		*args = []string{"url.dll,FileProtocolHandler"}
	case "darwin":
		*c = "open"
	case "linux":
		*c = "xdg-open"
	default:
		log.Fatal("Unsupported operating system\n")
	}
	*args = append(*args, uri)
	// Open URI in browser.
	if err := exec.Command(*c, *args...).Start(); err != nil {
		log.Fatal(err, "\n")
	}
}

func main() {
	flag.Parse()
	// Obtain absolute folder even in different OS.
	f, err := filepath.Abs(*folder)
	if err != nil {
		log.Fatal(err, "\n")
	}

	h := http.FileServer(http.Dir(f))
	p := fmt.Sprintf(":%d", *port)
	fmt.Print("File Server is Started to listen on http://127.0.0.1", p, " for browse ", f, "\nPlease press CTRL + C to finish the file server...\n")
	// If '-b' flag is set, open URI in browser.
	if *browse == true {
		go openURI("http://127.0.0.1" + p)
	}
	// Start this file server
	if err = http.ListenAndServe(p, h); err != nil {
		log.Fatal(err, "\n")
	}
}
