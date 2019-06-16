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

 Date: 16/06/2019 15:02:57
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for farm_milk_production_record
-- ----------------------------
DROP TABLE IF EXISTS `farm_milk_production_record`;
CREATE TABLE `farm_milk_production_record` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cow_id` int(11) NOT NULL,
  `record_time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `morn_production` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `noon_production` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `dusk_production` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `sum_production` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `morn_injection` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `noon_injection` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `dusk_injection` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of farm_milk_production_record
-- ----------------------------
BEGIN;
INSERT INTO `farm_milk_production_record` VALUES (1, 7, '2019-06-15 14:17:21', '1.2', '1.3', '1.4', '3.9', '1.5', '1.6', '2');
INSERT INTO `farm_milk_production_record` VALUES (3, 11, '2019-06-15 23:02:10', '2.22', '3.33', '4.44', '9.990000000000002', '5.55', '6.66', '7.77');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
