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

### Config file explanations:

**imagesdir:**

Path on which motion will save videos

**authorizedextensions:**

File types inside `imagesdir` that will be accessible from web UI (typically `.mp4`, `.mkv`, `.jpg`...)

**cameras:**

URLs of MJPEG streams, user will be able to choose stream in web UI

**commands:**

+ Command to start motion service
+ Command to stop motion service
+ Command to check if motion service is active (typically `systemctl check motion`)
+ Value returned by previous command if motion service is active (typically `active`)
+ Value returned by previous command if motion service is inactive (typically `inactive`)

**notsecuremodeport:**

Port on which motion_webmonitor listens in case `tls` is set to **false**. This is useful in case you plan to run motion_webmonitor behind a web proxy (nginx, caddy, apache2...)

**domains:**

In case you don't want to run motion_webmonitor behind a web proxy, it's able to serve through HTTPS, and so that fetches a Let's Encrypt SSL certificate. This field is used to tell Let's Encrypt the domain of the certificate (so the one you will use for your website).

**passwordfile:**

File containing credentials, as explained above. Format is `username:bcrypt_hashed_password`, 1 credential per line. **Don't forget to chmod the file correctly**
