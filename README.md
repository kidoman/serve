# serve

The simplest directory serving web server. 

The program is written in Go but having the Go compiler installed in not a requirement as the binaries available below are self sufficient.

## How to get it

A simple:

	go get github.com/kidoman/serve

is more than sufficient. Also, you can download precompiled versions to place in your ```/usr/local/bin``` from the below links:

* [Mac OS X 64 bit](https://dl.dropboxusercontent.com/u/6727135/Binaries/serve/darwin-amd64/serve)
* [Linux 64 bit](https://dl.dropboxusercontent.com/u/6727135/Binaries/serve/linux-amd64/serve)
* [Linux 32 bit](https://dl.dropboxusercontent.com/u/6727135/Binaries/serve/linux-386/serve)

## Platforms supported

Anything you can run a precompiled binary on (i.e. Windows, Mac OS X, Linux, etc.)

## How to use

Provided you have ```serve``` under your $PATH somewhere:

	serve .

This will serve the current directory at ```http://localhost:5000/```

	serve -p 9999 ~/my-awesome-blog

Will serve the contents of the folder ```~/my-awesome-blog``` at ```http://localhost:9999/```

	serve -x /my ~/precious

You guessed it, ```http://localhost:5000/my``` is now wired up to ```~/precious```

## Coming soon

* HTTPS support
