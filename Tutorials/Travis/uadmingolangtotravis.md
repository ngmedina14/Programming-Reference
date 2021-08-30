# Running Uadmin Go App in Travis
#### This is a bare minimum for running go in travis

----------------------

##### Prerequisite
- Github Repository
- Travis account using github account
- Uadmin App or any Go application
- Activate GitHub Apps Integration (go to Travis open settings)
- Sync Github repository located at the Travis settings

##### Getting Started in Travis

###### Go to your Github Repo and add a travis file

> *.travis.yml*

```yml
#this is required
language: go

#Here you can choose your version
go:
- 1.17.x
#i think this is important in a long run ? just include this
go_import_path: github.com/account/repo

# Skip dependency fetch. We store all dependencies under vendor/.
install: true
```

*even you dont have the `go test` file it will run fine*

