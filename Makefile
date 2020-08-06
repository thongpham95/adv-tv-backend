
all: setup_start clean

.PHONY: setup_start
setup_start:
	@docker build -t adv-server .
	@docker run --name adv-server -p 9000:9000 -d adv-server
	

.PHONY: clean
clean:
	@docker rmi $(shell docker images -qa -f 'dangling=true')


.PHONY: stop
stop:
	@docker stop adv-server
	@docker rm adv-server
