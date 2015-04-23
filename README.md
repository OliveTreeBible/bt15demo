**This is a work in progress please be sure to update the repository before the presentation to be sure you have the latest code.**

##Installation
####Preparing the Go Environment

This demo is written in [Go](http://golang.org)

For complete install instructions go to [Install Go](http://golang.org/doc/install)

#####Installing Go (the shorter version)
I have tested these instructions on a Mac, your mileage may vary on other platforms but the steps will be similar.

On Mac it is very likely you will need Xcode and its 'Command Line Tools' counterpart. I am not certain since all of my machines already have it installed.

You will also need [git](http://git-scm.com) and [Mercurial](http://mercurial.selenic.com) as Go simply uses on whatever version control system the package maintainer uses and both git and mercurial are used by the packages the demo needs.

1) Download go for your platform from here https://golang.org/dl/ 
	I downloaded this: https://storage.googleapis.com/golang/go1.4.2.darwin-amd64-osx10.8.tar.gz

2) Extract the files (in my case the command is)
	
```
$ sudo tar -C /usr/local -xzf go1.4.2.darwin-amd64-osx10.8.tar.gz
```
3) Add the binary path to your bash profile by opening `.bash_profile` and adding the line

```
export PATH=$PATH:/usr/local/go/bin
```
4) Either log out and back in **or** `source` your `.bash_profile`

5) Create your `go` workspace (this can be anywhere for me it is directly under my home directory)

```
$ mkdir go_workspace
```

6) Set the `$GOPATH` environment variable based on the workspace directory you created in step 5. Open your .bash_profile again and add the line

```
export GOPATH=$HOME/go_workspace
export PATH=$PATH:$HOME/go_workspace/bin
```
7) Either log out and back in **or** `source` your `.bash_profile`

####Getting the Demo's Dependencies and Source 
8) Now to install the packages we need to run the demo. These will be installed into your `$GOPATH`

```
$ go get github.com/blevesearch/segment
$ go get github.com/PuerkitoBio/goquery
$ go get code.google.com/p/go-sqlite/go1/sqlite3
```

9) Now grab the source for the demo

```
$ cd ~/go_workspace/src/github.com/
$ mkdir olivetreebible
$ olivetreebible
$ git clone https://github.com/OliveTreeBible/bt15demo.git
```

####Building and Running the Demo
10) Go into the project directory and build and run it

```
$ cd bt15demo/
$ go install
```

This compiled the source code to a binary and installed it to `$GOPATH/bin` and since we added that path to the `$PATH` shell variable in step 6 we can now simply run the app like this:

```
bt15demo
```

You will see output that looks something like

```
...
Minor typographical errors punctuation and inconsistencies have been silently normalized Archaic spelling has been retained 
====> 15 words 78aaee5c82ce7d28257b8259ffd291ce4bc5f2110bbbded9f298b988b62c559a

Page 365 Tell truth has been changed to Tell true 
====> 10 words 533bde1a50485de830d42664c307e9e8ace79e909bff7682aa79d1ca1efcb865
```

##SUCCESS!
