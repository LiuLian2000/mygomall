CREATE DATABASE IF NOT EXISTS `user`  
   DEFAULT CHARACTER SET = 'utf8mb4';;
USE `user`;
CREATE TABLE IF NOT EXISTS `user` (
  `id` bigint NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `email` varchar(255) NOT NULL,
  `password_hashed` text NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci


#product service database
#一个商品可能有多个类别(category)，product表用json存储所有类别的name，用类别的name从catagory表中可以查到catgory描述等其他信息
#商品id用雪花算法生成int64类型的数
CREATE DATABASE IF NOT EXISTS `product`  
   DEFAULT CHARACTER SET = 'utf8mb4';
USE `product`;
#商品表
CREATE TABLE IF NOT EXISTS `product` (
  `id` BIGINT NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `name` VARCHAR(255) NOT NULL,
  `description` TEXT,
  `picture` VARCHAR(255),
  `price` FLOAT NOT NULL,
  `categories` JSON,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
#商品类别表

CREATE TABLE IF NOT EXISTS `category` (
    `id` BIGINT NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `name` VARCHAR(255) NOT NULL UNIQUE,
    `description` TEXT,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

