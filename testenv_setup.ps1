New-Item test -ItemType Directory
Set-Location test
New-Item test1.txt -ItemType File
(Get-ChildItem test1.txt).CreationTime = (Get-Date).AddYears(-2)
New-Item test2.txt -ItemType File
