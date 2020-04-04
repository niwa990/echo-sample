CREATE TABLE `comments` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(200) COLLATE utf8_unicode_ci DEFAULT 'NONAME',
  `text` text COLLATE utf8_unicode_ci,
  `created_at` datetime NOT NULL default current_timestamp,
  `updated_at` datetime NOT NULL default current_timestamp on update current_timestamp,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
