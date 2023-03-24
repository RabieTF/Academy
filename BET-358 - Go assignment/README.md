# **Usage**: 
+ Download and install go on your machine.
+ Clone this folder.
```
git clone https://github.com/RabieTF/Academy/tree/main/BET-358%20-%20Go%20-%20-assignment
```

+ Move into module directory
```
cd "BET-358 - Go assignment"
```
+ Install necessary libraries
```
go get .
```
+ Execute by using entry file main.go
```
go run main.go
```


# **Endpoints**:
## **Users**: 
### **POST** /users: Creates a new account.
+ 201 if successful.
+ 500 if internal server during processing.
+ 400 if user doesn't respect correct format.
+ Example data :
``` 
{
    "name": "Your name",
    "email": "you_email@example.com",
    "password": "yourPassword"
}
``` 

### **POST** /login: Logs the user in, gives back a JWT token.
+ 200 if successful.
+ 400 if request body incorrect
+ 401 if wrong credentials.
+ Example data:
``` 
{
    "email": "you_email@example.com",
    "password": "yourPassword"
}
``` 

## **Shops**:
### **POST** /shops: Creates a new shop. **Requires authentification.**
+ 201 if successful.
+ 500 if internal error.
+ 400 if incorrect format.
+ Example data: 
```
{
    "name": "name_of_shop",
    "address": "physical_address_of_shop"
}
```

### **GET** /shops : Returns all the available shops. 
+ 200 and all the shops if successful.
+ 404 if there are no shops in database.
+ 500 if internal error.

### **GET** /shops/:id : Returns the shop with the same id in the parameter.
+ 200 and the requested shop if successful.
+ 404 if the requested shop doesn't exist in database.
+ 500 if internal error.

### **PUT** /shops/:id : Updates the shop with the same id in the paramater. **Requires authentification and user must own the shop**
+ 200 if successful.
+ 400 if bad formatting.
+ 403 if user isn't owner of this shop.
+ 404 if shop doesn't exist.
+ 500 if something went wrong.
+ Example data:
```
{
    "name": "name_of_shop",
    "address": "physical_address_of_shop"
}
```

### **DELETE** /shops/:id : Deletes shop with the same id as the paramater. **Requires authentification and user must own the shop**
+ 200 if successful.
+ 403 if user doesn't own this shop.
+ 404 if shop doesn't exist.
+ 500 if something went wrong


## **Products**:

### **POST** /products: Creates a new product. **Requires authentification.**
> Categories should be one string, separated by a comma, and they must be in the predefined categories, see categories endpoint below.
+ 201 if successful.
+ 400 if incorrect JSON format.
+ 403 if user is attempting to create a new product in a shop he doesn't own.
+ 500 if something went wrong.
+ Example data: 
```
{
    "ShopID": 1,
    "Name": "Burger",
    "Description": "A great burger",
    "Categories": "Food, Electronics"
}
```

### **GET** /products : Returns all available products.
+ 200 and all the products if successful.
+ 404 if there are no products in database.
+ 500 if internal error.

### **GET** /products/:id : Returns the product with the same id as parameter.
+ 200 and the requested product if successful.
+ 404 if the requested product doesn't exist.
+ 500 if internal error.

### **PUT** /products/:id : Updates the product with the same id in the parameter. **Requires authentification and user must own the shop where the product belongs**
> Categories should be one string, separated by a comma, and they must be in the predefined categories, see categories endpoint below.
+ 200 if successful.
+ 400 for bad formatting.
+ 403 if user isn't owner of the shop where the product belongs.
+ 404 if product doesn't exist.
+ 500 if something went wrong.
+ Example data: 
```
{
    "Name": "Burger",
    "Description": "A great burger",
    "Categories": "Food, Electronics"
}
```

### **DELETE** /products/:id : Deletes product with the same id as the paramater. **Requires authentification and user must own the shop where the product belongs**
+ 200 if successful.
+ 400 for bad formatting.
+ 403 if user isn't owner of the shop where the product belongs.
+ 404 if product doesn't exist.
+ 500 if something went wrong.

## **Categories**:

## **GET** /categories : returns the predefined categories from the database.
These predefined categories MUST be used when creating or updating a new product, otherwise you will receive an error.