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

 Date: 16/06/2019 15:02:40
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for farm_calve_record
-- ----------------------------
DROP TABLE IF EXISTS `farm_calve_record`;
CREATE TABLE `farm_calve_record` (
  `id` int(50) NOT NULL AUTO_INCREMENT,
  `cow_id` int(50) NOT NULL,
  `birth_time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `flowing_time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `fetus_time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `fetus_organ` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `fetus_birth_time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `placenta_time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `is_complete` tinyint(3) NOT NULL,
  `is_abortion` tinyint(3) NOT NULL,
  `yak_id` int(50) NOT NULL,
  `yak_index` double(50,0) NOT NULL,
  `milk_production` double(50,0) NOT NULL,
  `cream` double(50,0) NOT NULL,
  `protein` double(50,0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of farm_calve_record
-- ----------------------------
BEGIN;
INSERT INTO `farm_calve_record` VALUES (3, 7, '2019-04-21 14:46:43', '2019-04-23 14:46:43', '2019-04-23 14:46:43', '腿', '2019-04-23 14:46:43', '2019-04-23 14:46:43', 1, 2, 11, 1, 2, 3, 4);
INSERT INTO `farm_calve_record` VALUES (4, 101, '2019-04-21 14:46:43', '2019-04-23 14:46:43', '2019-04-23 14:46:43', 'head', '2019-04-23 14:46:43', '2019-04-23 14:46:43', 1, 2, 12, 1, 2, 3, 4);
INSERT INTO `farm_calve_record` VALUES (6, 11, '2019-01-01 14:46:43', '2019-01-01 14:46:43', '2019-01-01 14:46:43', 'head', '2019-01-01 14:46:43', '2019-01-01 14:46:43', 1, 1, 13, 1, 2, 3, 4);
INSERT INTO `farm_calve_record` VALUES (7, 11, '2019-01-01 14:46:43', '2019-01-01 14:46:43', '2019-01-01 14:46:43', 'head', '2019-01-01 14:46:43', '2019-01-01 14:46:43', 1, 1, 13, 1, 2, 3, 4);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
