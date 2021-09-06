Home24
===

This application analyzes the given url and returns the statistics of the given url.

## Setup
### Run the application
```sh
# download git repo
$ git clone git@github.com:lakshanwd/home24.git /path/to/repo

# navigate to the repo cloned location
$ cd /path/to/repo

# run the application
$ go run main.go
```
Open http://localhost:3000 on your browser

#### Assumptions made
* treated the urls with same host as internal links, others links reated as exernal links

#### Further improvements
A major drawback of this application is, it analyzes the availability of the links by making http calls to each one
of the links synchonusly.  If we were able to analyze several urls at once in a controlled pararell environment, it would
improve the proformance of the solutions drastically.

#### Screenshots
![screenshot1](/screenshots/screenshot1.png)
![screenshot2](/screenshots/screenshot2.png)
## Deployment

### On a Virtual Machine

* Build the application using following commands
    ```sh
    # navigate to the repo cloned location
    $ cd /path/to/repo

    # download dependancies 
    $ go mod download

    # build an executable of the application
    $ go build -o executable
    ```
* install nginx on the server
* deploy executable as a service
* configure a reverse-proxy such as nginx to listen on port 3000 and serve the application

### Serve on a docker environment

* execute the following commands
    ```sh
    # build docker image
    $ docker build --pull -f Dockerfile -t home24:latest .

    # run image
    $ docker run -d -p 80:3000 home24:latest
    ```
* Open browser application and navigate to http://localhost