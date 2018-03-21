**POST /registration/**

Register an attendee

Headers

| Header          | Description                  | Required  |
| --------------- | ---------------------------- | --------- |
| `Content-Type`  | application/json             | Yes       |
| `Authorization` | a valid authentication token | Yes       |

Request Body

````
{
	"attendee": {
		"shirtsize": "M",
		"diet": "NONE",
		"age": 19,
		"graduationyear": 2020,
		"transportation": "NOT_NEEDED",
		"school": "University of Illinois Urbana-Champaign",
		"major": "Computer Science",
		"gender": "MALE",
		"professionalinterest": "INTERNSHIP",
		"github": "ASankaran",
		"linkedin": "ArnavSankaran",
		"interests": "Something interesting",
		"isnovice": false,
		"isprivate": false,
		"haslightninginterest": false,
		"phonenumber": "5555555555",
		"collaborators": [
			{
				"collaborator": "Collab1@gmail.com"
			},
			{
				"collaborator": "Collab2@gmail.com"
			}
		],
		"longforms": [
			{
				"info": "This is a longform info."
			}
		],
		"extrainfos": [
			{
				"website": "example.com"
			}
		]
	}
}
````

Response Body

````
{
    "ID": 5,
    "CreatedAt": "2018-03-20T22:16:42.370643518-05:00",
    "UpdatedAt": "2018-03-20T22:16:42.370643518-05:00",
    "DeletedAt": null,
    "userid": 2,
    "shirtsize": "M",
    "diet": "NONE",
    "age": 19,
    "graduationyear": 2020,
    "transportation": "NOT_NEEDED",
    "school": "University of Illinois Urbana-Champaign",
    "major": "Computer Science",
    "gender": "MALE",
    "professionalinterest": "INTERNSHIP",
    "github": "ASankaran",
    "linkedin": "ArnavSankaran",
    "interests": "Something interesting",
    "status": "PENDING",
    "phonenumber": "5555555555",
    "reviewertime": "0001-01-01T00:00:00Z",
    "collaborators": [
        {
            "ID": 5,
            "CreatedAt": "2018-03-20T22:16:42.372116152-05:00",
            "UpdatedAt": "2018-03-20T22:16:42.372116152-05:00",
            "DeletedAt": null,
            "collaborator": "Collab1@gmail.com",
            "AttendeeID": 5
        },
        {
            "ID": 6,
            "CreatedAt": "2018-03-20T22:16:42.372706138-05:00",
            "UpdatedAt": "2018-03-20T22:16:42.372706138-05:00",
            "DeletedAt": null,
            "collaborator": "Collab2@gmail.com",
            "AttendeeID": 5
        }
    ],
    "longforms": [
        {
            "ID": 2,
            "CreatedAt": "2018-03-20T22:16:42.373153506-05:00",
            "UpdatedAt": "2018-03-20T22:16:42.373153506-05:00",
            "DeletedAt": null,
            "info": "This is a longform info.",
            "AttendeeID": 5
        }
    ],
    "extrainfos": [
        {
            "ID": 2,
            "CreatedAt": "2018-03-20T22:16:42.379617421-05:00",
            "UpdatedAt": "2018-03-20T22:16:42.379617421-05:00",
            "DeletedAt": null,
            "website": "example.com",
            "AttendeeID": 5
        }
    ]
}
````

**GET /registration/**

Get an logged in attendee

Headers

| Header          | Description                  | Required  |
| --------------- | ---------------------------- | --------- |
| `Content-Type`  | application/json             | Yes       |
| `Authorization` | a valid authentication token | Yes       |

Response Body

Same as response when originally creating the attendee

**GET /registration/{id}**

Get an attendee by id

Headers

| Header          | Description                  | Required  |
| --------------- | ---------------------------- | --------- |
| `Content-Type`  | application/json             | Yes       |
| `Authorization` | a valid authentication token | Yes       |

Response Body

Similar for the response for GET /registration/ but for the attendee with the given id
