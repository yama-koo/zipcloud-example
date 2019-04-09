create table zipcloud.zipcode (
  id int auto_increment not null primary key,
  prefecture_code int,
  zipcode varchar(50),
  prefecture varchar(50),
  municipalities varchar(50),
  section varchar(50),
  prefecture_phonetic varchar(50),
  municipalities_phonetic varchar(50),
  section_phonetic varchar(50)
);
