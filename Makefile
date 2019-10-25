TARGET=parser_xml_barcodes.exe

all: clean build

clean:
	rm -rf $(TARGET)

build:
	go build -o $(TARGET) main.go
