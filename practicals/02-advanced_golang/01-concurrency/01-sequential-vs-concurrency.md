- Sequential Execution
    - Process of executing one command after another.
    - This operation keeps the cpu hold for performing single operation
- Concurrency
    - It is the ability of the OS to execute more than one task simultaneously on CPU (Referring to single Core CPU)
    - This can be achieved by switching between task in small time quantum "for eg. doing small process of each task in round robin manner"
    - Concurrency is the notion of multiple things happening at the same time
    - It is the potential for multiple process to be in progress at the same time
    - Best example is Text Editor (Any Text editor) 
- Parallel Processing
    - Parallel processing or Parallelism is using a whole core for single process and if we have multiple core each core will handle one task in parallel 
    -  Eg hadoop based distributed data processing

- GoRoutine
    - Go routines are light weight thread that has a separate independent execution
    - Can execute concurrently with other go-routine
    - All go routine are managed by Go runtime scheduler

    - SYNTAX:
        ``` go <Function/Method> ```
    - Eg:
        ``` go calculate() ```

    - **Whatever the function/ method called after the 'go' keyword will be executed in a separate go routine**