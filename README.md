dos_ssh
=======

Use BIOS ram hacks to make a SSH server out of any INT 10 13h app (MS-DOS is one of those)

You can find a demo Youtube Video here below:

[![Youtube Video](http://img.youtube.com/vi/2JrugnykXmg/0.jpg)](http://www.youtube.com/watch?v=2JrugnykXmg)


howto run
=========

* Get a DOS compatible floppy disk image (e.g. from here: http://www.allbootdisks.com/download/dos.html)
* Install golang and qemu, e.g. `sudo apt-get install qemu golang`
* Compile dos_ssh: `go get; go build`
* Start qemu with this image: `qemu-system-i386 -fda Dos6.22.img -boot a -vnc :0 -s`
* Run dos_ssh: `./dos_ssh`
* Connect to ssh server: `ssh localhost -p 2222`

