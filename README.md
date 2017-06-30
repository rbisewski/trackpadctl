# trackpadctl - Trackpad control tool for Linux, written in golang.

This is a rather plain golang-based method of using xinput to control certain
trackpads settings on Linux.

Part of my reason for writing this was a need for a more minimalist way of
painlessly stopping laptop trackpads on bootup.

When it comes to trackpads on Linux, I expect there is a great deal of
non-conformity, so go ahead and shoot me an email if you see any obvious
mistakes on your hardware.

Maybe one day it will have more features, but for now it is more of a
simple tool. Feel free to fork it and use it for other projects if you find
it useful.


# Requirements

The following is needed in order for this to function as intended:

* Linux kernel 3.0+
* golang
* xinput

Older kernels could still give some kind of result, but I *think* most of
the newer distros use the latest version of xinput, which should allow for
broad compatibility. Feel free to email me if this is incorrect.


# Running

1) Install the program.

    make install

2) On bootup the trackpad settings should be enforced.


# Todos

* Remove hard coded values
* Allow for more devices, possibly a config file

# Authors

Written by Robert Bisewski at Ibis Cybernetics. For more information, contact:

* Website -> www.ibiscybernetics.com

* Email -> contact@ibiscybernetics.com
