CREATE DATABASE wallet_db_alif_winanda

CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	email varchar(256) UNIQUE,
	password varchar(256),
	reset_code varchar(256),
	created_at default current_timestamp,
	updated_at timestamp
);


CREATE SEQUENCE custom_sequencer INCREMENT 1 MINVALUE 270001 MAXVALUE 270999 START 270001;

CREATE TABLE wallets (
	
	wallet_number int PRIMARY KEY DEFAULT NEXTVAL('custom_sequencer'),
	user_id int,
	balance int
);


CREATE TABLE transactions(

	id SERIAL PRIMARY KEY,
	wallet_number int,
	transaction_type varchar,
	source_id int,
	fund_id int,
	target_id int,
	amount int,
	created_at default current_timestamp,
	description varchar
);

CREATE TABLE funds(
	
	id SERIAL PRIMARY KEY,
	source_name varchar
);

--realpass -> pass + nama, ex : rengoku -> passrengoku
INSERT INTO users (email, password)VALUES ('rengoku@kny.com', '$2a$04$cYsCz47rFurDNaWUti2fAOI57Y6S4/0LWh48GtqDudtxzFbGuajw.');
INSERT INTO users (email, password)VALUES ('zenitsu@kny.com', '$2a$04$pYeSyzJ6ozfodRlrHVRtROzZiZ2tONHnFdE5qmZ3yKVkjNudS4Gsm');
INSERT INTO users (email, password)VALUES ('tomioka@kny.com', '$2a$04$cWZ9cRt7n7L/8tm6T392COt9NOytQdiI8bFkf.VCtKIP8A112k.7e');
INSERT INTO users (email, password)VALUES ('deku@mha.com', '$2a$04$FYaYA/qvNDAakkXsQ0KT6O.rQwX1GgorUdQf0soLENosGFxpGauZW');
INSERT INTO users (email, password)VALUES ('bakugo@mha.com', '$2a$04$0SHV.WZphk9AEl0IKkeZxuMTUb1bjIScqd1zH35bd4v.Gmji84/Fq');


INSERT INTO wallets(user_id,balance) values(1,500000);
INSERT INTO wallets(user_id,balance) values(2,1500000);
INSERT INTO wallets(user_id,balance) values(3,2500000);
INSERT INTO wallets(user_id,balance) values(4,3500000);
INSERT INTO wallets(user_id,balance) values(5,4500000);

INSERT INTO funds(id,source_name) values (1,'Bank Transfer');
INSERT INTO funds(id,source_name) values (2,'Credit Card');
INSERT INTO funds(id,source_name) values (3,'Cash');

INSERT INTO transactions (wallet_number,transaction_type,source_id,fund_id,target_id,amount,created_at,description)
VALUES (270001,'Top Up',1,1,1,75000,(select timestamp '2020-10-10 20:00:00' +
       random() * (timestamp '2020-10-10 20:00:00' -
                   timestamp '2022-10-10 20:00:00'))
, 'Top Up via Bank Transfer' );
INSERT INTO transactions (wallet_number,transaction_type,source_id,fund_id,target_id,amount,created_at,description)
VALUES (270001,'Transfer',1,0,2,85000,(select timestamp '2020-10-10 20:00:00' +
       random() * (timestamp '2020-10-10 20:00:00' -
                   timestamp '2022-10-10 20:00:00'))
, 'Transfer Uang' );
INSERT INTO transactions (wallet_number,transaction_type,source_id,fund_id,target_id,amount,created_at,description)
VALUES (270001,'Transfer',1,0,3,65000,(select timestamp '2020-10-10 20:00:00' +
       random() * (timestamp '2020-10-10 20:00:00' -
                   timestamp '2022-10-10 20:00:00'))
, 'Transfer test' );

INSERT INTO transactions (wallet_number,transaction_type,source_id,fund_id,target_id,amount,created_at,description)
VALUES (270001,'Transfer',1,0,4,952000,(select timestamp '2020-10-10 20:00:00' +
       random() * (timestamp '2020-10-10 20:00:00' -
                   timestamp '2022-10-10 20:00:00'))
, 'Transfer test' );
INSERT INTO transactions (wallet_number,transaction_type,source_id,fund_id,target_id,amount,created_at,description)
VALUES (270001,'Transfer',1,0,5,110000,(select timestamp '2020-10-10 20:00:00' +
       random() * (timestamp '2020-10-10 20:00:00' -
                   timestamp '2022-10-10 20:00:00'))
, 'Transfer tes22' );
-------------------------------
INSERT INTO transactions (wallet_number,transaction_type,source_id,fund_id,target_id,amount,created_at,description)
VALUES (270002,'Transfer',2,0,5,75000,(select timestamp '2020-10-10 20:00:00' +
       random() * (timestamp '2020-10-10 20:00:00' -
                   timestamp '2022-10-10 20:00:00'))
, 'Transfer ke satu' );
INSERT INTO transactions (wallet_number,transaction_type,source_id,fund_id,target_id,amount,created_at,description)
VALUES (270002,'Top Up',2,2,2,85000,(select timestamp '2020-10-10 20:00:00' +
       random() * (timestamp '2020-10-10 20:00:00' -
                   timestamp '2022-10-10 20:00:00'))
, 'Top Up via Credit Card' );
INSERT INTO transactions (wallet_number,transaction_type,source_id,fund_id,target_id,amount,created_at,description)
VALUES (270002,'Transfer',2,0,3,65000,(select timestamp '2020-10-10 20:00:00' +
       random() * (timestamp '2020-10-10 20:00:00' -
                   timestamp '2022-10-10 20:00:00'))
, 'Transfer test' );

INSERT INTO transactions (wallet_number,transaction_type,source_id,fund_id,target_id,amount,created_at,description)
VALUES (270002,'Transfer',2,0,4,952000,(select timestamp '2020-10-10 20:00:00' +
       random() * (timestamp '2020-10-10 20:00:00' -
                   timestamp '2022-10-10 20:00:00'))
, 'Transfer test' );
-----
INSERT INTO transactions (wallet_number,transaction_type,source_id,fund_id,target_id,amount,created_at,description)
VALUES (270003,'Transfer',3,0,1,75000,(select timestamp '2020-10-10 20:00:00' +
       random() * (timestamp '2020-10-10 20:00:00' -
                   timestamp '2022-10-10 20:00:00'))
, 'Transfer ke satu' );
INSERT INTO transactions (wallet_number,transaction_type,source_id,fund_id,target_id,amount,created_at,description)
VALUES (270003,'Transfer',3,0,2,85000,(select timestamp '2020-10-10 20:00:00' +
       random() * (timestamp '2020-10-10 20:00:00' -
                   timestamp '2022-10-10 20:00:00'))
, 'Transfer' );
INSERT INTO transactions (wallet_number,transaction_type,source_id,fund_id,target_id,amount,created_at,description)
VALUES (270003,'Top Up',3,3,3,65000,(select timestamp '2020-10-10 20:00:00' +
       random() * (timestamp '2020-10-10 20:00:00' -
                   timestamp '2022-10-10 20:00:00'))
, 'Top Up via Cash' );

INSERT INTO transactions (wallet_number,transaction_type,source_id,fund_id,target_id,amount,created_at,description)
VALUES (270003,'Transfer',3,0,5,952000,(select timestamp '2020-10-10 20:00:00' +
       random() * (timestamp '2020-10-10 20:00:00' -
                   timestamp '2022-10-10 20:00:00'))
, 'Transfer test' );
------
INSERT INTO transactions (wallet_number,transaction_type,source_id,fund_id,target_id,amount,created_at,description)
VALUES (270004,'Transfer',4,0,1,75000,(select timestamp '2020-10-10 20:00:00' +
       random() * (timestamp '2020-10-10 20:00:00' -
                   timestamp '2022-10-10 20:00:00'))
, 'Transfer ke satu' );
INSERT INTO transactions (wallet_number,transaction_type,source_id,fund_id,target_id,amount,created_at,description)
VALUES (270004,'Transfer',4,0,2,85000,(select timestamp '2020-10-10 20:00:00' +
       random() * (timestamp '2020-10-10 20:00:00' -
                   timestamp '2022-10-10 20:00:00'))
, 'Transfer' );
INSERT INTO transactions (wallet_number,transaction_type,source_id,fund_id,target_id,amount,created_at,description)
VALUES (270004,'Top Up',4,2,4,65000,(select timestamp '2020-10-10 20:00:00' +
       random() * (timestamp '2020-10-10 20:00:00' -
                   timestamp '2022-10-10 20:00:00'))
, 'Top Up via Credit Card' );

INSERT INTO transactions (wallet_number,transaction_type,source_id,fund_id,target_id,amount,created_at,description)
VALUES (270004,'Transfer',4,0,5,952000,(select timestamp '2020-10-10 20:00:00' +
       random() * (timestamp '2020-10-10 20:00:00' -
                   timestamp '2022-10-10 20:00:00'))
, 'Transfer test' );
-----
INSERT INTO transactions (wallet_number,transaction_type,source_id,fund_id,target_id,amount,created_at,description)
VALUES (270005,'Transfer',5,0,2,75000,(select timestamp '2020-10-10 20:00:00' +
       random() * (timestamp '2020-10-10 20:00:00' -
                   timestamp '2022-10-10 20:00:00'))
, 'Transfer ke satu' );
INSERT INTO transactions (wallet_number,transaction_type,source_id,fund_id,target_id,amount,created_at,description)
VALUES (270005,'Transfer',5,0,3,85000,(select timestamp '2020-10-10 20:00:00' +
       random() * (timestamp '2020-10-10 20:00:00' -
                   timestamp '2022-10-10 20:00:00'))
, 'Transfer' );
INSERT INTO transactions (wallet_number,transaction_type,source_id,fund_id,target_id,amount,created_at,description)
VALUES (270005,'Top Up',5,1,5,65000,(select timestamp '2020-10-10 20:00:00' +
       random() * (timestamp '2020-10-10 20:00:00' -
                   timestamp '2022-10-10 20:00:00'))
, 'Top Up via Bnk' );


select * from users
select * from wallets
select * from transactions
select * from funds;


ALTER TABLE public.users ADD CONSTRAINT users_un UNIQUE (email);

