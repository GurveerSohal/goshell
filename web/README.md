# Shell on the web

Note: Whatever you do in that shell, will reflect on your actual machine.
If you don't want that, then you can build the docker image using the Dockerfile, but that is a work in progress.
You can do that using the following:
```
docker build -t go-web-shell .
docker run -it -p 8080:8080 go-web-shell
```
If you do this you can ignore the steps for the first terminal

You'll need to open 2 terminals. 

In the first one, do
```
go run .
```

In the second one, do
```
npm install
npm run dev
```

Then go to localhost://3000, and you should be able to use the shell as normal.
