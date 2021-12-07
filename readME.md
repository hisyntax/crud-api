# CRUD API task with a User authentication 

This is a REST API with Golang, MongoDB and using jwt-token to authorize users

## Clone the repository
git clone [https://githb.com/hisyntax/crud-api](https://githb.com/hisyntax/crud-api)

## After cloning the repo navigate to the project directory

## Create a .env file and then copy the content of the example.env file into it and add the requires values

## Install Dependencies

gi get ./...

## Run the app

go run main.go or gowatch

# REST API

The REST API to the app is described below

## Signup

### Request

`POST- /user/signup`
     127.0.0.1:8000/api/v1/user/signup
    {
        "email": "your email address",
        "password": "your password"
    }

### Response
    {
        "user_id": {
            "InsertedID": "61af77afa3e61ca5f77d5914"
        }
    }


## Signin

### Request

`POST /user/signin`
    127.0.0.1:8000/user/signin
    {
        "email": "your email address",
        "password": "your password"
    }

### Response

    {
        "ID": "61af77afa3e61ca5f77d5914",
        "email": "email",
        "Password": "hashed password",
        "token": "token",
        "refresh_token": "refresh token",
        "created_at": "2021-12-07T15:03:11Z",
        "updated_at": "2021-12-07T15:03:11Z",
        "user_id": "61af77afa3e61ca5f77d5914"
    }

## Create a post

### Request
`POST /post/create`

    127.0.0.1:8000/api/v1/post/create

    Token needs to be passed into the header 
    token : "your generated token"

    {
        "title": "This is a test post",
        "body": "This is the body post"
    }

### Response

    {
        "created_at": "2021-12-07T16:17:43+01:00",
        "title": "This is a test post",
        "user_id": {
            "InsertedID": "61af7b17a3e61ca5f77d5916"
        }
    }

## Get a post

### Request
`GET /post/:post_id`
    127.0.0.1:8000/api/v1/post/:post_id

    Token needs to be passed into the header 
    token : "your generated token"


### Response

    {
        "post": {
            "ID": "61af7b17a3e61ca5f77d5916",
            "title": "This is a test post",
            "body": "This is the test body",
            "created_at": "2021-12-07T15:17:43Z",
            "updated_at": "2021-12-07T15:17:43Z",
            "post_id": "61af7b17a3e61ca5f77d5916"
        }
    }


## Get all post

### Request
`GET /post/posts`
    127.0.0.1:8000/api/v1/post/posts

    Token needs to be passed into the header 
    token : "your generated token"


### Response

    {
        "post": [
            {
            "ID": "61af514cb8a0837743114e13",
            "title": "this is a new title",
            "body": "this is the body",
            "created_at": "2021-12-07T12:19:24Z",
            "updated_at": "2021-12-07T12:19:24Z",
            "post_id": "61af514cb8a0837743114e13"
            },
            {
            "ID": "61af595bc01f395bd3c749c3",
            "title": "this is the second title",
            "body": "this is the body of the second post",
            "created_at": "2021-12-07T12:53:47Z",
            "updated_at": "2021-12-07T12:53:47Z",
            "post_id": "61af595bc01f395bd3c749c3"
            },
            {
            "ID": "61af5985c01f395bd3c749c4",
            "title": "this is the third title",
            "body": "this is the body",
            "created_at": "2021-12-07T12:54:29Z",
            "updated_at": "2021-12-07T12:54:29Z",
            "post_id": "61af5985c01f395bd3c749c4"
            },
            {
            "ID": "61af598ac01f395bd3c749c5",
            "title": "this is the fourth title",
            "body": "this is the body",
            "created_at": "2021-12-07T12:54:34Z",
            "updated_at": "2021-12-07T12:54:34Z",
            "post_id": "61af598ac01f395bd3c749c5"
            },
            {
            "ID": "61af7b17a3e61ca5f77d5916",
            "title": "This is a test post",
            "body": "This is the test body",
            "created_at": "2021-12-07T15:17:43Z",
            "updated_at": "2021-12-07T15:17:43Z",
            "post_id": "61af7b17a3e61ca5f77d5916"
            }
        ]
    }


## Update a post

### Request
`PUT /post/:post_id`
    127.0.0.1:8000/api/v1/post/:post_id

    Token needs to be passed into the header 
    token : "your generated token"

### Request 

    {
        "title": "This is the updated test post",
        "body": "This is the updated test body"
    }


### Response

    {
        "message": "Successfully updated the Post",
        "title": "This is the updated test post"
    }   


## DELETE a post

### Request
`DELETE /post/:post_id`
    127.0.0.1:8000/api/v1/post/:post_id

    Token needs to be passed into the header 
    token : "your generated token"


### Response

    {
        "message": "Successfully Deleted!"
    }
