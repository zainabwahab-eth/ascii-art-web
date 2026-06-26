.PHONY: watch

watch:
	while true; do \
		go run . & \
		inotifywait -r -e modify,create,delete --include='.*\.(go|html|css)' . ; \
		kill $$(lsof -t -i:8080); \
	done