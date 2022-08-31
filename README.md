# go_pgsql_orm


## summary
this is a simple example of go-pgsql-orm usage.
i designed this simple lib to help me  learn go and postgresql better.

there are many improvements that can be done to this lib.

with this small lib you can:

- create schema
- create table
- insert data
- read one record

### this is a simple example of go-pgsql-orm usage.:

how to create schema:

```go
package main

import (
	"fmt"
	"time"
	managers "github.com/aliaqa256/go_pgsql_orm/managerspkg"
)

type User struct {

	Id uint  `orm:"serial PRIMARY KEY"`
	CreatedAt time.Time 
	UpdatedAt time.Time 
	UserName string `orm:"varchar(255) not null"`
	Email string `orm:"varchar(255) not null"`
	Phone string `orm:"varchar(255) not null"`
	Age  float32  `orm:"int"`
}


func main() {
    modelpkg.Migrate(a)
}
```

how to insert data:

```go
import (
	"fmt"
	"time"
	managers "github.com/aliaqa256/go_pgsql_orm/managerspkg"
)
func main() {
    user := &User{
        UserName: "Aliaqa",
        Email: "alilotfi256@gmail.com"
        age: 25,
        Phone: "09121234567",
    }
    managers.Create(a)
}
```

how to get id of one record:

```go
dbc:=managers.GetId(a,map[string]string{
 	"username": "aliaqa",
    "email": "alilotfi256@gmail.com"
 })

 fmt.Println(dbc.Resualt)
 ```

 how to get < filed > of one record:

```go                   
                    ///   struct,id,fieldname
	dbc:=managers.GetField(a,26,"age")
	fmt.Println(dbc.Resualt)
```

if you have any suggestions or questions please contact me at :
```
https://alilotfidev.ir
alilotfi256@gmail.com
```
