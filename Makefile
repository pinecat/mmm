GO=go
BUILDDIR=build
INSTALLDIR=/usr/local/bin
CONFIGDIR=/usr/local/etc/mmm

all: mmm

mmm:
	mkdir -p $(BUILDDIR)
	go build -o build/mmm ./mmm.go

install:
	cp $(BUILDDIR)/mmm $(INSTALLDIR)/mmm
	mkdir -p $(CONFIGDIR)
	chmod 666 $(CONFIGDIR)

uninstall:
	rm -f $(INSTALLDIR)/mmm
	rm -rf $(CONFIGDIR)

clean:
	rm -rf $(BUILDDIR)