/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50628
Source Host           : 20.30.1.120:32777
Source Database       : mangmang

Target Server Type    : MYSQL
Target Server Version : 50628
File Encoding         : 65001

Date: 2019-04-18 17:54:29
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for business_card
-- ----------------------------
DROP TABLE IF EXISTS `business_card`;
CREATE TABLE `business_card` (
  `business_id` varchar(36) NOT NULL,
  `user_id` varchar(36) NOT NULL COMMENT '用户ID',
  `name` varchar(50) NOT NULL COMMENT '姓名',
  `company` varchar(255) NOT NULL COMMENT '公司',
  `position` varchar(255) NOT NULL COMMENT '职位',
  `phone` varchar(20) NOT NULL COMMENT '电话',
  `qq` varchar(15) DEFAULT NULL COMMENT 'QQ号',
  `wx` varchar(50) DEFAULT NULL COMMENT '微信号',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `data_status` smallint(1) NOT NULL DEFAULT '1' COMMENT '0删除，1有效',
  PRIMARY KEY (`business_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `business_card_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='电子名片表';

-- ----------------------------
-- Records of business_card
-- ----------------------------

-- ----------------------------
-- Table structure for mechanism
-- ----------------------------
DROP TABLE IF EXISTS `mechanism`;
CREATE TABLE `mechanism` (
  `id` varchar(36) NOT NULL,
  `name` varchar(255) NOT NULL COMMENT '公司名字',
  `address` varchar(255) DEFAULT NULL COMMENT '地址',
  `phone` varchar(20) DEFAULT NULL COMMENT '电话',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `data_status` smallint(1) NOT NULL DEFAULT '1' COMMENT '状态0有效，1无效',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='公司';

-- ----------------------------
-- Records of mechanism
-- ----------------------------

-- ----------------------------
-- Table structure for project
-- ----------------------------
DROP TABLE IF EXISTS `project`;
CREATE TABLE `project` (
  `project_id` varchar(36) NOT NULL,
  `project_name` varchar(50) NOT NULL COMMENT '项目名称',
  `project_content` text COMMENT '项目描述',
  `create_time` datetime NOT NULL,
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `data_status` smallint(1) NOT NULL DEFAULT '1' COMMENT '0删除或无效，1有效',
  PRIMARY KEY (`project_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of project
-- ----------------------------

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `role_id` varchar(36) NOT NULL,
  `role_level` int(11) NOT NULL DEFAULT '1' COMMENT '权限大小',
  `role_name` varchar(50) NOT NULL COMMENT '角色名称',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `data_status` smallint(1) NOT NULL DEFAULT '1' COMMENT '0无效，1有效',
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of role
-- ----------------------------

-- ----------------------------
-- Table structure for task
-- ----------------------------
DROP TABLE IF EXISTS `task`;
CREATE TABLE `task` (
  `task_id` varchar(36) NOT NULL,
  `father_task_id` varchar(36) DEFAULT NULL COMMENT '父任务',
  `project_id` varchar(36) NOT NULL COMMENT '项目ID',
  `task_name` varchar(255) NOT NULL,
  `task_type` varchar(10) NOT NULL DEFAULT '需求' COMMENT '任务类型',
  `task_content` text COMMENT '任务内容',
  `task_schedule` int(11) NOT NULL DEFAULT '0' COMMENT '进度',
  `task_status` varchar(10) NOT NULL DEFAULT '新建' COMMENT '任务状态',
  `starting_time` datetime DEFAULT NULL COMMENT '开始时间',
  `planned_completion_time` datetime DEFAULT NULL COMMENT '计划完成时间',
  `create_time` datetime NOT NULL,
  `update_time` datetime NOT NULL,
  `data_status` smallint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`task_id`),
  KEY `project_id` (`project_id`),
  CONSTRAINT `task_ibfk_1` FOREIGN KEY (`project_id`) REFERENCES `project` (`project_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of task
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` varchar(36) NOT NULL COMMENT '用户ID',
  `name` varchar(20) NOT NULL COMMENT '用户名',
  `avatar_url` varchar(255) DEFAULT NULL COMMENT '头像',
  `email` varchar(255) DEFAULT NULL COMMENT '邮箱',
  `phone` varchar(11) DEFAULT NULL COMMENT '电话',
  `sex` smallint(1) NOT NULL DEFAULT '0' COMMENT '性别：0保密，1男，2女',
  `birthday` date DEFAULT NULL COMMENT '生日',
  `address` varchar(255) DEFAULT NULL COMMENT '地址',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `data_status` smallint(1) NOT NULL DEFAULT '1' COMMENT '数据状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户信息表';

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('2bca066a-2446-4b1d-b24c-0ed6e0dad162', 'qin', '', '', '18328020611', '0', '2019-04-03', '', '2019-04-03 15:44:11', '2019-04-03 15:44:11', '1');

-- ----------------------------
-- Table structure for user_login_method
-- ----------------------------
DROP TABLE IF EXISTS `user_login_method`;
CREATE TABLE `user_login_method` (
  `id` varchar(36) NOT NULL,
  `user_id` varchar(36) NOT NULL COMMENT '用户ID',
  `login_type` varchar(2) NOT NULL COMMENT '登陆方式',
  `identification` varchar(50) NOT NULL COMMENT '登陆账户或其他ID',
  `access_code` varchar(255) NOT NULL COMMENT '登陆密码或者其他授权',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `data_status` smallint(1) NOT NULL DEFAULT '1' COMMENT '账户状态0禁用,1有效',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `user_login_method_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户登录方式表';

-- ----------------------------
-- Records of user_login_method
-- ----------------------------
INSERT INTO `user_login_method` VALUES ('a849823a-9b2c-4aa5-9a44-214aecb1921a', '2bca066a-2446-4b1d-b24c-0ed6e0dad162', 'P', '18328020611', '63faf9a9e04759f4ece532f53c0c8129', '2019-04-03 15:44:11', '2019-04-03 15:44:11', '1');

-- ----------------------------
-- Table structure for user_project_mapping
-- ----------------------------
DROP TABLE IF EXISTS `user_project_mapping`;
CREATE TABLE `user_project_mapping` (
  `id` varchar(36) NOT NULL,
  `user_id` varchar(36) NOT NULL COMMENT '用户id',
  `project_id` varchar(36) NOT NULL COMMENT '项目id',
  `role_id` varchar(36) NOT NULL COMMENT '角色ID',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `data_status` smallint(1) NOT NULL DEFAULT '1' COMMENT '0，无效，1有效',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `project_id` (`project_id`),
  KEY `role_id` (`role_id`),
  CONSTRAINT `user_project_mapping_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `user_project_mapping_ibfk_2` FOREIGN KEY (`project_id`) REFERENCES `project` (`project_id`),
  CONSTRAINT `user_project_mapping_ibfk_3` FOREIGN KEY (`role_id`) REFERENCES `role` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user_project_mapping
-- ----------------------------
