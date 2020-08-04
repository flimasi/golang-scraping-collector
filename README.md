## golang-scraping-collector
Golang Scraping Store

### Pré require install go in system

### Create Database

CREATE SCHEMA `scraping` ;

CREATE TABLE `scraping`.`scraping` (
  `idscraping` INT NOT NULL AUTO_INCREMENT,
  `content` TEXT NULL,
  `source` VARCHAR(45) NULL,
  `date` DATETIME NULL,
  PRIMARY KEY (`idscraping`));

### Set Sources


### Collect Data

./collect.go