# Setup | Installation
## Windows
- Install Go from https://go.dev/dl/
- Press Win+R to open Run command.
- Type 'cmd' then click enter.
- On the newly open terminal, type 'cd Desktop'
- Then clone this repository by 'git clone https://github.com/mahkassem/gogin.git'
- After cloning type 'cd gogin'
- Change '.env.example' to '.env' and change the environment variables to fit your needs.
- Run 'go run main.go' in your terminal.
- Enjoy :) 
## Linux | MacOS
- Soon

##  Usage
>**$** go run main.go `-m`
  - `-m` means that the program will perform migration.
  - program will **stop** after.
>**$** go run main.go `-d`
  - `-d` means that the program will drop all database tables.
  - program will **stop** after.
>**$** go run main.go `-md`
  - combined flags from `-m` & `-d`
  - order does **NOT** matter, dropping tables will always happen first.
  - program will **stop** after.
>**$** go run main.go `-md` `-n`
  - `-n` means normal startup.
  - program will drop all database tables, perform migration and will **NOT STOP** after. 
 
# Important URLS

| Name  |  URL |
|---|---|
| Documentation (Go Gin)  | https://gin-gonic.com/docs  |
| Article (REST API)  | https://www.jetbrains.com/guide/go/tutorials/rest_api_series/gin/  |
| Article (Chapters)  | https://masteringbackend.com/posts/gin-framework  |
| GO ORM  | https://gorm.io/index.html  |
| Github Repository for Go Gin Examples  | https://github.com/gin-gonic/examples  |
| Documentation (Go Gin) Examples  | https://gin-gonic.com/docs/examples/  |