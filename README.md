- **Call Method:** This is the core method responsible for passing requests to the Telegram API.

- Encapsulation: The Call method will be wrapped in other methods such as GetUpdate or SendMessage to interact with the Telegram API as needed.

````
+-------------------------------+
|        Call Method            |
+-------------------------------+
|                               |
| Inputs:                       |
| 1. method (API method name)   |
| 2. token (Telegram token)     |
| 3. params (parameters for API)|
|                               |
+---------------|---------------+
                |
                v
+-------------------------------+
|  Construct URL                |
|  - Uses token and method      |
+---------------|---------------+
                |
                v
+-------------------------------+
|  Prepare Request Body         |
|  - Serializes params to JSON  |
+---------------|---------------+
                |
                v
+-------------------------------+
|  Send HTTP POST Request       |
|  - To Telegram API URL        |
+---------------|---------------+
                |
                v
+-------------------------------+
|  Receive Response             |
|  - From Telegram API          |
+---------------|---------------+
                |
                v
+-------------------------------+
|  Parse Response               |
|  - Unmarshal JSON to struct   |
+---------------|---------------+
                |
                v
+-------------------------------+
|  Return Response              |
|  - To the caller method       |
+-------------------------------+
````