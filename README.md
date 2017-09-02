It is a simple Authentication in ROR, which uses Json Web Token.

Create a database and a table named user to store user information.
when you store the password in database encrypt using Bcrypt hashing algorithm and store in user table.

User table contains id, username and password.

CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(45) DEFAULT NULL,
  `password` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1


Encrypt the string password which you have using Bcrypy Hashing algorithm, You can generate it online and store it in password field against username.

POST  http://127.0.0.1:8080/login

request is in following format (Application/json)

{
	"username":"admin",
	"password":"password"
}

It will generate the token. And you should pass the token in the next request header.


GET http://127.0.0.1:8080/simple

request header will be 

Authorization Bearer {Your-token-here}
