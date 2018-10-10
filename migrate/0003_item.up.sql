SET FOREIGN_KEY_CHECKS = 0;
CREATE TABLE IF NOT EXISTS `item` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `order_id` BIGINT UNSIGNED NOT NULL,
  `item_name` VARCHAR(45) NOT NULL,
  `tax_code_id` BIGINT UNSIGNED NOT NULL,
  `price` DECIMAL(20,2) UNSIGNED NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  INDEX `fk_item_1_idx` (`order_id` ASC),
  INDEX `fk_item_2_idx` (`tax_code_id` ASC),
  CONSTRAINT `fk_item_1`
    FOREIGN KEY (`order_id`)
    REFERENCES `order` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_item_2`
    FOREIGN KEY (`tax_code_id`)
    REFERENCES `tax_code` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
