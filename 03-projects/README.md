# Golang Projects

Reference:
- https://www.youtube.com/playlist?list=PL5dTjWUk_cPYztKD7WxVFluHvpBNM28N9

### Important steps

#### go mod init
- Always init your project as you start which will create go.mod files.
- This will manage all your dependencies of your project. If you are using any other module as a part of your project which will be managed here
##### Why Do You Need go mod init?
- Dependency Management: It allows you to manage the versions of third-party libraries or packages your project depends on.
- Reproducibility: The go.mod file ensures that your project will work the same way on any system, by keeping track of the exact versions of dependencies.
- Go Build Process: Once the go.mod file is initialized, you don’t need to worry about GOPATH or manually managing dependencies. Go automatically fetches and installs dependencies based on your go.mod.
##### Example
`go mod init example/<projname>`

#### go mod tidy
- The go mod tidy command is used to clean up your Go project's dependencies by removing any unnecessary ones and adding any that are missing. Essentially, it ensures that your go.mod and go.sum files are in a correct and minimal state, with only the necessary dependencies included.
##### Why Use go mod tidy?
- Maintain Clean Dependencies: Over time, as you add or remove packages, the go.mod file might accumulate unnecessary dependencies. go mod tidy ensures only the required dependencies remain.
- Avoid Dependency Issues: It helps prevent potential issues where the project may not work because a required dependency is missing from go.mod or an outdated dependency remains in the file.
- Version Control: Running go mod tidy before committing changes ensures your go.mod and go.sum are in an optimal state for other developers working on the project. It makes your dependency management clean and reliable.
- Improves Build Performance: By ensuring that only necessary dependencies are included, you might see slight improvements in the build process, as Go won’t have to fetch unnecessary packages.
##### Example
`go mod tidy`

#### go build
- This is used to build the project.
- This will create binary for your os platform. exe for windows and executable file for mac and linux. the file will be generated in the same name as your project folder
##### Example
`go build`

#### go run
- This will run the go project with raw file i.e main.go file
##### Example
`go run main.go`

