create table as_category_dishes
(
  id int not null auto_increment
    primary key,
  category_name varchar(255) not null,
  dish_create_time datetime not null,
  dish_summary text not null,
  dish_modify_time datetime not null,
  constraint category_name
  unique (category_name)
)
;

create table as_dishes
(
  id int not null auto_increment
    primary key,
  dish_name varchar(255) not null,
  dish_price float not null,
  dish_unit varchar(255) not null,
  dish_description mediumtext null,
  dish_create_time datetime not null,
  dish_modify_time datetime not null,
  dish_category_id int not null,
  constraint dish_name
  unique (dish_name),
  constraint as_dishes_ibfk_1
  foreign key (dish_category_id) references app_demo.as_category_dishes (id)
)
;

create index dish_category_id
  on as_dishes (dish_category_id)
;

create table user_info
(
  id int not null auto_increment
    primary key,
  user_name varchar(255) not null,
  password varchar(255) not null,
  email varchar(255) not null,
  constraint user_name
  unique (user_name),
  constraint email
  unique (email)
)
;

