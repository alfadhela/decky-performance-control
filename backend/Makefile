# This is the default target, which will be built when 
# you invoke make
.PHONY: all
all: backend

# This rule tells make how to build hello from hello.cpp
backend:
	mkdir -p ./out
	go build -ldflags "-s -w" -o ./out/backend ./src/main.go 

# This rule tells make to delete hello and hello.o
.PHONY: clean 
clean:
	rm -f backend