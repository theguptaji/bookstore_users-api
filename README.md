# bookstore_users-api
RESTFUL API in golang. Developed with MVC architecture standards.

## To Run 
* Setup a Mysql database.
* Prepare a database named "userDB".
* Create a table with following config
```sql
CREATE TABLE `users` (
	`id` BIGINT(64) NOT NULL AUTO_INCREMENT,
	`first_name` VARCHAR(45),
	`last_name` VARCHAR(45),
	`email` VARCHAR(45) NOT NULL,
	`data_created` VARCHAR(45) NOT NULL,
	`status` VARCHAR(30) NOT NULL,
	`password` VARCHAR(45),
	PRIMARY KEY (`id`)
);
```
* We basically wants a schema like this
```
+--------------+-------------+------+-----+---------+----------------+
| Field        | Type        | Null | Key | Default | Extra          |
+--------------+-------------+------+-----+---------+----------------+
| id           | bigint      | NO   | PRI | NULL    | auto_increment |
| first_name   | varchar(45) | YES  |     | NULL    |                |
| last_name    | varchar(45) | YES  |     | NULL    |                |
| email        | varchar(45) | NO   | UNI | NULL    |                |
| date_created | varchar(45) | NO   |     | NULL    |                |
| status       | varchar(45) | NO   |     | NULL    |                |
| password     | varchar(32) | NO   |     | NULL    |                |
+--------------+-------------+------+-----+---------+----------------+
```
* Now set the following environment variables
`mysql_users_username`, `mysql_users_password`, `mysql_users_host`, `mysql_users_schema`
* At this point you can clone the repo and run the code
```
git clone <repo>
go run main.go
```
* The application will be exposed at port 8080.

## To Test
* Prepare `curl` or http rest client like `postman` or `insomnia`.

### Create a User by
```
curl -x POST http://localhost:8080/users -d @data.json
```
where data.json could be:-
```json
{
  "id": 1,
  "first_name": "aman",
  "last_name": "gupta",
  "email": "aman@email.com"
  "password": "some-password"
}
```
* Make sure, your POST request data is following the standards of insert validation defined in `domain/users/user_dto.go`, you can also edit this function as per your needs :-
```go
func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	// TODO: Add validation for password
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
```

### Get a User by
* This API offers two kinds of get user request, `PUBLIC` and `PRIVATE`
* For public request set a header `X-public=true` and insert a user_id you want to get correspounding user to.
* For private request no need to set header, but insert the user_id
```
curl -x GET http://localhost:8080/users/<id>
```
* For private request you get following response:-
```json
{
  "id": 1,
  "first_name": "aman",
  "last_name": "gupta",
  "email": "aman@email.com",
  "date_created": "15046-04-15 14:15:52",
  "status": "active"
}
```
* For public request you get following response:-
```json
{
  "id": 1,
  "date_created": "15046-04-15 14:15:52",
  "status": "active"
}
```

### Update a User by
```
curl -x UPDATE http://localhost:8080/users/<id> -d @data.json
```
or
```
curl -x PATCH http://localhost:8080/users/<id> -d @data.json
```
* Here patch will not update if an empty key is passed

### Delete a User by
```
curl -x DELETE http://localhost:8080/users/<id> 
```
