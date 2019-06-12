/*
Navicat MySQL Data Transfer

Source Server         : 连接
Source Server Version : 80013
Source Host           : localhost:3306
Source Database       : farm

Target Server Type    : MYSQL
Target Server Version : 80013
File Encoding         : 65001

Date: 2019-06-12 14:09:33
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for farm_calve_record
-- ----------------------------
DROP TABLE IF EXISTS `farm_calve_record`;
CREATE TABLE `farm_calve_record` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cow_id` int(11) DEFAULT NULL,
  `fetus_organ` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `is_complete` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `is_abortion` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `yak_id` int(11) DEFAULT NULL,
  `yak_index` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `milk_production` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `cream` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `protein` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `birth_time` varchar(0) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `flowing_time` varchar(0) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `fetus_time` varchar(0) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `fetus_birth_time` varchar(0) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `placenta_time` varchar(0) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of farm_calve_record
-- ----------------------------

-- ----------------------------
-- Table structure for farm_user
-- ----------------------------
DROP TABLE IF EXISTS `farm_user`;
CREATE TABLE `farm_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `role` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `telephone_number` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `e_mail` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of farm_user
-- ----------------------------
INSERT INTO `farm_user` VALUES ('1', 'bobo', '123456', '1', '123456', '111@x.com');
INSERT INTO `farm_user` VALUES ('2', 'bobo1', '123456', '1', '15623253450', 'zjb770168382@qq.com');
