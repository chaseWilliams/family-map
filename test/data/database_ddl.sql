-- sqlite


-- CREATE TABLES
create table Persons
(
	person_id varchar(30) not null
		constraint Persons_pk
			primary key,
	username varchar(30) not null,
	first_name varchar(30) not null,
	last_name varchar(30) not null,
	gender varchar(1) not null,
	father_id varchar(30)
		references Persons
			on update set null on delete set null,
	mother_id varchar(30)
		references Persons
			on update set null on delete set null,
	spouse_id varchar(30)
		references Persons
			on update set null on delete set null
);

create unique index Persons_person_id_uindex
	on Persons (person_id);




create table Users
(
	username varchar(30) not null
		constraint Users_pk
			primary key,
	password varchar(30) not null,
	email varchar(30) not null,
	first_name varchar(30) not null,
	last_name varchar(30) not null,
	gender varchar(1) not null,
	person_id varchar(30) not null
);



create table Auth
(
	username varchar(30) not null
		constraint auth_username
			references Users
				on update cascade on delete cascade,
	auth_token varchar(50) not null
		constraint Auth_pk
			primary key
);

create unique index Auth_auth_token_uindex
	on Auth (auth_token);


create table Events (
  event_id varchar(30) not null constraint Events_pk primary key,
  username varchar(30) not null references Users on update cascade on delete cascade,
  person_id varchar(30) not null references Persons on update cascade on delete cascade,
  latitude float not null,
  longitude float not null,
  country varchar(30) not null,
  city varchar(30) not null,
  event_type varchar(30) not null,
  year integer not null
);
create unique index Events_event_id_uindex on Events (event_id);

create table Cities
(
	id int
		constraint Cities_pk
			primary key,
	city_type varchar(30),
	wiki_data_id varchar(30),
	city varchar(100) not null,
	name varchar(100) not null,
	country varchar(100) not null,
	country_code varchar(30) not null,
	region varchar(30) not null,
	region_code varchar(30),
	latitude float not null,
	longitude float not null
);

