.PHONY: fmt vet templ css air clean

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

templ:
	templ generate --watch

css:
	bunx @tailwindcss/cli -i ./views/templates/styles/index.css -o ./views/static/styles.css --watch

air: vet
	air

clean:
	rm -rf ./tmp/*
