# Lumos

Your magic to light you wisdom by delegating to your search engine of course


Building
---

 - Go
 - [Godep]( https://github.com/tools/godep)

 Dependencies are provided on `vendor/` directory, if you whish use `godep get`
 to install them for you.

Running
---

Using google as default engine
 ```
 lumos -q golang

 ```

 Custom engine (Support for: DuckduckGo, Bing and Google)
 ```
 lumos -p duck -q golang
 ```

*Nox*
