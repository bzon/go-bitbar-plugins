BIN?=bin/currency_converter.30m.go
BITBAR_PLUGINS_DIR=${HOME}/Bitbar

b:
	rm -fr bin/
	go build -o $(BIN)
d: b
	cp $(BIN) $(BITBAR_PLUGINS_DIR)/
