# Version
VERSION = `date +%y.%m`

# If unable to grab the version, default to N/A
ifndef VERSION
    VERSION = "n/a"
endif

#
# Makefile options
#


# State the "phony" targets
.PHONY: all clean build install uninstall


all: build

build:
	@go build

clean:
	@go clean

install: build
	@echo installing executable file to /usr/bin/trackpadctl
	@sudo cp trackpadctl /usr/bin/trackpadctl
	@echo installing cron file to /etc/cron.d/trackpadctl
	@sudo cp trackpadctl.cron /etc/cron.d/trackpadctl

uninstall: clean
	@echo removing executable file from /usr/bin/trackpadctl
	@sudo rm /usr/bin/trackpadctl
	@echo removing cron file from /etc/cron.d/trackpadctl
	@sudo rm /etc/cron.d/trackpadctl
