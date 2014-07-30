# serve

The simplest directory serving web server.

The program is written in Go but having the Go compiler installed in not a requirement as the binaries available below are self sufficient.

## How to get it

A simple:

	go get github.com/kidoman/serve

or:

	brew install kidoman/tools/serve

will get the job done. Also, you can download precompiled releases to place in your ```/usr/local/bin``` from:

https://github.com/kidoman/serve/releases/latest

PS: We should hopefully be in the official brew repo soon. Then a ```brew install serve``` will be sufficient.

## Platforms supported

Anything you can run a precompiled binary on (i.e. Windows, Mac OS X, Linux, etc.)

## How to use

Provided you have ```serve``` under your $PATH somewhere:

	serve

This will serve the current directory at ```http://localhost:5000/```

	serve -p 9999 ~/my-awesome-blog

Will serve the contents of the folder ```~/my-awesome-blog``` at ```http://localhost:9999/```

	serve -x /my ~/precious

You guessed it, ```http://localhost:5000/my``` is now wired up to ```~/precious```

	serve -o ~/sesame

Wires up ```http://localhost:5000``` to ```~/sesame``` and opens the URL in your favorite browser while it is at it.

## Coming soon

* HTTPS support
