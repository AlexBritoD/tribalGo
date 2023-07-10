CREATE TABLE songs (
	idpk serial4 NOT NULL,
	id int4 NULL,
	"name" varchar(255) NOT NULL,
	artist varchar(255) NOT NULL,
	duration varchar(10) NOT NULL,
	album varchar(255) NOT NULL,
	artwork varchar(255) NOT NULL,
	price varchar(10) NOT NULL,
	origin varchar(20) NOT NULL,
	CONSTRAINT songs_pkey PRIMARY KEY (idpk)
);