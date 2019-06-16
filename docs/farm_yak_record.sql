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

 Date: 16/06/2019 15:03:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for farm_yak_record
-- ----------------------------
DROP TABLE IF EXISTS `farm_yak_record`;
CREATE TABLE `farm_yak_record` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `yak_id` int(11) NOT NULL,
  `variety` varchar(50) NOT NULL,
  `birth_time` varchar(50) NOT NULL,
  `father_number` int(11) NOT NULL,
  `mother_number` int(11) NOT NULL,
  `weight` varchar(50) NOT NULL,
  `length` varchar(50) NOT NULL,
  `height` varchar(50) NOT NULL,
  `bust` varchar(50) NOT NULL,
  `front_rear_distance` varchar(50) NOT NULL,
  `left_right_distance` varchar(50) NOT NULL,
  `left_front_length` varchar(50) NOT NULL,
  `right_rear_length` varchar(50) NOT NULL,
  `record_time` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of farm_yak_record
-- ----------------------------
BEGIN;
INSERT INTO `farm_yak_record` VALUES (1, 11, 'bobo', '2019-04-21 14:46:43', 5, 7, '77.7', '88.8', '99.9', '101.1', '33.3', '44.4', '55.5', '66.6', '2019-06-15 18:06:58');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
