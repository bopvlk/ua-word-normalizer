.SILENT:

inst-lsb-release:
	apt install lsb-release

prepare: inst-lsb-release
	curl -fsSL https://packages.redis.io/gpg | sudo gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg

echo: prepare
	echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/redis.list

redis-update: echo
	apt-get update

inst-redis: redis-update
	apt-get install redis