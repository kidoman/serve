# serve

The simplest directory serving web server.

## Platforms supported

Anything you can run a precompiled binary on. Or simply do a:

	go run main.go

if you would rather go that way.

## How to use

Provided you have ```serve``` under your path somewhere (it will soon be possible to just do a ```brew install serve```):

	serve .

This will serve the current directory at ```http://localhost:5000/```

	serve -p 9999 ~/my-awesome-blog

Will serve the contents of the folder ```~/my-awesome-blog``` at ```http://localhost:9999/```

	serve -pf /my ~/precious

You guessed it, ```http://localhost:5000/my``` is now wired up to ```~/precious```

## Coming soon

* HTTPS support
