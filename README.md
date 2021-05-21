# motion_webmonitor

If you are using motion with one or more Raspberry Pi (or whatever) to monitor your home, you should have noticed that
the web interface is really not great (there isn't even HTTPS). The objective is therefore to make it possible to use
a "normal" web interface in order to manage one or more cameras.

This software should be able to manage other IP cameras than just Raspberry Pi with motion. 

### How to install ?

```
go build
./motion_webmonitor
```
Next create a .passwd file respecting the following format:

```
user1:bcryptHashedPassword
user2:$2y$12$9k223DMvQgKh7df3K.gCDukgUD3LKBCxwOS8MabsVt4zx3TMyCAP.
```
Chmod it correctly the file in order to prevent unwanted modifications.