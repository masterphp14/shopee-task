SET FOREIGN_KEY_CHECKS = 0;
CREATE TABLE IF NOT EXISTS `tax_code` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `code` VARCHAR(45) NOT NULL,
  `refundable` ENUM('yes', 'no') NULL,
  `type` VARCHAR(45) NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;

INSERT INTO `tax_code` (`id`, `code`, `refundable`, `type`) VALUES
('1', '1', 'yes', 'Food & Beverage'),
('2', '2', 'no', 'Tobacco'),
('3', '3', 'no', 'Entertainment');
