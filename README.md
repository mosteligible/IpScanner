# IpScanner

Keeps checking the IP address of server it is installed on and sends notification if the IP changes. This is useful when the server's IP changes over time .i.e. the server's IP address is not reserved.

The problem I had was, I was serving different services for my home via my router however Public IP address for my home was changing. So I made this program to help me keep track of changing IP addresses.


## Setup running the program on your machine even though it restarts

- Build the program with following command:
```
go build -ldflags "-X main.webhookURL=<your slack webhook url>
```

- Create a file `ipchecker.service` in `/etc/systemd/system/`
```
sudo touch /etc/systemd/system/ipchecker.service
```

- Edit the file `ipchecker.service`
```
sudo nano /etc/systemd/system/ipchecker.service
```

- Add the following content (or change them based on your requirements)
```
[Unit]
Description=ipchecker_obtains_ip_of_srever_its_installed_on_sends_ip_if_it_changes
After=network.target

[Service]
ExecStart=/path/to/your/ipchecker/binary
Type=simple
Restart=on-failure
RestartSec=300

[Install]
WantedBy=multi-user.target
```

- Change permission for the file
```
sudo chmod 664 /etc/systemd/system/ipchecker.service
```

- Enable the service (if it is not enabled), reload the daemon and start `ipchecker.service`
```
# check if service is enabled
sudo systemctl is-enabled ipchecker.service

# if not enabled, enable it
sudo systemctl enable ipchecker.service

# reload and start
sudo systemctl daemon-reload
sudo systemctl start ipchecker.service
```

- check if `ipchecker.service` started successfully
```
sudo systemctl status ichecker.service

# output will look like following

● ipchecker.service - ipchecker_obtains_ip_of_srever_its_installed_on_sends_ip_if_it_changes
     Loaded: loaded (/etc/systemd/system/ipchecker.service; enabled; vendor preset: enabled)
     Active: active (running) since Tue 2022-06-07 15:30:50 EDT; 3h 45min ago
   Main PID: 2329 (ipchecker)
      Tasks: 10 (limit: 9097)
     Memory: 6.5M
     CGroup: /system.slice/ipchecker.service
             └─2329 /path/to/your/ipchecker/binary

Jun 07 15:30:50 mydevice systemd[1]: Started ipchecker_obtains_ip_of_srever_its_installed_on_sends_ip_if_it_changes.
Jun 07 15:30:50 mydevice ipchecker[2329]: 2022-06-07 15:30:50.314285042 -0400 EDT m=+0.210756808 : Ip Detection Request Sent. Most recent Ip: 111.22.66.852
Jun 07 15:30:50 mydevice ipchecker[2329]: Ip Address change detected, changed from <  > to < 111.22.66.852 >
Jun 07 15:30:50 mydevice ipchecker[2329]: Get failed with error:  200 OK
Jun 07 16:30:50 mydevice ipchecker[2329]: 2022-06-07 16:30:50.669412942 -0400 EDT m=+3600.565884850 : Ip Detection Request Sent. Most recent Ip: 111.22.66.852
Jun 07 17:30:50 mydevice ipchecker[2329]: 2022-06-07 17:30:50.789912636 -0400 EDT m=+7200.686384423 : Ip Detection Request Sent. Most recent Ip: 111.22.66.852
Jun 07 18:30:50 mydevice ipchecker[2329]: 2022-06-07 18:30:50.925316342 -0400 EDT m=+10800.821788139 : Ip Detection Request Sent. Most recent Ip: 111.22.66.852
```
