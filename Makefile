.PHONY: local_env_up
local_env_up:
	/usr/local/bin/docker compose -f /Users/a.gusev/GolandProjects/awesomeProject1/docker-compose.yml -p awesomeproject1 up -d

.PHONY: local_env_down
local_env_down:
	/usr/local/bin/docker compose -f /Users/a.gusev/GolandProjects/awesomeProject1/docker-compose.yml -p awesomeproject1 down


.PHONY: local_env_restart
local_env_restart: local_env_down local_env_up


.PHONY: local_run
local_run:
	go run main.go
