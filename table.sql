SET timezone = 'Asia/Bangkok';

create table speaker (
    id      SERIAL PRIMARY KEY,
    name    varchar(250) not null
);

create table events (
    id             SERIAL PRIMARY KEY,
    name            varchar(250) not null,
    location        varchar(250) not null,
    start_at        timestamptz not null
);

create table event_speaker (
    id          SERIAL PRIMARY KEY,
    event_id    integer not null,
    speaker_id  integer not null,
    constraint fk_event_speaker_event_fk foreign key(event_id) references events(id),
    constraint fk_event_speaker_speaker_fk foreign key(speaker_id) references speaker(id),
    constraint UQ_eventID_speakerID UNIQUE(event_id, speaker_id)
);

create table visitor (
    id              SERIAL PRIMARY KEY,
    event_id       	integer not null,
    name            varchar(250) not null,
    email           varchar(250) not null,
    invite_code     varchar(250) not null unique,
    is_accept       boolean not null default FALSE,
    constraint fk_visitor_event foreign key(event_id) references events(id)
);


insert into speaker (name) values('Speaker01'), ('Speaker02');

insert into events (name, "location", start_at) 
values('Golang class 102', 'Bangkok', '2022-12-23 10:30:25');

insert into event_speaker (event_id, speaker_id) values(2, 2);

insert into visitor (event_id, name, email, invite_code) 
values 
	(1, 'visitor01', 'visitor01@gmail.com', 'generate_random_code_invite_code_visitor01'),
	(1, 'visitor02', 'visitor02@gmail.com', 'generate_random_code_invite_code_visitor02'),
	(1, 'visitor03', 'visitor03@gmail.com', 'generate_random_code_invite_code_visitor03');