# HTTP_Server_Test_Task

Using in-memory storage, HTTP Server implements possibility to perform create and read books and articles from the application. It follows REST approach with JSON messages. Business requirements: it is not allowable to create a new item with the same name (duplications by name are not acceptable) or with other types except “ARTICLE” or “BOOK”.

Instructions:

JSON structure to send data (to save record or to get stored record):

{
  “type”: [string],
  “name”: [string]
}

JSON structure of the data that can be returned:

{
  “type”: [string],
  “name”: [string]
}

or

{
  “success”: [bool],
  “name”:    [string]
}

HTTP requests (all data receives and returns in JSON):

1. /literature   - returns all literature stored in memory

2. /store        - to save new record in storage, returns special response about execution success

3. /getlitname   - returns record that contains name provided in request or special response in case of unsuccessful execution

4. /getlittype   - returns all records that contains type provided in request or special response in case of unsuccessful execution

