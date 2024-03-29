<p align="center">
  <img src="https://github.com/Superredstone/Port-Scanner/blob/main/assets/icon.png"></img> <br> <br>
  <img src="https://img.shields.io/github/license/Superredstone/Port-Scanner"></img>
  <img src="https://img.shields.io/discord/821836676607115304?color=Blue&logo=Discord&logoColor=Blue"></img>
</p>

# Screenshots
<p>
  <img width=50% height=50% src="https://github.com/Superredstone/Port-Scanner/blob/main/assets/screenshot1.png"></img><br>
  <img width=50% height=50% src="https://github.com/Superredstone/Port-Scanner/blob/main/assets/screenshot2.png"></img><br>
  <img width=30% height=30% src="https://github.com/Superredstone/Port-Scanner/blob/main/assets/screenshot3.jpg"></img>
</p>

# How to build (Please look at this guide, don't try it lonely)
It is recommended to build on Linux or MacOS

```bash
git clone https://github.com/Superredstone/Port-Scanner.git
cd Port-Scanner/
```

## Windows 
### amd64
```bash
env GOOS=windows GOARCH=amd64 go build -o PortScanner-windows-amd64.exe -ldflags="-H windowsgui" .
```
### arm
```bash
env GOOS=windows GOARCH=arm go build -o PortScanner-windows-amd64.exe -ldflags="-H windowsgui" .
```
### Adding rsrc
```bash
go get github.com/akavel/rsrc
rsrc -arch amd64 -ico assets/icon.ico
```

## Linux
```bash
env GOOS=linux GOARCH=amd64 go build -o PortScanner-linux-amd64 .
```

## MacOS
*You can build MacOS only if you are on MacOS*
```bash
env GOOS=darwin GOARCH=amd64 go build -o PortScanner-macos-amd64 .
```

## iOS
*You can build for iPhones only if you are on MacOS*
```bash
env GOOS=
```

## Android
To build Android you need [Android Studio SDK](https://developer.android.com/studio?gclsrc=ds&gclsrc=ds&gclid=CLvek6-H-fACFedDHQkdISICpw) and [JDK](https://www.oracle.com/java/technologies/javase-jdk16-downloads.html)

*DOON'T REMOVE THE DOT AT THE END OF THE COMMAND*
```bash
gogio -target android -o build/PortScanner-android.apk -icon assets/icon.png -appid com.portscanner.superredstone .
```

# Donate :heart:
Every donation is appriciated <3 <br> <br>
<a href='https://ko-fi.com/A0A64PC0Y' target='_blank'><img height='36' style='border:0px;height:36px;' src='https://cdn.ko-fi.com/cdn/kofi3.png?v=2' border='0' alt='Buy Me a Coffee at ko-fi.com' /></a>

# License
GNU GPL V3
