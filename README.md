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
Of Marian's history 402 If to the Flowers of your young hearts Instructions dews are given Oh be earnest as our Marian was To find the road to Heaven 402 If to the Flowers of your young hearts Instructions dews are given Oh be earnest as our Marian was To find the road to Heaven 
====> 55 words ad4079a72f481c8c0dd83e1311da6c8dd08cf8439b91972758fa26c70ec7cc74

FOOTNOTES 1 Bedford jail in which Bunyan was twelve years a prisoner 2 Tophet here means hell 3 Idle one 4 An old word meaning money or riches 5 This word means pleasant or delightful 6 Perspective glass is an old name for a telescope or spy glass 7 An atheist is one who does not believe that there is a God 8 That is of the body and blood of Christ 9 An instrument of music used in the time of John Bunyan somewhat like a very small piano 10 An old English coin bearing the figure of an angel 11 The word let here means hindrance 
====> 108 words debf61244bc38adc724d1cc4a9c2035e82573f7f2ffbe1b312135d9c4898743d
...
```

##SUCCESS!
