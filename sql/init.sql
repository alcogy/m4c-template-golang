CREATE TABLE customer (
  id serial primary key,
  name varchar(128) not null default '',
  zip varchar(8) not null default '',
  address varchar(255) not null default '',
  tel varchar(13) not null default '',
  email varchar(255) not null default ''
);

INSERT INTO customer VALUES (
  DEFAULT,
  'Alcogy Inc',
  '060-0032',
  'Sun Mountain Bldg 3F 1-3-3 Kita 2 Jo Higashi Chuo-ku Sapporo-shi Hokkaido',
  '080-6332-4661',
  'info@alcogy.com'
);
