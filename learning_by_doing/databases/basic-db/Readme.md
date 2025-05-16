# Basic Database Connection
## Key Takeaway for Coding
- When using native library for Database we need drivers, the reason is covered later
- To connect to the DB you need to pass DSN(DataSourceName) as input
- The DSN should contain username password and db address with port
```
Syntax: "<username>:<password>@tcp(<hostaddress>)/<dbname>?parseTime=true"
```
- DB connection comes with err handling by default so the syntax is 
```
db, err := sql.Open(<dbtype>,dsn)
```
- Handle err with panic option
- Always close the connection when not needed

## Learnings
- "database/sql" is a standard library for SQL database interactions in Go. Since it doesn't provide drivers for specific databases we need to import the drivers externally. 
- We are using MySQL hence we used '_ "github.com/go-sql-driver/mysql'.
- How you know the syntax: We can google to get suitable library for mysql and in the corresponding library documentation we will have the syntax. No need to worry on that. 
- Github reference for this library https://github.com/go-sql-driver/mysql

### Why we need to import driver
- "database/sql is a general API". i.e it only provides functionality like Open(), Query(), Exec(), Ping(), etc which are common to db. But it doesn't know how to speak to specific flavour of db like mysql, postgres etc. 
- So in order for ""database/sql" talk to corresponding flavour or db efficiently we need to import the db driver.



