/*
 Navicat MySQL Data Transfer

 Source Server         : farm
 Source Server Type    : MySQL
 Source Server Version : 80016
 Source Host           : localhost:3306
 Source Schema         : farm

 Target Server Type    : MySQL
 Target Server Version : 80016
 File Encoding         : 65001

 Date: 16/06/2019 15:02:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for farm_dhi_record
-- ----------------------------
DROP TABLE IF EXISTS `farm_dhi_record`;
CREATE TABLE `farm_dhi_record` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cow_id` int(11) NOT NULL,
  `record_time` varchar(50) NOT NULL,
  `birth_time` varchar(50) NOT NULL,
  `day_age` int(11) NOT NULL,
  `count` int(5) NOT NULL,
  `milk_production` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
