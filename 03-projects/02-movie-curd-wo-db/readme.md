## Go Mod Init
Before starting any go project create a module so that all your dependencies are managed in the module.
```
go mod init example.com/greetings 
```
Here **example.com** is used to tell audience that this is just an example as most of the go modules are managed in github

## Import Module Usages
| Module | Usage |
|:-:|:-|
|`fmt` | As we will be printing out messages such as sever connected etc |
|`log` | We will use logs to print errors and log information and warnings |
|`encoding/json` | This is used to pass the input as Json so that we can pass it to the API for writing |
|`math/rand` | This is used to generate random number which is used as Id for our movies |
|`net/http` | This creates a web server |
|`strconv` | The ID created by Math.Random is a integer but we prefer to store the ID as string so we convert int ID to string and store it |
|`github.com/gorilla/mux` | Request handler to match URL pattern and routes the traffic to corresponding functions |
