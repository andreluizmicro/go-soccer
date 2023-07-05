-- go_soccer.countries definition

CREATE TABLE `countries` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` varchar(100) NOT NULL,
  `name` varchar(100) NOT NULL,
  `capital` varchar(100) NOT NULL,
  `area` varchar(100) NOT NULL,
  `language` varchar(100) NOT NULL,
  `timezone` varchar(100) NOT NULL,
  `continent` varchar(100) NOT NULL,
  `official_color` varchar(100) NOT NULL,
  `population` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;