# motion_webmonitor

If you are using motion with one or more Raspberry Pi (or whatever) to monitor your home, you should have noticed that
the web interface is really not great (there isn't even HTTPS). The objective is therefore to make it possible to use
a "normal" web interface in order to manage one or more cameras.

This software should be able to manage other IP cameras than just Raspberry Pi with motion.

### How to install ?

At first create a .passwd file respecting the following format:

```
user1:bcryptHashedPassword
user2:$2y$12$9k223DMvQgKh7df3K.gCDukgUD3LKBCxwOS8MabsVt4zx3TMyCAP.
```

Chmod it correctly the file in order to prevent unwanted modifications.

```
go build -ldflags "-s -w"
sudo ./motion_webmonitor
```

#### Config file

Write a JSON configuration file like

```
{
  "imagesdir" : "/var/lib/motion",
  "authorizedextensions" : [
    ".mp4",
    ".mkv"
  ],
  "cameras" : [
    "http://192.168.1.10:3000/",
    "http://192.168.1.11:3000/"
  ],
  "commands" : [
    "service motion start",
    "service motion stop",
    "systemctl check motion",
    "active",
    "inactive"
  ],
  "notsecuremodeport" : 8080,
  "tls": true,
  "domains": ["www.example.com"],
  "passwordfile" : "/path/to/.passwd"
}
```

And run the program: `sudo ./motion_webmonitor config.json`