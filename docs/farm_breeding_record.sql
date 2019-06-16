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

 Date: 16/06/2019 15:02:28
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for farm_breeding_record
-- ----------------------------
DROP TABLE IF EXISTS `farm_breeding_record`;
CREATE TABLE `farm_breeding_record` (
  `cow_id` int(50) NOT NULL,
  `last_birth_time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `start_time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `end_time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `first_time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `first_number` int(50) NOT NULL,
  `second_time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `second_number` int(50) NOT NULL,
  `third_time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `third_number` int(50) NOT NULL,
  `fourth_time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `fourth_number` int(50) NOT NULL,
  `final_time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `final_number` int(50) NOT NULL,
  `id` int(50) unsigned zerofill NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of farm_breeding_record
-- ----------------------------
BEGIN;
INSERT INTO `farm_breeding_record` VALUES (7, '2019-04-21 14:46:43', '2019-04-22 14:46:43', '2019-05-22 14:46:43', '2019-04-22 14:56:43', 1, '2019-04-23 14:46:43', 2, '2019-04-24 14:46:43', 3, '2019-04-25 14:46:43', 4, '2019-04-26 14:46:43', 5, 00000000000000000000000000000000000000000000000003);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
