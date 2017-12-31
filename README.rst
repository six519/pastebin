pastebin
========

Pastebin API Client For Go

Installation
============

`go get github.com/six519/pastebin`

Create Paste
============
::

    package main

    import "fmt"
    import "github.com/six519/pastebin"

    func main() {

        client := pastebin.Pastebin{Api_dev_key: "<DEV KEY>"}
        options := pastebin.PastebinOption{}

        ret, err := client.Paste("hello world", options)

        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println(ret)
        }
    }

List Trending Pastes
====================
::

    package main

    import "fmt"
    import "github.com/six519/pastebin"

    func main() {

        client := pastebin.Pastebin{Api_dev_key: "<DEV KEY>"}
        ret, err := client.Trends()

        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println(ret)
        }
    }

Delete Paste
============
::

    package main

    import "fmt"
    import "github.com/six519/pastebin"

    func main() {

        client := pastebin.Pastebin{Api_dev_key: "<DEV KEY>"}

        client.Login("<USERNAME>", "<PASSWORD>")
        ret, err := client.Delete("<PASTE KEY>")

        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println(ret)
        }
    }

Getting User Information
========================
::

    package main

    import "fmt"
    import "github.com/six519/pastebin"

    func main() {

        client := pastebin.Pastebin{Api_dev_key: "<DEV KEY>"}

        client.Login("<USERNAME>", "<PASSWORD>")
        ret, err := client.UserDetails()

        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println(ret)
        }
    }

Listing Pastes
==============
::

    package main

    import "fmt"
    import "github.com/six519/pastebin"

    func main() {

        client := pastebin.Pastebin{Api_dev_key: "<DEV KEY>"}
        options := pastebin.PastebinOption{}

        client.Login("<USERNAME>", "<PASSWORD>")
        ret, err := client.List(options)

        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println(ret)
        }
    }	

Getting Raw Paste
=================
::

    package main

    import "fmt"
    import "github.com/six519/pastebin"

    func main() {

        client := pastebin.Pastebin{Api_dev_key: "<DEV KEY>"}
        
        client.Login("<USERNAME>", "<PASSWORD>")
        ret, err := client.ShowPaste("<PASTE KEY>")

        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println(ret)
        }
    }	
