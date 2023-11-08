src := main.go
builddir := ./build

# Blatently stole this name from Tsoding
program := $(builddir)/binviz

clean:
	rm -fr $(builddir)/*

build: clean
	go build -o $(program) $(src)
