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

 Date: 12/06/2019 13:58:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for farm_user
-- ----------------------------
DROP TABLE IF EXISTS `farm_user`;
CREATE TABLE `farm_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `role` int(4) NOT NULL,
  `telephone_number` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `e_mail` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of farm_user
-- ----------------------------
BEGIN;
INSERT INTO `farm_user` VALUES (1, 'bobo', '123456', 1, '123456', '111@x.com');
INSERT INTO `farm_user` VALUES (4, 'whiskey11', '12345678', 1, '13593693130', 'shgal@qq.com');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
