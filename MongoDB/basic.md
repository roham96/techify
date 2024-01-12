
## Connect MongoDB Shell
#
Connect to 
mongodb://localhost:27017 by default
```
mongo
```

Connect to local network
```
mongo "mongodb://127.0.0.1:27017"
```
Connect to MongoDB Atlas
```
mongo "mongodbURI" --username <username>"
```
</br>

## Database Operation
#
Show All Databases
```
show dbs
```

Show Current Database
```
db
```
Create / Switch to new Database
```
use shopify
```
Delete Database
```
db.dropDatabase()
```
</br>

## Collections
#
Show all of collections of current DB
```
show collections
````
Create new Collection
```
db.createCollection('articles')
```

</br>

## Create Documents
#
Insert one document
```
db.articles.insertOne({
    name: "Learn MongoDB",
    price: 80
})
```
or
```
db.articles.insert({
    name: "Learn MongoDB",
    price: 80
})
```
Insert multiple document
```
db.articles.insert(
    {
    name: "Learn MongoDB",
    price: 80
    },
    {
    name: "Learn GoLang",
    price: 120
    }
)
```
</br>

## Read Documents
#
Find one document
```
db.articles.findOne()
```
Find multiple documents
```
db.articles.find()
```
Find multiple documents with json format
```
db.articles.find().pretty()
```
Find documents by field value
```
db.articles.find({
    "name": "Learn GoLang",
})
```
Find documents by Element in Array
```
db.articles.find({
    comments:{
        $elemMatch: {
            "user": "Roham"
        }
    }
})
```
</br>

## Update Documents
#
Update one 
```
db.articles.updateOne({
    {"_id": 1},
    {
        $set:{
        "name": "Learn GoLang"
        }
    }
})
```
Update Multiple 
```
db.articles.update({
    {"price": 120},
    {
        $set:{
        "price": 110
        }
    }
})
```
Increment field Value
```
db.articles.update({
    {"_id": 2},
    {
        $inc:{
        "views": 85
        }
    }
})
```
Rename field
```
db.articles.updateOne({
    {"_id": 3},
    {
        $rename:{
        "views": 'likes'
        }
    }
})
```
Update Sub-Documents
```
db.articles.update({
    {name: "Learn MongoDB"},
    {
        $set:{
            comments:[
                {
                    "user":"Mehdi",
                    "body":"Comment one",
                },   {
                    "user":"Yaghoob",
                    "body":"Comment Two",
                },
            ]
        }
    }
})
```
</br>

## Delete Documents
#
Delete a single document
```
db.articles.remove({
    "name": "Learn GoLang"
})
```
</br>

## Count
#
Returns the count of documents
```
db.articles.find().count()
```
Returns the count of documents that would be matched by the find() query
```
db.articles.find({
    "views": 85
}).count()
```
</br>

## Sorting
#
Ascending order of documents
```
db.articles.find().sort({
    "name": 1
}).pretty()
```
Descending order of documents
```
db.articles.find().sort({
    "name": -1
}).pretty()
```
</br>

## Results by Pagination
#
Skip 7 results
```
db.articles.find({}).skip(7)
```
Fetch only 5 results
```
db.articles.find({}).limit(5)
```
Sort by  name, Skip first 7 results, fetch only next 5 documents
```
db.articles.find({}).sort({
    "name": 1
}).skip(7).limit(5)
```
</br>

## Add Index
#
Create Index on single field
```
db.articles.createIndex({
    "name": 1
})
```
Create compound Index
```
db.articles.createIndex({
    "name": 1,
    "date": 1
})
```
</br>

## Drop Index
#
Drop Index
```
db.articles.dropIndex{"field_name")
```
</br>

## Text Search
#
Create Text Index on field
```
db.articles.createIndex({
    "content": "text"
})
```
Search by Text
```
db.articles.find({
    $content: {
        $search: "PHP"
    }
})
```
</br>

## Find by range query
#
Find articles where views are grater than 60
```
db.articles.find({
    "views": {'$gt': 60}
})
```
Find articles where views are grater than or equal to 60
```
db.articles.find({
    "views": {'$gte': 60}
})
```
Find articles where views are less than 60
```
db.articles.find({
    "views": {'$lt': 60}
})
```
Find articles where views are less than or equal to 60
```
db.articles.find({
    "views": {'$lte': 60}
})
```
##
### Thank you for supporting !