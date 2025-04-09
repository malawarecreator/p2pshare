# p2pshare

well guys guess what<br>

p2psharing! securely!<br>


well, to start the server (globally in da WWW), you gotta first forward port 8888 using vscode or something else like your network control panel<br>

that's like the only extra step you have to take in order to run it. again, the links will be provided when you run the command.<br>

usage: `p2pshare <filepath> <authtoken>`<br>

example output:
```
C:\Users\benjamin\code\p2pshare>p2pshare go.mod ben
Fetching public IP address...

--- Server Starting ---
Serving file: go.mod
Required auth token: ben
Listening on port 8888...

--- Access URLs ---
From THIS machine:    http://localhost:8888/download?authtoken=ben
From SAME network:    http://<YOUR_LOCAL_IP>:8888/download?authtoken=ben (Find local IP with 'ipconfig' or 'ifconfig')
From OUTSIDE network: http://76.218.110.235:8888/download?authtoken=ben (Requires firewall & router port forwarding for port 8080)

Server running. Press Ctrl+C to stop.
Rejected download attempt from 127.0.0.1:62290: incorrect token ''
Serving file 'go.mod' (Type: video/mpeg) to 127.0.0.1:62288
^C
```
<br>

make sense?<br>

## Also:
you are welcome to build this for other platforms (im on windows so i can only make windows distros) and then use THIS program to send me the file (i'm so clever). u gotta dm me the link on discord @biggyballz_69 (that is my actual username)

