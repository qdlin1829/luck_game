/*
Navicat MySQL Data Transfer

Source Server         : dev
Source Server Version : 50726
Source Host           : 127.0.0.1:3306
Source Database       : gin-vue

Target Server Type    : MYSQL
Target Server Version : 50726
File Encoding         : 65001

Date: 2020-04-10 20:02:38
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for g_color_list
-- ----------------------------
DROP TABLE IF EXISTS `g_color_list`;
CREATE TABLE `g_color_list` (
  `cid` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `color_name` varchar(30) NOT NULL COMMENT '彩种类型',
  `state` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态: 1 正常 0禁用',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`cid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT=' 彩种类型';

-- ----------------------------
-- Records of g_color_list
-- ----------------------------

-- ----------------------------
-- Table structure for g_lottery_history
-- ----------------------------
DROP TABLE IF EXISTS `g_lottery_history`;
CREATE TABLE `g_lottery_history` (
  `id` int(11) NOT NULL,
  `color_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '彩种类型 1幸运飞艇',
  `sq` char(20) NOT NULL COMMENT '期号',
  `new_sq` char(20) NOT NULL COMMENT '下一期号',
  `open_time` int(1) unsigned NOT NULL DEFAULT '0' COMMENT ' 开奖时间',
  `new_open_time` int(11) unsigned NOT NULL COMMENT '下一期开奖时间',
  `sumfs` tinyint(1) unsigned NOT NULL COMMENT '总合计',
  `sumdx` char(3) NOT NULL COMMENT '总合，单双',
  `sumds` char(3) NOT NULL COMMENT '总和大小',
  `number` varchar(30) NOT NULL COMMENT '开奖号码:,切格',
  `creata_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_sq` (`sq`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='开奖历史表';

-- ----------------------------
-- Records of g_lottery_history
-- ----------------------------

-- ----------------------------
-- Table structure for g_order_goods
-- ----------------------------
DROP TABLE IF EXISTS `g_order_goods`;
CREATE TABLE `g_order_goods` (
  `og_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `order_id` int(11) unsigned NOT NULL COMMENT '订单关联id',
  `color_type` tinyint(100) unsigned NOT NULL DEFAULT '1' COMMENT '彩种类型 1幸运飞艇',
  `sq` varchar(20) NOT NULL COMMENT '期号',
  `pos` tinyint(1) unsigned NOT NULL COMMENT '位置',
  `number` varchar(5) NOT NULL COMMENT '购买数： 大小单双 1-9',
  `goods_num` smallint(5) unsigned NOT NULL DEFAULT '1' COMMENT '购买数量',
  `amount` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '下单金额',
  PRIMARY KEY (`og_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='订单彩种表';

-- ----------------------------
-- Records of g_order_goods
-- ----------------------------

-- ----------------------------
-- Table structure for g_order_info
-- ----------------------------
DROP TABLE IF EXISTS `g_order_info`;
CREATE TABLE `g_order_info` (
  `order_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `order_sn` varchar(30) NOT NULL COMMENT '订单号',
  `order_type` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT ' 类型:1幸运飞艇',
  `order_status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态: 1待开奖 101未中奖 201已中奖 301无效',
  `user_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `sq` varchar(20) NOT NULL COMMENT '期号',
  `pos` tinyint(1) unsigned NOT NULL COMMENT '位置0-10 0位是总和',
  `number` varchar(10) NOT NULL COMMENT '购买数:大小单双1-10',
  `pay_amount` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '赔付金额',
  `odds` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '赔率',
  `order_amount` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '订单金额',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`order_id`),
  UNIQUE KEY `uq_ordersn` (`order_sn`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COMMENT='订单主表';

-- ----------------------------
-- Records of g_order_info
-- ----------------------------
INSERT INTO `g_order_info` VALUES ('1', '232', '1', '1', '6', '20169', '0', '大', '0.00', '2.12', '1.00', '0', '0');
INSERT INTO `g_order_info` VALUES ('2', '456', '1', '1', '6', '20168', '0', '单', '0.00', '1.94', '1.00', '0', '0');
INSERT INTO `g_order_info` VALUES ('3', '333', '1', '1', '1', '20162', '1', '3', '0.00', '1.94', '10.00', '0', '0');
INSERT INTO `g_order_info` VALUES ('5', '2321', '1', '1', '6', '201691', '0', '大', '0.00', '2.12', '1.00', '0', '0');
INSERT INTO `g_order_info` VALUES ('6', '4561', '1', '1', '6', '201681', '0', '单', '0.00', '1.94', '1.00', '0', '0');
INSERT INTO `g_order_info` VALUES ('7', '3331', '1', '1', '6', '201621', '1', '3', '0.00', '1.94', '10.00', '0', '0');

-- ----------------------------
-- Table structure for g_user
-- ----------------------------
DROP TABLE IF EXISTS `g_user`;
CREATE TABLE `g_user` (
  `user_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `username` varchar(30) NOT NULL COMMENT '帐号',
  `password` varchar(100) NOT NULL COMMENT '密码',
  `state` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT ' 状态: 1正常 0禁用',
  `create_time` int(11) unsigned NOT NULL COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `uq_username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- ----------------------------
-- Records of g_user
-- ----------------------------
INSERT INTO `g_user` VALUES ('1', 'admin', '0192023a7bbd73250516f069df18b500', '0', '1586488743', '1586488743');
INSERT INTO `g_user` VALUES ('2', 'admin3', '0192023a7bbd73250516f069df18b500', '0', '1586488783', '1586488783');
INSERT INTO `g_user` VALUES ('3', 'admina', '0192023a7bbd73250516f069df18b500', '0', '1586488849', '1586488849');
INSERT INTO `g_user` VALUES ('4', 'admina3', '0192023a7bbd73250516f069df18b500', '0', '1586488905', '1586488905');
INSERT INTO `g_user` VALUES ('5', 'admina32', '0192023a7bbd73250516f069df18b500', '0', '1586489030', '1586489030');
INSERT INTO `g_user` VALUES ('6', 'admin323', '0192023a7bbd73250516f069df18b500', '0', '1586489166', '1586489166');
INSERT INTO `g_user` VALUES ('7', 'atew', '0192023a7bbd73250516f069df18b500', '0', '1586489185', '1586489185');
INSERT INTO `g_user` VALUES ('8', 'test123', '0192023a7bbd73250516f069df18b500', '0', '1586489273', '1586489273');
INSERT INTO `g_user` VALUES ('9', 'test1237', '0192023a7bbd73250516f069df18b500', '0', '1586489305', '1586489305');
INSERT INTO `g_user` VALUES ('10', 'test12374', '0192023a7bbd73250516f069df18b500', '0', '1586489453', '1586489453');
INSERT INTO `g_user` VALUES ('11', 'test123743', '0192023a7bbd73250516f069df18b500', '0', '1586489506', '1586489506');
INSERT INTO `g_user` VALUES ('12', 'test123743999', '0192023a7bbd73250516f069df18b500', '0', '1586489614', '1586489614');
INSERT INTO `g_user` VALUES ('13', 'test1237439993', '0192023a7bbd73250516f069df18b500', '0', '1586489667', '1586489667');

-- ----------------------------
-- Table structure for g_user_balance
-- ----------------------------
DROP TABLE IF EXISTS `g_user_balance`;
CREATE TABLE `g_user_balance` (
  `balance_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `amount` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '用户余额',
  `lock_amount` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '锁定金额',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT ' 更新时间',
  PRIMARY KEY (`balance_id`),
  UNIQUE KEY `uq_userid` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户余额表';

-- ----------------------------
-- Records of g_user_balance
-- ----------------------------

-- ----------------------------
-- Table structure for g_user_flow
-- ----------------------------
DROP TABLE IF EXISTS `g_user_flow`;
CREATE TABLE `g_user_flow` (
  `fid` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `flow_type` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '1支出，2收入',
  `user_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `amount` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '金额',
  `before_amount` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '修改前金额',
  `after_amount` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '修改后金额',
  `content` varchar(200) NOT NULL COMMENT '备注',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`fid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户流水表';

-- ----------------------------
-- Records of g_user_flow
-- ----------------------------
