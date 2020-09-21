GO=go
BUILDDIR=build
INSTALLDIR=~/bin

all: mmm

mmm:
	mkdir -p $(BUILDDIR)
	go build -o build/mmm src/mmm.go

install:
	cp $(BUILDDIR)/mmm $(INSTALLDIR)/mmm

uninstall:
	rm -f $(INSTALLDIR)/mmm

clean:
	rm -rf $(BUILDDIR)