CREATE TABLE `auths` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `email` varchar(255) UNIQUE NOT NULL,
  `password` varchar(255),
  `role` varchar(255) DEFAULT "user",
  `login_at` datetime
);

CREATE TABLE `users` (
  `auth_id` int PRIMARY KEY,
  `name` varchar(255),
  `gender` varchar(255),
  `phone_number` varchar(255),
  `pict_url` varchar(255),
  `created_at` datetime,
  `updated_at` timestamp,
  `deleted_at` datetime
);

CREATE TABLE `addresses` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `street` varchar(255),
  `city_id` varchar(255),
  `province_id` varchar(255),
  `address_tag` varchar(255),
  `activate` varchar(1) DEFAULT "n",
  `created_at` datetime,
  `updated_at` timestamp,
  `deleted_at` datetime
);

CREATE TABLE `merchants` (
  `auth_id` int PRIMARY KEY,
  `name` varchar(255),
  `phone_number` varchar(255),
  `street` varchar(255),
  `city_id` varchar(255),
  `province_id` varchar(255),
  `pict_url` varchar(255),
  `created_at` datetime,
  `updated_at` timestamp,
  `deleted_at` datetime
);

CREATE TABLE `products` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `merchant_id` int,
  `category` varchar(255),
  `name` varchar(255),
  `desc` varchar(255),
  `price` int,
  `stock` int,
  `weight` int,
  `img_url` varchar(255),
  `created_at` datetime,
  `updated_at` timestamp,
  `deleted_at` datetime
);

CREATE TABLE `orders` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `total_price` int,
  `total_weight` int,
  `status` varchar(255),
  `created_at` datetime,
  `updated_at` timestamp,
  `deleted_at` datetime
);

CREATE TABLE `order_items` (
  `order_id` int,
  `product_id` int,
  `quantity` int,
  `notes` varchar(255),
  `status` varchar(255),
  PRIMARY KEY (`order_id`, `product_id`)
);

CREATE TABLE `reviews` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `order_id` int,
  `product_id` int,
  `rating` int,
  `notes` tex,
  `img_url` varchar(255),
  `created_at` timestamp
);

CREATE TABLE `merchant_tags` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255)
);

CREATE TABLE `tags` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `merchant_id` int,
  `merchant_tag_id` int
);

ALTER TABLE `users` ADD FOREIGN KEY (`auth_id`) REFERENCES `auths` (`id`);

ALTER TABLE `merchants` ADD FOREIGN KEY (`auth_id`) REFERENCES `auths` (`id`);

ALTER TABLE `addresses` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`auth_id`);

ALTER TABLE `orders` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`auth_id`);

ALTER TABLE `order_items` ADD FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`);

ALTER TABLE `reviews` ADD FOREIGN KEY (`order_id`, `product_id`) REFERENCES `order_items` (`order_id`, `product_id`);

ALTER TABLE `order_items` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);

ALTER TABLE `products` ADD FOREIGN KEY (`merchant_id`) REFERENCES `merchants` (`auth_id`);

ALTER TABLE `tags` ADD FOREIGN KEY (`merchant_tag_id`) REFERENCES `merchant_tags` (`id`);

ALTER TABLE `tags` ADD FOREIGN KEY (`merchant_id`) REFERENCES `merchants` (`auth_id`);
