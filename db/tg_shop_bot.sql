-- MySQL Script generated by MySQL Workbench
-- Tue Dec 31 15:25:48 2024
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema tg_shop
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema tg_shop
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `tg_shop` DEFAULT CHARACTER SET utf8 ;
USE `tg_shop` ;

-- -----------------------------------------------------
-- Table `tg_shop`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tg_shop`.`users` (
  `userId` INT NOT NULL,
  `userSecret` VARCHAR(45) NULL,
  `userOrders` INT NULL,
  `userOrders_orderId` INT NOT NULL,
  PRIMARY KEY (`userId`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tg_shop`.`vendorItems`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tg_shop`.`vendorItems` (
  `itemId` INT NOT NULL,
  `vendorId` INT NOT NULL,
  `itemPrice` FLOAT NULL,
  `itemName` VARCHAR(45) NULL,
  `itemDescription` VARCHAR(200) NULL,
  PRIMARY KEY (`itemId`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tg_shop`.`vendorCategories`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tg_shop`.`vendorCategories` (
  `categoryId` INT NOT NULL,
  `vendorId` INT NOT NULL,
  `categoryName` VARCHAR(45) NULL,
  `vendorItems_itemId` INT NOT NULL,
  PRIMARY KEY (`categoryId`),
  INDEX `fk_vendorCategories_vendorItems1_idx` (`vendorItems_itemId` ASC) VISIBLE,
  CONSTRAINT `fk_vendorCategories_vendorItems1`
    FOREIGN KEY (`vendorItems_itemId`)
    REFERENCES `tg_shop`.`vendorItems` (`itemId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tg_shop`.`orders`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tg_shop`.`orders` (
  `orderId` INT NOT NULL,
  `userId` INT NULL,
  `itemId` INT NULL,
  `time` TIMESTAMP NULL,
  `users_userId` INT NOT NULL,
  `vendorId` INT NULL,
  `categoryId` INT NULL,
  `vendorCategories_categoryId` INT NOT NULL,
  `vendorItems_itemId` INT NOT NULL,
  PRIMARY KEY (`orderId`),
  INDEX `fk_vendorOrders_users1_idx` (`users_userId` ASC) VISIBLE,
  INDEX `fk_orders_vendorCategories1_idx` (`vendorCategories_categoryId` ASC) VISIBLE,
  INDEX `fk_orders_vendorItems1_idx` (`vendorItems_itemId` ASC) VISIBLE,
  CONSTRAINT `fk_vendorOrders_users1`
    FOREIGN KEY (`users_userId`)
    REFERENCES `tg_shop`.`users` (`userId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_orders_vendorCategories1`
    FOREIGN KEY (`vendorCategories_categoryId`)
    REFERENCES `tg_shop`.`vendorCategories` (`categoryId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_orders_vendorItems1`
    FOREIGN KEY (`vendorItems_itemId`)
    REFERENCES `tg_shop`.`vendorItems` (`itemId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tg_shop`.`vendors`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tg_shop`.`vendors` (
  `vendorId` INT NOT NULL,
  `vendorName` VARCHAR(45) NULL,
  `vendorOrders_vendorOrderId` INT NOT NULL,
  `vendorCategories_categoryId` INT NOT NULL,
  `vendorItems_itemId` INT NOT NULL,
  PRIMARY KEY (`vendorId`),
  INDEX `fk_vendors_vendorOrders1_idx` (`vendorOrders_vendorOrderId` ASC) VISIBLE,
  INDEX `fk_vendors_vendorCategories1_idx` (`vendorCategories_categoryId` ASC) VISIBLE,
  INDEX `fk_vendors_vendorItems1_idx` (`vendorItems_itemId` ASC) VISIBLE,
  CONSTRAINT `fk_vendors_vendorOrders1`
    FOREIGN KEY (`vendorOrders_vendorOrderId`)
    REFERENCES `tg_shop`.`orders` (`orderId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_vendors_vendorCategories1`
    FOREIGN KEY (`vendorCategories_categoryId`)
    REFERENCES `tg_shop`.`vendorCategories` (`categoryId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_vendors_vendorItems1`
    FOREIGN KEY (`vendorItems_itemId`)
    REFERENCES `tg_shop`.`vendorItems` (`itemId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;