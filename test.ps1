Write-Output "create test directories and files with modified creation date"
New-Item test -ItemType Directory
Set-Location test
New-Item test1.txt -ItemType File
(Get-ChildItem test1.txt).CreationTime = (Get-Date).AddYears(-2)
New-Item test2.txt -ItemType File
New-Item anothertest -ItemType Directory
Set-Location anothertest
New-Item test3.txt -ItemType File
(Get-ChildItem test3.txt).CreationTime = (Get-Date).AddYears(-2)
Set-Location ..
Set-Location ..
Write-Output "build main.go and run it for folder test"
go build .\main.go
.\main.exe "test"
Write-Output "check if creation dates are within a minute of the current date"
$testdate = (Get-Date).AddMinutes(-1)
if ((Get-ChildItem test\test1.txt).CreationTime -gt $testdate) {
    Write-Output "date of test\test1.txt successfully bumped"
} else {
    Write-Output "error"
}
if ((Get-ChildItem test\anothertest\test3.txt).CreationTime -gt $testdate) {
    Write-Output "date of test\anothertest\test3.txt successfully bumped"
} else {
    Write-Output "error"
}
Remove-Item test -Recurse