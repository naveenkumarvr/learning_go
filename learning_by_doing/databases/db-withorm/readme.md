# DB with ORM
- ORM stands for Object Relational Mapping
- It is a technique(tool) that helps us to interact with db in go lang struct instead of writing an sql Query
```
Example
SELECT * FROM users id = 1;

Instead of writing like this we can write

// Here db.First is one of the Gorm method and we are pointing to the struct user and telling gorm to get the id = 1
db.First(&user, 1) 

// Your struct can be 
type user struct {
ID uint `grom:"primaryKey"`
Name string
Email string
}
```

## Gorm Syntax
```
gorm.Open(dialector, config)
```
- dialector : Tells gorm which database you are using (Mysql postrges etc)
- config: gives additional settings to customize Gorm behavior

Modified Syntax with some DB
```
dsn := "root:password@tcp(127.0.0.1:3306)/mysql?charset=utf8&parseTime=True&loc=Local"

gorm.Open(mysql.Open(dsn), &gorm.Config{})
```
- In the above code we are saying hey GORM use mysql driver and open the connection using DSN.(Remember the DSN which contains all the details about your server such as username passwrod, server ip, port and DB name.)

### gorm Config
- The last part of gorm is the config , i.e ***&gorm.Config{}***. This is pointer to the Gorm Default config struct. If you want to change the behaviour of the Gorm you can modify or mention the parameter here
- Some example such as more details you can find here https://gorm.io/docs/gorm_config.html
```
&gorm.Config{
    Logger: logger.Default.LogMode(logger.Info), // show SQL logs
    DisableForeignKeyConstraintWhenMigrating: true,
}
```

## Why we need var DB *gorm.DB
Now I had question why the variable DB is declared and why it is pointing to gorm.DB

- We declare DB variable which is a pointer to **gorm.DB struct**
- The gorm.DB struct contains the information like db connections, config and other internal states
We use a pointer (*gorm.DB) because:
- Gorm returns and excepts pointers from gorm.Open() and in its method calls
- Using a pointer allows us to share the single db connection acrros the whole application without copying the entire struct
- When we call methods like DB.AutoMigrate()or DB.Create() those methods are needs pointer as receivers so the value will be automatically dereferenced. 