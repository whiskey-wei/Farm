/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80011
 Source Host           : localhost:3306
 Source Schema         : farm

 Target Server Type    : MySQL
 Target Server Version : 80011
 File Encoding         : 65001

 Date: 12/06/2019 21:43:16
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for farm_calve_record
-- ----------------------------
DROP TABLE IF EXISTS `farm_calve_record`;
CREATE TABLE `farm_calve_record`  (
  `id` int(50) NOT NULL AUTO_INCREMENT,
  `cow_id` int(50) NOT NULL,
  `birth_time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `flowing_time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `fetus_time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `fetus_organ` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `fetus_birth_time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `placenta_time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `is_complete` tinyint(3) NULL DEFAULT NULL,
  `is_abortion` tinyint(3) NULL DEFAULT NULL,
  `yak_id` int(50) NULL DEFAULT NULL,
  `yak_index` double(50, 0) NULL DEFAULT NULL,
  `milk_production` double(50, 0) NULL DEFAULT NULL,
  `cream` double(50, 0) NULL DEFAULT NULL,
  `protein` double(50, 0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of farm_calve_record
-- ----------------------------
INSERT INTO `farm_calve_record` VALUES (3, 7, '2019-04-21 14:46:43', '2019-04-23 14:46:43', '2019-04-23 14:46:43', '腿', '2019-04-23 14:46:43', '2019-04-23 14:46:43', 1, 2, 11, 1, 2, 3, 4);
INSERT INTO `farm_calve_record` VALUES (4, 101, '2019-04-21 14:46:43', '2019-04-23 14:46:43', '2019-04-23 14:46:43', 'head', '2019-04-23 14:46:43', '2019-04-23 14:46:43', 1, 2, 11, 1, 2, 3, 4);

SET FOREIGN_KEY_CHECKS = 1;
