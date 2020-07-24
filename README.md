# Document creation date bumper
## Dependencies
go version go1.14.3
## What does it do?
Changes a documents creation date (**Windows only!**) to system time by copying it when the creation date is 23 months in the past.
## Motivation
Microsoft OneDrive does not allow changing the creation date. If you use the win32 API the changes on file level are not synced into the cloud. The official Office365 API does not allow for a change of creation date either. So the only remaining solution is to create a copy of the file. 
## Warnings!
This will result into OneDrive **uploading the file all over** which creates a lot of traffic when large files are affected,
## Usage
Run test.ps1. It will build main.go, create testfolders and files and checks if everything works as expected. If there are no errors you can run main.exe with the PowerShell-command .\main.exe "C:\path\to\my\folder"