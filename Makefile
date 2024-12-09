install:
	go build -o flint main.go; rm -r /home/jacex/.local/bin/flint; mv flint /home/jacex/.local/bin;