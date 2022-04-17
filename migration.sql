CREATE DATABASE IF NOT EXISTS `vuln_db`;

CREATE TABLE vuln_db.users (
                               user_id BIGINT auto_increment NOT NULL,
                               name varchar(100) NULL,
                               email varchar(100) NULL,
                               password varchar(100) NULL,
                               is_active BOOL DEFAULT true NOT NULL,
                               CONSTRAINT users_PK PRIMARY KEY (user_id)
)
    ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE vuln_db.products (
                                  product_id BIGINT auto_increment NOT NULL,
                                  product_name varchar(100) NULL,
                                  product_desc TEXT NULL,
                                  price INT NOT NULL,
                                  CONSTRAINT products_PK PRIMARY KEY (product_id)
)
    ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;
