run:
	lsof -i :8080 | awk 'NR!=1 {print $$2}' | xargs -r kill -9
	go run .


history:
	cat ~/.zsh_history | grep "make" | sort | uniq -c | sort -nr