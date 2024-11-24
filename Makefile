# Nodemon with verbose output (-V) watches *.go files for changes (-e .go -w .)
# Executes "go run main.go" on changes
# Runs once (--count=1) and listens for multiple termination signals
# SIGTERM (terminate), SIGHUP (hangup), SIGINT (interrupt), SIGQUIT (quit), SIGABRT (abort), SIGUSR1 (user-defined)
run:
	@bun run nodemon -V -e .go -w . -x go run main.go --count=1 --race -V --signal SIGTERM SIGHUP SIGINT SIGQUIT SIGABRT SIGUSR1

watch:
	@bun run nodemon -e .templ -w . -x "make compile"

compile: tailwindcss templ

tailwindcss:
	@bun run tailwindcss --config tailwind.config.js -i input.css -o ./public/css/styles.css --minify

templ:
	@templ generate ./views

.PHONY: run compile tailwindcss templ
