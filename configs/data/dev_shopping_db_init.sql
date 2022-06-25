USE dev;

create table product
(
    id    mediumtext   null,
    name  varchar(128) null,
    price float        null
);

INSERT INTO dev.product (id, name, price) VALUES ('1', 'shoes', 120.0);
INSERT INTO dev.product (id, name, price) VALUES ('2', 'pen', 80.0);
INSERT INTO dev.product (id, name, price) VALUES ('3', 't-shirt', 9.15);

create table advertisement
(
    id      int  null,
    content text null
);

INSERT INTO dev.advertisement (id, content) VALUES (1, 'Look, this shirt is cheap.');
INSERT INTO dev.advertisement (id, content) VALUES (2, 'Good sale, xxx...');
INSERT INTO dev.advertisement (id, content) VALUES (3, 'Can you imagine the ...');


create table product_detail
(
    id           int          null,
    name         varchar(128) null,
    product_type varchar(128) null,
    price        float        null,
    picture_uri  text         null
);

INSERT INTO dev.product_detail (id, name, product_type, price, picture_uri) VALUES (1, 'shoes', 'clothes', 120, null);
INSERT INTO dev.product_detail (id, name, product_type, price, picture_uri) VALUES (2, 'pen', 'usage', 80, null);
INSERT INTO dev.product_detail (id, name, product_type, price, picture_uri) VALUES (3, 't-shirt', 'clothes', 9.15, null);
