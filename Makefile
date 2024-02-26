.PHONY: watch-templ watch-server watch-all
watch-all:
	$(MAKE) -j2 watch-templ watch-server
watch-templ:
	templ generate --watch
watch-server:
	air