# Email Service - for the coding challenge!

Email Service - for the coding challenge

### The problem:
Create a service that accepts the necessary information and sends emails. It should provide an abstraction between two different email service providers. If one of the services goes down, your service can quickly fail over to a different provider without affecting your customers.


### The solution:
I create a service(focus on backend) that sends the email to the recipient by taking the emails of recipient and sender, the subject and content as inputs.
It is backed by Mailgun and Sendgrid. It always try to use Mailgun first, and if there's anything wrong, it tried to use Sendgrid to send the email.

## Installation/Deployment
* Install Docker
* Build docker images by: ```docker-compose build```
* Put the Sendgrid API key, mailgun API key and the base API url into the ```.env``` file. You can get those by creating free account on [sendgrid](https://sendgrid.com/) and [mailgun](http://www.mailgun.com)
* You can run it locally by ```docker-compose up```  and you can access it on http://127.0.0.1:8080/

## How to use this service
The service APIs are RESTful APIs.

The main API calls should be made with HTTP POST. (Help API can be called with GET)
Any non-0 status code in HTTP response code is an error. The returned message tells more detailed information.

### Main API 
```
URL: /email
```

There is no UI for this project. It is a REST API. It is accessible through HTTP POST requests, expecting a JSON object as input. And it will return an object as output too.


method: POST

input: 
- One from email address
- cc or bcc (optional)
- email subject
- email content (Email subject and email content cannot be both empty)

input format: json

JSON key | Meaning
-------- | -------
from     | string, the sender email address
to       | string the to email address
cc       | string the cc email address (optional)
bcc      | string the bcc email address (optional)
subject  | the email subject
text     | full text content of the email to be sent


Following is a sample input json:
```
{
    'from':'test_from@mail.com',
    'to':'test_to1@mail.com',
    'cc':'test_cc@mail.com',
    'bcc':test_bcc@mail.com,
    'subject':'test subject',
    'text':'This is the test text as the email content. Again, this is the test text as the email content.'
}
```

output:
- status code 
- message 

output format: json
 
Following is a sample output json:
```
{'status': 0, 'message':'Success'}
```
This is for successfully transaction.

You can consider any non-0 status code as an error. The message will give details. 
Following are typical errors, in the format of status code and message:

status code | message
----------- | -------
1           | from email address invalid
2           | to email address invalid
3           | subject is empty
4           | text is empty
5           | all email sender failed

## Testing

Following are the scripts to test the service:

- send email to one address
```
curl -i -H "Content-Type: application/json" -X POST -d \
 '{"from":"nattawut.ru@gmail.com","to":"nattawut.ru@gmail.com", "subject":"test subject","text":"This is an test email."}' \
  http://localhost:8080/email
 
```
output:
```
{
  "message": "success", 
  "status": 0
}
```