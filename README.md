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
adds new person

### `PUT /EditPerson`
Updates person

### `DELETE /DeletePerson`
Removes person

### `GET /GetById/{id}`
Retrieves person of specified id.

### `GET /GetByWeight/{weight}`
Retrieves persons who have specified weight.

### `GET /GetByHeight/{height}`
Retrieves persons who have specified height.

