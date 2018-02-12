
SET NAMES utf8;
SET time_zone = '+00:00';

DROP DATABASE IF EXISTS `cybergame`;
CREATE DATABASE `cybergame` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `cybergame`;

DROP TABLE IF EXISTS `teams`;
CREATE TABLE `teams` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sid` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `description` TEXT,
  `deleted` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `players`;
CREATE TABLE `players` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sid` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `deleted` int(11) NOT NULL,
  `description` TEXT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `teamplayers`;
CREATE TABLE `teamplayers` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sid` int(11) NOT NULL,
  `player_id` int(11) NOT NULL,
  `team_id` int(11) NOT NULL,
  `fromtime` TIMESTAMP NULL DEFAULT NULL,
  `totime` TIMESTAMP NULL DEFAULT NULL,
  `deleted` TEXT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sid` int(11) NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(50) NOT NULL,
  `facebook_url` varchar(100) NOT NULL,
  `twitch_url` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tours`;
CREATE TABLE `tours` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sid` int(11) NOT NULL,
  `name` int(11) NOT NULL,
  `description` TEXT,
  `title` varchar(100) NOT NULL,
  `location` varchar(50) NOT NULL,
  `series` varchar(50) NOT NULL,
  `prize` varchar(50) NOT NULL,
  `fromtime` TIMESTAMP NULL DEFAULT NULL,
  `totime` TIMESTAMP NULL DEFAULT NULL,
  `deleted` TEXT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `matches`;
CREATE TABLE `matches` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sid` int(11) NOT NULL,
  `name` int(11) NOT NULL,
  `team_id1` int(11) NOT NULL,
  `team_id2` int(11) NOT NULL,
  `score_id1` int(11) NOT NULL,
  `score_id2` int(11) NOT NULL,
  `tour_id` int(11) NOT NULL,
  
  
  `description` TEXT,
  `location` varchar(50) NOT NULL,
  `series` varchar(50) NOT NULL,
  `prize` varchar(50) NOT NULL,
  `time` TIMESTAMP NULL DEFAULT NULL,
  `totime` TIMESTAMP NULL DEFAULT NULL,
  `deleted` TEXT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `stages`;
CREATE TABLE `stages` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sid` int(11) NOT NULL,
  `name` int(11) NOT NULL,
  `team_id1` int(11) NOT NULL,
  `team_id2` int(11) NOT NULL,
  `score_id1` int(11) NOT NULL,
  `score_id2` int(11) NOT NULL,
  `tour_id` int(11) NOT NULL,
  
  
  `description` TEXT,
  `location` varchar(50) NOT NULL,
  `series` varchar(50) NOT NULL,
  `prize` varchar(50) NOT NULL,
  `time` TIMESTAMP NULL DEFAULT NULL,
  `totime` TIMESTAMP NULL DEFAULT NULL,
  `deleted` TEXT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `seasons`;
CREATE TABLE `seasons` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sid` int(11) NOT NULL,
  `name` int(11) NOT NULL,
  `team_id1` int(11) NOT NULL,
  `team_id2` int(11) NOT NULL,
  `score_id1` int(11) NOT NULL,
  `score_id2` int(11) NOT NULL,
  `tour_id` int(11) NOT NULL,
  
  
  `description` TEXT,
  `location` varchar(50) NOT NULL,
  `series` varchar(50) NOT NULL,
  `prize` varchar(50) NOT NULL,
  `time` TIMESTAMP NULL DEFAULT NULL,
  `totime` TIMESTAMP NULL DEFAULT NULL,
  `deleted` TEXT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


