![alt text](./screenshots/kumo-01.png)
# Kumo CLI
A CLI made in Go to fetch HTML pages from the internet.

## Features
- Set custom path to store all your searched files
- Set the amount of files you want to fetch from one search
- With the list flag you can search specific links to download

## Setup
- Have Go installed in your system
- Clone the repository
- Go to your home directory and create a file ".kumo.yaml" containing
```
kumoPath: "/Users/siddharthsingh/Desktop"
search: 2
```
- cd into the Kumo folder and run the following command
```
$ go install github.com/44t4nk1/kumo
``` 
- And you are all set! You can start using Kumo from your terminal now

## Use
```
$ kumo search binary search tree
$ kumo search --list binary search tree
```
- Will put the HTML pages for the search in the path you mentioned.
