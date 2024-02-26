create table admins
(
	id text not null,
	username varchar(30) not null,
	password text not null,
	primary key(id)
)

insert into admins(id, username, password)
	values
		('A001', 'farid', 'far');

create table blogs
(
	id text not null,
	title varchar(300),
	text_blog text not null,
	admin_id text not null,
	primary key(id),
	constraint fk_blogs_admin_id foreign key(admin_id) references admins(id)
)
create table categories
(
	id serial not null,
	category varchar(50),
	primary key(id)
)

insert into categories(category)
	values('buku'),('handpon'),('polpen')

select * from categories

create table blogs_categories
(
	id serial not null,
	category_id int not null,
	blog_id text not null,
	primary key(id),
	constraint fk_category_id foreign key(category_id) references categories(id),
	constraint fk_blog_id foreign key(blog_id) references blogs(id)
	
)

drop table blogs_categories

select * from blogs_categories
select * from blogs
select * from admins
delete from blogs where admin_id = 'A001'
delete from categories where category = 'sepeda'
delete from admins where username = 'kangkung'