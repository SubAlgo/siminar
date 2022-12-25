# DB Design
## TABLE: EVENTS
Field Name | Data type | Constraints | Default value | Format | Description
--- | --- | --- | --- |--- |---
id | serial | pk | | |  
name |varchar(250) |not null | | | event name
location | varchar(250) | not null | | | event location
start_at | timestamptz |not null | | 2022-12-24 10:30:25 |set timezone = Asia/Bangkok

## TABLE: SPEAKER
Field Name | Data type | Constraints | Default value | Format | Description
--- | --- | --- | --- |--- |---
id | serial | pk | | |  
name |varchar(250) |not null | | | speaker name


## TABLE: EVENT_SPEAKER
Field Name | Data type | Constraints | Default value | Format | Description
--- | --- | --- | --- |--- |---
id | serial | pk | | |  
event_id | integer |not null, FK (events.id) | | | 
speaker_id | integer |not null, FK (speaker.id) | | | 

### constraint `UQ_eventID_speakerID UNIQUE(event_id, speaker_id)`


## TABLE: VISITOR
Field Name | Data type | Constraints | Default value | Format | Description
--- | --- | --- | --- |--- |---
id | serial | pk | | |  
event_id | integer | FK(events.id)
name | varchar(250) | not null
email | varchar(250) | not null
invite_code | varchar(250) | not null, unique
is_accept | boolean | not null | false

## Create docker database
`docker run -d --name=postgres12 --restart=always \ 
    -p 5432:5432 \
    -e POSTGRES_USER=postgres \
    -e POSTGRES_PASSWORD=test1234 \
    -e POSTGRES_DB=seminar \
    postgres:12.1
`
## Create and init data to DB
ดูที่ table.sql [https://github.com/SubAlgo/siminar/blob/master/table.sql]


## Run app
`go run main.go`

## CURL 
### search visitor by event_id
`curl --location --request GET 'http://127.0.0.1:3000/visitor?event_id=1'`

### search visitor by event_name
`curl --location --request GET 'http://127.0.0.1:3000/visitor?event_name=Golang class 101'`











