
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE users (
  `id` int(11) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`),
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE users;
