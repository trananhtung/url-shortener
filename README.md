# URL Shortener Repository (Go Lang)
This repository is a URL shortener implemented using the Go programming language. It provides a simple and efficient solution for shortening long URLs into shorter, more manageable ones. The repository contains all the necessary code and components to create, manage, and redirect shortened URLs.

# Features
1. URL Shortening: The core functionality of the repository is the ability to shorten long URLs into concise, unique identifiers. This process involves generating a shortened code or key that maps to the original URL. When a user visits the shortened URL, they are redirected to the original destination.

2. Customization: Users can customize the shortened URLs by providing their desired custom slugs. This feature allows them to create memorable and meaningful short URLs. If the custom slug is already taken, the system generates an alternative unique slug automatically.

3. Redirection: The repository handles the redirection of shortened URLs to the original long URLs seamlessly. When a visitor clicks on a shortened URL, the system recognizes the associated key and redirects the user to the correct destination URL.

4. Persistence: The repository ensures the persistence of short URLs and their mappings to the original URLs. It employs a database or a storage mechanism to store and retrieve the necessary information, ensuring that shortened URLs remain functional over time.

5. Analytics: The system may include analytics functionality to track and analyze the usage of shortened URLs. This feature allows users to monitor the number of clicks, geographic distribution, referral sources, and other relevant metrics.

# API Endpoints

## POST /shorten

Create a shortened URL.

**Request**

Request body should contain a JSON object with the following fields:

```json
{
  "url": "The original URL to be shortened"
}
```
Response

Returns the shortened URL in the response body.

## POST /shorten/:hash
Create a shortened URL with a custom hash.

Path Parameters

hash: The custom hash to be used for the shortened URL
Request

Request body should contain a JSON object with the following fields:

```json
{
  "url": "The original URL to be shortened"
}
```

Response

Returns the shortened URL in the response body.

## GET /:hash

Redirect to the original URL.


## GET /record

Get all the records of ip address and url.

## GET /all

Get all the shortened URLs.

# POSTMAN Collection

```json
{
	"info": {
		"_postman_id": "136a03e0-f21b-46f6-aea2-26cd6f9b0c45",
		"name": "URL Shortener",
		"schema": "https://schema.getpostman.com/json/collection/v2.0.0/collection.json",
		"_exporter_id": "19335676",
		"_collection_link": "https://www.postman.com/tunganhtran/workspace/hr-system/collection/19335676-136a03e0-f21b-46f6-aea2-26cd6f9b0c45?action=share&creator=19335676&source=collection_link"
	},
	"item": [
		{
			"name": "get",
			"request": {
				"method": "GET",
				"header": [],
				"url": "http://localhost:8080/3pZ6Xv2v8nX8Rj"
			},
			"response": []
		},
		{
			"name": "get all record",
			"request": {
				"method": "GET",
				"header": [],
				"url": "http://localhost:8080/all"
			},
			"response": []
		},
		{
			"name": "get all shorten link",
			"request": {
				"method": "GET",
				"header": [],
				"url": "http://localhost:8080/all"
			},
			"response": []
		},
		{
			"name": "create shorten link",
			"request": {
				"method": "POST",
				"header": [
					{
						"key": "Content-Type",
						"value": "application/json",
						"type": "text"
					}
				],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"url\": \"google.com\"\n}"
				},
				"url": {
					"raw": "http://localhost:8080/shorten",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"shorten"
					],
					"query": [
						{
							"key": "Content-Type",
							"value": "application/json",
							"disabled": true
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "create shorten link customize",
			"request": {
				"method": "POST",
				"header": [
					{
						"key": "Content-Type",
						"value": "application/json",
						"type": "text"
					}
				],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"url\": \"google.com\"\n}"
				},
				"url": {
					"raw": "http://localhost:8080/shorten",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"shorten"
					],
					"query": [
						{
							"key": "Content-Type",
							"value": "application/json",
							"disabled": true
						}
					]
				}
			},
			"response": []
		}
	]
}
```
