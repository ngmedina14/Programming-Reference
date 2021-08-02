# Best Practices for uadmin code setup

## main.go
	func main(){
		uadmin.Register
		uadmin.Settings
		http.Handle Function
		uadmin.Custom Translation
		Functions for push
		Other Function
		uadmin.Start Server
		}

	func DBConfig()


------------



## Naming Standard and when 2 words
#### Naming files uses snake_case
Word should start with Lowercase letter and use Underscore _ for space
##### files
- **some_example.go**
<br>

#### Naming type and struct and func uses CamelCase

Every words should start with Uppercase letter and should not use space 

```
type NameSample struct{

SomeExample string
OtherExampleID string

 } 
```


------------



## uadmin Custom UI file location
###### template/
- **custom_ui.html**

###### view/
- **backend_custom_ui.go**
