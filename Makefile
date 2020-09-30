GO=go
INSTALLDIR=/usr/local/bin

all: mmm

mmm:
	go build mmm.go

install:
	cp mmm $(INSTALLDIR)/

uninstall:
	rm -f $(INSTALLDIR)/mmm

clean:
	rm -rf $(BUILDDIR)