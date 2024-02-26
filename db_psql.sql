create table news
(
	id text not null,
	title varchar(500) not null,
	news_image varchar(200) not null,
	news_text text not null,
	created_at timestamp not null default current_timestamp,
	updated_at timestamp not null,
	primary key(id)
);
 
create index news_title_index on news(title)

alter table news

ALTER TABLE admin
DROP CONSTRAINT fk_admin_news CASCADE;

alter table news rename to blogs

alter table news
	add column id_category text
alter table news
	add constraint fk_news_category foreign key(id_category) references categories(id)
alter table news
	add primary key(id)

alter table news
	drop column id_admin

create table admin
(
	id text not null,
	username varchar(40) not null,
	password text not null,
	email varchar(150) not null,
	created_at timestamp not null default current_timestamp,
	primary key(id)
);

alter table admin
	add constraint admin_email_unique unique(email)

create table categories
(
	id text not null,
	category varchar(50),
	primary key(id)
)

insert into categories(id, category, created_at, updated_at)
	values
		('C0001C', 'Politik', 111, 111),
		('C0002C', 'Teknologi', 111, 111),
		('C0003C', 'Perang', 111, 111)
select * from categories
create index categories_category_index on categories(category)
	

insert into admin(id, username, password, email)
	values
		('A0001A', 'farid anang samudra', 'farid123', 'faridanangs@gmail.com'),
		('A0002A', 'raika aisyah', 'raika123', 'raika@gmail.com'),
		('A0003A', 'wagas try', 'wagas123', 'wagas@gmail.com'),
		('A0004A', 'diana milda', 'diana123', 'diana@gmail.com');

alter table admin
	add column id_news text;
alter table admin
	add constraint fk_admin_news foreign key(id_news) references news(id)
select * from admin;

insert into blogs(id, title, news_image, news_text, updated_at,id_category, id_admin, created_at)
	values
		('N0001N', 'perang rusia ukraina', 'http://example.com', 'lorem ispum dolor sit amet lolipop', 111111, 'C0003C', '3314fce3-5800-44e6-ba5b-7ce9847f4c9b', 111111),
		('N0002N', 'kecanggihan ai', 'http://example.com', 'lorem ispum dolor sit amet lolipop', 111111,'C0002C', '7fa88835-90e3-4fbc-9b81-9bb62f734283', 111111),
		('N0003N', 'politik dinasti', 'http://example.com', 'lorem ispum dolor sit amet lolipop', 111111, 'C0001C', '0954ddbf-d053-4a80-a44d-d9386faea75b', 111111),
		('N0004N', 'prabowo gibran', 'http://example.com', 'lorem ispum dolor sit amet lolipop', 111111, 'C0001C', '0954ddbf-d053-4a80-a44d-d9386faea75b', 111111),
		('N0005N', 'anis muhaimin', 'http://example.com', 'lorem ispum dolor sit amet lolipop', 111111,'C0001C', '0954ddbf-d053-4a80-a44d-d9386faea75b', 111111),
		('N0006N', 'perang palestina israil', 'http://example.com', 'lorem ispum dolor sit amet lolipop', 111111,'C0003C', '3314fce3-5800-44e6-ba5b-7ce9847f4c9b', 111111),
		('N0007N', 'perang dunia 2', 'http://example.com', 'lorem ispum dolor sit amet lolipop', 111111, 'C0003C', '3314fce3-5800-44e6-ba5b-7ce9847f4c9b', 111111),
		('N0008N', 'laptop terbaru yang bisa terbang', 'http://example.com', 'lorem ispum dolor sit amet lolipop', 111111, 'C0002C', '7fa88835-90e3-4fbc-9b81-9bb62f734283', 111111),
		('N0009N', 'hape terbaru yang bisa hilang', 'http://example.com', 'lorem ispum dolor sit amet lolipop', 111111, 'C0002C', '7fa88835-90e3-4fbc-9b81-9bb62f734283', 111111);

select * from blogs;
select * from categories;
select * from admin;

alter table admin
	drop constraint fk_admin_news;
	

select * from blogs
	join categories as c on c.id = news.id_category
	
alter table news
	add column id text not null;

alter table news
	add column id_admin text;

alter table news
	add constraint fk_news_admin foreign key(id_admin) references admin(id)
		
update from news set
	id_admin = 'A0001A'
	where

delete from news where id_category in('C0003C', 'C0001C', 'C0002C')



select A.id, A.username, A.email, N.title, N.news_image, N.news_text, C.category from admin as A
	join blogs as N on A.id = N.id_admin
	join categories as C on C.id = N.id_category



create table blog_category(
	id serial not null,
	id_category text not null,
	id_blog text not null,
	primary key(id),
	constraint fk_blog_category_id_category foreign key(id_category) references categories(id),
	constraint fk_blog_category_id_blog foreign key(id_blog) references blogs(id),

)

alter table admin
	add column updated_at timestamp

select * from admin
select * from blogs

