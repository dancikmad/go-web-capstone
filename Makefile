# Makefile 
#
# # Set default target 
.PHONY: up down open 

# Build and run containers in detached mode 
up:
		docker-compose up --build -d 


# Open API Url in the browser 
open:
		@echo "Opening http://localhost:4002"
		open http://localhost:4001

down:
		docker-compose down 

run: up open 
