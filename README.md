# serve

The simplest directory serving web server.

## How to get

A simple:

	go get github.com/kidoman/serve

is more than sufficient. I would prefer you used the ```brew``` way though.

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
