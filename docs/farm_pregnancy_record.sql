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

 Date: 16/06/2019 15:03:11
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for farm_pregnancy_record
-- ----------------------------
DROP TABLE IF EXISTS `farm_pregnancy_record`;
CREATE TABLE `farm_pregnancy_record` (
  `cow_id` int(50) NOT NULL,
  `check_time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `final_time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `estimate_time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `ovary_diameter` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `follicle_diameter` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `uterine_inflammation` tinyint(2) NOT NULL,
  `vaginal_inflammation` tinyint(2) NOT NULL,
  `is_pregnancy` tinyint(4) NOT NULL,
  `is_empty` tinyint(4) NOT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of farm_pregnancy_record
-- ----------------------------
BEGIN;
INSERT INTO `farm_pregnancy_record` VALUES (7, '2006-01-02 15:04:55', '2019-04-26 14:46:43', '2020-02-25 14:46:43', '44.2', '33.1', 1, 1, 1, 1, 1);
INSERT INTO `farm_pregnancy_record` VALUES (7, '2006-01-02 15:04:55', '2019-04-26 14:46:43', '2020-02-25 14:46:43', '44.2', '33.1', 1, 1, 1, 1, 2);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
