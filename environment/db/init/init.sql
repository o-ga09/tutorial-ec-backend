CREATE USER 'repl'@'api-dbsrv01' IDENTIFIED BY 'repl';
GRANT REPLICATION SLAVE ON *.* TO 'repl'@'api-dbsrv01';

CREATE TABLE `orders` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `updated_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  `deleted_at` datetime(6) DEFAULT NULL,
  `user_id` varchar(255) NOT NULL,
  `total_amount` int(11) NOT NULL,
  `order_at` datetime(6) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `order_products` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `updated_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  `deleted_at` datetime(6) DEFAULT NULL,
  `order_id` varchar(255) NOT NULL,
  `product_id` varchar(255) NOT NULL,
  `price` int(11) NOT NULL,
  `quantity` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `products` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `updated_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  `deleted_at` datetime(6) DEFAULT NULL,
  `product_id` varchar(255) NOT NULL,
  `owner_id` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `price` int(11) NOT NULL,
  `stock` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `updated_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  `deleted_at` datetime(6) DEFAULT NULL,
  `user_id` varchar(255) NOT NULL,
  `email` varchar(255) DEFAULT NULL,
  `phonenumber` varchar(255) DEFAULT NULL,
  `last_name` varchar(255) DEFAULT NULL,
  `first_name` varchar(255) DEFAULT NULL,
  `prefecture` varchar(255) DEFAULT NULL,
  `city` varchar(255) DEFAULT NULL,
  `extra` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `owners` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `updated_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  `deleted_at` datetime(6) DEFAULT NULL,
  `owner_id` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO orders (`user_id`, `total_amount`, `order_at`) VALUES
  ('user-1', 10000, '2023-11-14 12:00:00'),
  ('user-2', 20000, '2023-11-15 13:00:00'),
  ('user-3', 30000, '2023-11-16 14:00:00');

INSERT INTO order_products (`order_id`, `product_id`, `price`, `quantity`) VALUES
  ('order-1', 'product-1', 1000, 1),
  ('order-1', 'product-2', 2000, 2),
  ('order-2', 'product-3', 3000, 3),
  ('order-3', 'product-4', 4000, 4);

INSERT INTO products (`product_id`, `owner_id`, `name`, `description`, `price`, `stock`) VALUES
  ('product-1', 'owner-1', '商品1', '商品の説明1', 1000, 10),
  ('product-2', 'owner-2', '商品2', '商品の説明2', 2000, 20),
  ('product-3', 'owner-3', '商品3', '商品の説明3', 3000, 30),
  ('product-4', 'owner-4', '商品4', '商品の説明4', 4000, 40);

INSERT INTO users (`user_id`, `email`, `phonenumber`, `last_name`, `first_name`, `prefecture`, `city`, `extra`) VALUES
  ('user-1', 'user1@example.com', '090-1234-5678', '山田', '太郎', '東京', '渋谷区', '道玄坂'),
  ('user-2', 'user2@example.com', '090-8765-4321', '佐藤', '花子', '大阪', '大阪市', '難波'),
  ('user-3', 'user3@example.com', '090-1122-3344', '斎藤', '健', '神奈川', '横浜市', '港北');

INSERT INTO owners (`owner_id`, `email`, `name`) VALUES
  ('owner-1', 'owner1@example.com', '店舗1'),
  ('owner-2', 'owner2@example.com', '店舗2'),
  ('owner-3', 'owner3@example.com', '店舗3'),
  ('owner-4', 'owner4@example.com', '店舗4');
