create database online_learning_db;

create type user_level as enum ('newbie', 'junior', 'senior', 'master');

create table users (
	id SERIAL primary key not null,
	email VARCHAR(256) unique not null,
	password VARCHAR(256) not null,
	is_admin BOOLEAN not null,
	fullname VARCHAR(2566) not null,
	address VARCHAR(256) not null,
	phone_number VARCHAR(16) not null,
	level user_level not null,
	referral VARCHAR(256) not null,
	ref_referral VARCHAR(256)
);

create table categories (
	id SERIAL primary key not null,
	name VARCHAR(256) not null
);

create table tags (
	id SERIAL primary key not null,
	name VARCHAR(256) not null
);

create table courses (
	id SERIAL primary key not null,
	title VARCHAR(256) not null,
	slug VARCHAR(256) not null,
	summary_description VARCHAR(256) not null,
	content VARCHAR(256) not null,
	img_thumbnail VARCHAR(256) not null,
	img_url VARCHAR(256) not null,
	author_name VARCHAR(256) not null,
	category_id INTEGER not null,
	tag_id INTEGER not null,
	created_at TIMESTAMP not null default CURRENT_TIMESTAMP,
    updated_at TIMESTAMP not null default CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP null,
	foreign key (category_id) references categories (id),
	foreign key (tag_id) references tags (id)
);

create table gifts (
	id SERIAL primary key not null,
	name VARCHAR(256) not null,
	stock INTEGER not null
);

create table vouchers(
	id SERIAL primary key not null,
	name VARCHAR(256) not null,
	voucher_code VARCHAR(256) not null,
	benefit DECIMAL not null
);

create type course_status as enum ('completed', 'on progress');

create table user_courses(
	id SERIAL primary key not null,
	user_id INTEGER not null,
	course_id INTEGER not null,
	status course_status not null,
	foreign key (user_id) references users (id),
	foreign key (course_id) references courses (id)
);

create table favorites(
	id SERIAL primary key not null,
	user_id INTEGER not null,
	course_id INTEGER not null,
	added_date DATE not null,
	foreign key (user_id) references users (id),
	foreign key (course_id) references courses (id)
);

create type track_status as enum ('completed', 'process', 'canceled');

create table tracks(
	id SERIAL primary key not null,
	user_id INTEGER not null,
	status track_status not null,
	name VARCHAR(256) not null,
	total INTEGER not null,
	generated_date TIMESTAMP not null,
	estimated_date DATE not null,
	generated_by VARCHAR(256) not null,
	foreign key (user_id) references users (id)
);

create table user_gifts(
	id SERIAL primary key not null,
	user_id INTEGER not null,
	gift_id INTEGER not null,
	tracking_id INTEGER not null,
	foreign key (user_id) references users (id),
	foreign key (gift_id) references gifts (id),
	foreign key (tracking_id) references tracks (id)
);


create type voucher_status as enum ('used', 'unused');

create table user_vouchers(
	id SERIAL primary key not null,
	user_id INTEGER not null,
	voucher_id INTEGER not null,
	status voucher_status not null,
	expired_date TIMESTAMP not null,
	foreign key (user_id) references users (id),
	foreign key (voucher_id) references vouchers (id)
);

create type invoice_status as enum ('completed', 'process', 'canceled');

create table invoices(
	id SERIAL primary key not null,
	user_id INTEGER not null,
	voucher_id INTEGER,
	status invoice_status not null,
	total numeric not null,
	payment_date TIMESTAMP not null,
	benefit_discount DECIMAL,
	foreign key (user_id) references users (id),
	foreign key (voucher_id) references vouchers (id)
);

create table transactions(
	id SERIAL primary key not null,
	invoice_id INTEGER not null,
	course_id INTEGER not null,
	sold_price numeric not null,
	foreign key (invoice_id) references invoices (id),
	foreign key (course_id) references courses (id)
);



