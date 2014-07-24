ocd
===

This project contains parsers for RSS (0.9 through 2) and Atom feeds using the xml package in the standard library. main.go exists mainly for testing, once it is built, a single URL can be passed to it and it will print out all of the items/entries in the feed.

    ./ocd https://news.ycombinator.com/rss
    
The current setup should allow new modules to be easily added without conflict. If you would like to add a module, please send a pull request with a working implementation. See the modules directory for examples. All modules should be properly namespaced.
