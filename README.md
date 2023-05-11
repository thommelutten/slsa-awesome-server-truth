# SLSA Awesome Server
Awesome Server written Golang used for SLSA exercises

To run the awesome-server run the following in a terminal

```
go run awesome-server
```

Awesome-server can be accessed in the following ways:

Powershell:
```
Invoke-RestMethod http://localhost:8090/awesome
```

Bash:
```
curl localhost:8090/awesome
```

Both should output the following to terminal

```
Who's awesome? You're awesome!
```