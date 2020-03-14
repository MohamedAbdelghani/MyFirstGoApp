## Starting the Development Server

Navigate to the directory where you want the project to live and run the following commands:

```
git clone https://github.com/MohamedAbdelghani/MyFirstGoApp.git
```
```
cd MyFirstGoApp
```
```
go run main.go
```
The api should be running on localhost:8080

# Routes
### `POST /AddPerson`
Adds new person

### `PUT /EditPerson/{id}`
Updates person with the specified id

### `DELETE /DeletePerson/{id}`
Removes person with the specified id

### `GET /GetById/{id}`
Retrieves person with the specified id.

### `GET /GetByWeight/{weight}`
Retrieves persons who have specified weight.

### `GET /GetByHeight/{height}`
Retrieves persons who have specified height.

