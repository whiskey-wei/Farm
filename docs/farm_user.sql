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

 Date: 16/06/2019 15:03:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for farm_user
-- ----------------------------
DROP TABLE IF EXISTS `farm_user`;
CREATE TABLE `farm_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '账号',
  `password` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `role` int(2) NOT NULL,
  `telephone_number` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `e_mail` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of farm_user
-- ----------------------------
BEGIN;
INSERT INTO `farm_user` VALUES (1, 'test', '12345678', 1, '13593693130', 'shgal@qq.com');
INSERT INTO `farm_user` VALUES (2, 'whiskey', '12345678', 1, '13593693130', 'shgal@qq.com');
INSERT INTO `farm_user` VALUES (3, 'admin', '12345678', 2, '13593693130', 'shgal@qq.com');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
