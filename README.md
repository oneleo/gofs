# gofs
 Gofs is a small file server for Browsing specify local folder and files.

## PreRequire
- [Git](https://git-scm.com/) - Distributed version control system.
- [Golang](https://golang.org/) - Programming language.

## Install
```golang
$> go get -u github.com/oneleo/gofs
$> go install github.com/oneleo/gofs
```

## Usage
- Start this file server listen on port ":80" for browse "./" current folder.
```bash
$> gofs
```

- Start this file server listen on port ":80" for browse "./" current folder, and open in browser.
```bash
$> gofs -b
```

- Start this file server listen on port ":8081" for browse "./www" specify folder, and open in browser.
```bash
$> gofs -p 8081 -f ./www -b
```

- Help message
```bash
$> gofs -h
OR
$> gofs --help
```

```text
$> gofs -h
Gofs is a small file server for Browsing specify local folder and files.
Usage:

        gofs [-p port] [-f folder] [-b]

Flags:
  -b    Open URL in browser (shorthand)
  -f string
        Set shared folder (shorthand) (default "./")
  -p int
        Set listening port (shorthand) (default 80)

Examples:
  gofs
        Start this file server listen on port ":80" for browse "./" current folder.
  gofs -b
        Start this file server listen on port ":80" for browse "./" current folder, and open in browser.
  gofs -p 8081 -f ./www -b
        Start this file server listen on specify port ":8081" for browse "./www" specify folder, and open in browser.
```

## References
- [GitHub - Simple static file server](https://github.com/golang-id/gost)
- [GitHub Gist - Open browser in golang](https://gist.github.com/hyg/9c4afcd91fe24316cbf0)
- [Stackoverflow - Start web server to open page in browser](https://stackoverflow.com/questions/39320371/how-start-web-server-to-open-page-in-browser-in-golang)