ff: an ffmpeg (or ffprobe) command-line parameter builder
========================================================

Implements useful functions for building params for a call to ffmpeg or ffprobe.
--------------------------------------------------------------------------------

It was written, partially, as an exploration of TDD (Test Driven Development).

*To install:*

    go get -t github.com/xdave/ff
(the -t flag ensures deps are pulled in required to run the tests)

*To run the tests:*

    go test -v github.com/xdave/ff

*To run the tests in your browser:*

    cd $GOPATH/src/github.com/xdave/ff
    $GOPATH/bin/goconvey
(Then, open your browser to the URL given in the output)

* See [Documentation](http://godoc.org/github.com/xdave/ff) for this pkg.
* The tests are written using [GoConvey](http://goconvey.co/) (BDD-style).


Copyright 2015 Dave Gradwell

Released under a BSD-style license (see LICENSE file)
