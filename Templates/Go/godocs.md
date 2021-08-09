# RUN GODOC OFFLINE

> Rules to Follow <br>
> - Go to your Project Folder first and open the terminal (because when your run it outside your folder it will not read it.. so its one at a time)
> - 

> Guide

Title (Title must have a Capital letter to be recognize as Title)

Second line is always the description of the title

Third line (if the third line is capital then its a new Title but if not then its the continuation of the second line)

  `1 Tab is consider a code container like this`
  
  
> Best Practices

### Before Package 

Comments before the package will appear in the Overiew of the Godoc.<br>
It must consist of 3 paragraph
- Title (should be capital letter)
- General description of where it is used
- Where it is used/called in the system. 

```
/*
Customer

Customer model collect all the data of the user and later used to their order and recommendation

Customer model is used by Payment,Order,Official Receipt models.

*/
package models

import

...
```

### Before *type* Example

Comment should be before the *type*, It will show as the description of the type Name<br>
Comment should explain When it will be used in the system

```
//Customer is the person ordering before entering the Menu Category
type Customer struct {
	uadmin.Model
	GuestName   string `uadmin:"search" sql:"type:text;"`
	MacAddress  string `uadmin:"search" sql:"type:text;"`
	FacebookAPI string `sql:"type:text;"`
}
```

### Before func Example

Comment should be before the *func*, It will show as the description of the func Name<br>
Comment should explain the general description of the function

```
//String function used as Identifier
func (c *Customer) String() string {
	return c.GuestName
}
```

> INSTALLATION

` go get golang.org/x/tools/cmd/godoc `

> EXECUTION (Go to your folder and Run this to the terminal)

` godoc -http=:6060 `

> ACCESS (Open the link in your browser)

` http://localhost:6060/pkg/github.com/YourFolder/YourProject `
