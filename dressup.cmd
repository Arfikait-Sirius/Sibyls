@ECHO OFF

set DIR_SIBYLS=%1\sibyls
set DIR_GIRLS=%DIR_SIBYLS%\girls

IF NOT exist %DIR_SIBYLS% (
     mkdir %DIR_SIBYLS%
     mkdir %DIR_GIRLS%
)

copy .\dir_Girls\Emily\emily.go %DIR_GIRLS%\. > NUL
copy .\dir_Girls\Louise\louise.go %DIR_GIRLS%\. > NUL

set DIR_HOME=%cd%
cd %1
go mod init %~n1
cd %DIR_HOME%

echo Completed!
