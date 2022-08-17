# Task_App


# Project Title
```
The project I developed is a proprietary issue tracking product that enables bug tracking and agile project management.
My project, which is a demo software, is a project software that can be used in more than ten industries. 
This software is a simple backend architecture with no frontend. In the future, the advanced version of the 
project will be developed by me.
```
## What Does It Do?
```
1- Developing new projects.

2- Managing the process in long-term projects quickly and without errors

3- Detecting and correcting the mistakes made thanks to its sophistication.

4- To work more systematically and to complete existing works to the dead line.

5- To make reporting and business data analysis.

6- Importing databases.
```
## Getting Started

### Prerequisites

```
In order to run the project successfully, the requirements are golang version 1.18.2 
as programming language, a suitable IDE and postman or similar Api test tool for testing 
are required.
```

### Installing

```
For Go installation, you can refer to my previous installation post.

https://github.com/FurkanSamaraz/GoKurulum

1- Go to the location of main.go

2- Download repository.

	go get -u github.com/gmvbr/httptest
	
	go get -u github.com/gofiber/fiber/v2

3- go run main.go


```
## Running the project


```
1- First create a user by going to the "/UserCreate" extension in Postman.

2- Go to "/StatusCreate" extension in Postman and add status.

3- Go to the "/EntryCreate" extension in the Postman and create an entry.

4- Add entryrelation to the entry you created by going to the "/EntryReltCreate" 
extension in the Postman.

5- Go to the "/EntryComCreate" extension in the Postman and assign a comment to 
the entry you have created by specifying the user.

6- You can update the allowed items of the entry you have created by going to the 
"/EntryUpdate" extension in the Postman.

7-You can update the allowed items of the entryrelation you have created by going 
to the "/EntryReltUpdate" extension in the post.

8-You can update the allowed items of the status you have created by going to the 
"/StatusUpdate" extension in the post.

9-You can pull the entries created in that date range by specifying the date range 
on the creation date of the entry you created by going to the "/EntryTimeCreGet" 
extension in the post.

10-You can pull the updated entries in that date range by specifying the date range 
on the update dates of the entry you created by going to the "/EntryTimeUpdGet" 
extension in the post.

11-You can pull the entries you have created by going to the "/EntryGet" extension 
in Postman.

12-You can get the entries you have created by going to the "/EntryStatusGet" 
extension in Postman according to the status filtering.

13-You can pull the entries you have created by going to the "/EntryTagGet" 
extension in Postman according to Tag filtering.

14-You can pull the entries you have created by going to the "/EntryAllGet" 
extension in the post with the assigned entryrelation information.

15-You can pull the entries you have created by going to the "/EntryTrueGet" 
extension in Postman according to True/False filtering.

16-You can save the entries you have created by going to the "/EntryTopAllGet" 
extension in the post, the comments, entryrelation, user information etc. assigned 
to the entry. You can pull the features as a whole.

17- You can go to the "/EntryComGet" extension in the Postman and retrieve the 
comments assigned to the entry you have created.

18- You can delete the comments assigned to the entry you have created by going 
to the "/EntryComRemove" extension in the Postman.

19- You can pull the entryrelations you have created by going to the "/EntryeReltGet" 
extension in the Postman.

20- By going to the "/EntryeReltMainGet" extension in the Postman, you can filter 
the entryrelations you have created by filtering whichever entry is assigned to it.

21- You can get the statuses you have created by going to the "/StatusGet" extension 
in Postman.

22- You can pull the users you have created by going to the "/UsersGet" extension 
in Postman.

23- You can pull the users you have created by going to the "/UsersActiveGet" extension 
in the Postman according to the True/False filtering.

24- By going to the "/UsersEntryAllGet" extension in the Postman, you can filter all 
the properties assigned to the users you have created, by filtering them.




```
![User_Diagram](https://user-images.githubusercontent.com/92402372/185144033-d490217e-fc78-41f8-aef5-c47c711d0e06.png)
![Status_Diagram](https://user-images.githubusercontent.com/92402372/185144099-eb0d2bbe-f824-4d7a-b8bd-e343b5a1ba4c.png)
![Entry_Relation_Diagram](https://user-images.githubusercontent.com/92402372/185144154-0cd8cb2a-ac88-4f0e-872f-69bee1161733.png)
![Comments_Diagram](https://user-images.githubusercontent.com/92402372/185144193-1e17b9b7-40e1-472c-81bc-d9ec3e9be493.png)
![Entry_Diagram](https://user-images.githubusercontent.com/92402372/185144223-87e58ae4-e048-48f9-a0f5-f6ef8add8426.png)

## Videos
Visit my dropbox extension for videos;

https://www.dropbox.com/s/onuem2hjskfjkf0/1.mov?dl=0

https://www.dropbox.com/s/egke24wsh4po2lr/2.mov?dl=0

https://www.dropbox.com/s/65j3ph5emriftn7/3.mov?dl=0

## Used technologies

```
Technologies used in the project;

1- vscode

2- Golang

3- Golang fiber library for APIs

4- httptest library for Golang tests

5- sqlite was used as database to store data, then revised with postgresql

6- Postman to test API

7- Flow chart for drawing the project chart and ease of visual understanding

```
## Running the tests

```
_test.go tests were written for the files with the .go extension in the repository. 
You can run and check the tests by going to the appropriate file location with the 
'go test -v' command or by opening the _test.go file and clicking the play button 
next to the functions.

```

### And coding style tests

```
There are two kinds of test algorithms in the _test.go files in the repository. 
1 for Get requests. 2. Test algorithms for Post requests. httptest library is 
used for these test algorithms.

```


## Deployment

```
You can make the project live by creating a subscription with heroku, which 
I like in terms of interface and usage, or you can do the same by creating 
a container with docker.

```
## Build With

```
Specify OS or Architecture in build:
env GOOS=linux go build main.go # builds for Linux
env GOARCH=arm go build main.go # builds for ARM architecture

Build multiple files:
go build main.go assets.go # outputs an executable: main

Building a package:
go build . # outputs an executable with name as the name of enclosing folder

```
