
@set GOROOT=D:\go
@set VENDER=%cd%\vendor
@set GOPATH=%VENDER%;C:\Users\dyh\go;%GOROOT%
@SETX GOPATH %GOPATH%
@set PATH=%PATH%;%GOROOT%\bin;%VENDER%\bin;

@cmd.exe