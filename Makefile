default:  test install

test:
	./test.sh

install:
	go install -ldflags="-s -w" ./...

.PHONY: test \
	install
