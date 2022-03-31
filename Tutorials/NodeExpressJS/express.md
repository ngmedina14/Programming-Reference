# Reference for Express JS

### Preparation

- Create a proj folder
- Initialize and configure npm `npm init`
- Install dependencies

> Install to project folder directory
> - `npm install nodemon` monitoring script
> - `npm install express --save-dev` Node.js web application framework that provides a robust set of features for web and mobile applications.
> - `npm install mysql2`Fast mysql driver. Implements core protocol, prepared statements, ssl and compression in native JS, made for mysql 8.0 to latest
> - `npm install express-handlebars` Template. The constructor function which holds the internal implementation on its prototype.
> - `npm install dotenv --save` Loads environment variables from .env file.
> - `npm install hbs` Express.js template engine plugin for Handlebars.
> - `npm i bcrypt` Hash and verify passwords 


### File Structure
```
Project-Folder
  ├── controllers  ---------- < handlers for all transactions
  │   └── authAccount.js  --- < transaction for users account and validate inputs
  ├── node_modules  --------- < all dependencies is saved here
  ├── routes  --------------- < API or redirections or display page 
  │   ├── Auth.js  ---------- < redirection of all transaction
  │   └── PageRoutes.js  ---- < navigate pages
  ├── views  ---------------- < HTML template or handlebars 
  │   ├── list.hbs
  │   ├── login.hbs
  │   ├── profile.hbs
  │   └── registration.hbs
  ├── .env  ----------------- < database environtment values
  ├── app.js  --------------- < main file
  ├── package.json  --------- < application configurations
  └── package-lock.json
```

### Sample code

index.js
```javascript
const express = require('express');
const app = express();
const port = 2000;

app.use(express.json())
const userRoutes = require('./routes/User')

app.use('/users',userRoutes)

app.listen(port,()=>{
    console.log("Server Started")
})
```
---------------------------
User.js
```javascript
const express = require('express');
const router = express.Router();//essential to get the POST and GET method

router.get('/fetch',(request,response)=>{
    // response.send("list of users");
    // response.json({
    //     name:"Neil",
    //     gender:"male",
    // })
    response.send(request.query);
});
router.post('/insert',(request,response)=>{
    // res.json({
    //     message:"test"
    // })
    console.log(request.body);
    response.send(request.headers);
})

module.exports = router;
```
### Run the App
`nodemon app.js`
