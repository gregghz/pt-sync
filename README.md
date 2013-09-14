This project is a syncronization server for the eBook application
known as [PageTurner](https://github.com/NightWhistler/PageTurner). It
is an alternative to
[PageTurnerWeb](https://github.com/NightWhistler/PageTurnerWeb). The
main goal is for this server to be easily deployable to Google
Appengine. It uses the Go SDK. Currently it is only suitable for
personal use. That is to say there is no way to distinguish between
different users syncing the same book at present.

It runs like any
[Go Appengine project](https://developers.google.com/appengine/docs/go/gettingstarted/introduction).

Just start up your instance, and put your server info into PageTurner
(you may have to build it from source at present to get the option to
change the sync server from the gui).
