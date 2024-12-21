


==> Model : 

A Model is a class that's your primary tool for interacting with MongoDB. An instance of a Model is called a Document.

In Mongoose, the term "Model" refers to subclasses of the mongoose.Model class. You should not use the mongoose.Model class directly. The mongoose.model() and connection.model() functions create subclasses of mongoose.Model as shown below  :: https://mongoosejs.com/docs/api/model.html#Model().

From ChatGpt : 
1.A model in Mongoose is a wrapper around a MongoDB collection and provides an .interface for interacting with the collection.
2.Models are created based on schemas. They allow you to perform CRUD operations (Create, Read, Update, Delete) on MongoDB documents and execute queries against the database.
3.Models represent specific collections in the database and have methods for querying, updating, deleting, and validating documents.

1.A model in Mongoose is a constructor function that represents a MongoDB collection and provides an .interface for querying, creating, updating, and deleting documents in that collection.
2.Models are created based on a Mongoose schema. Each model corresponds to a MongoDB collection, and instances of the model represent documents in that collection.
3.Mongoose provides methods on the model for interacting with the database, such as find, create, update, and delete.




