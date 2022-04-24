CREATE DATABASE IF NOT EXISTS pismo;

CREATE TABLE IF NOT EXISTS `pismo`.`accounts` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `document_number` VARCHAR(45) NOT NULL,
    PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE);

CREATE TABLE IF NOT EXISTS `pismo`.`operations_types` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `description` VARCHAR(45) NOT NULL,
    PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE);

CREATE TABLE IF NOT EXISTS `pismo`.`transactions` (
  `transaction_id` INT NOT NULL AUTO_INCREMENT,
  `account_id` INT NOT NULL,
  `operation_type_id` INT NOT NULL,
  `amount` FLOAT NOT NULL,
  `event_date` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`transaction_id`),
  UNIQUE INDEX `transaction_id_UNIQUE` (`transaction_id` ASC) VISIBLE,
  INDEX `account_fk_idx` (`account_id` ASC) VISIBLE,
  INDEX `operation_type_fk_idx` (`operation_type_id` ASC) VISIBLE,
  CONSTRAINT `account_fk`
    FOREIGN KEY (`account_id`)
    REFERENCES `pismo`.`accounts` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `operation_type_fk`
    FOREIGN KEY (`operation_type_id`)
    REFERENCES `pismo`.`operations_types` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);

INSERT INTO `pismo`.`operations_types`(`description`) VALUES ('COMPRA A VISTA');    
INSERT INTO `pismo`.`operations_types`(`description`) VALUES ('COMPRA PARCELADA');
INSERT INTO `pismo`.`operations_types`(`description`) VALUES ('SAQUE');
INSERT INTO `pismo`.`operations_types`(`description`) VALUES ('PAGAMENTO');