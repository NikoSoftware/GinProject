/*
 Navicat Premium Data Transfer

 Source Server         : xiaomotou
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : www.xiaomotou.cn:3306
 Source Schema         : student

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 24/10/2022 00:47:01
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for course
-- ----------------------------
DROP TABLE IF EXISTS `course`;
CREATE TABLE `course`  (
  `c_no` varchar(10) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `c_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
  `t_no` varchar(20) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`c_no`, `t_no`) USING BTREE,
  INDEX `Teno`(`t_no`) USING BTREE,
  CONSTRAINT `Teno` FOREIGN KEY (`t_no`) REFERENCES `teacher` (`t_no`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of course
-- ----------------------------
INSERT INTO `course` VALUES ('c001', 'J2SE', 't002');
INSERT INTO `course` VALUES ('c002', 'Java Web', 't002');
INSERT INTO `course` VALUES ('c003', 'SSH', 't001');
INSERT INTO `course` VALUES ('c004', 'Oracle', 't001');
INSERT INTO `course` VALUES ('c005', 'SQL SERVER 2005', 't003');
INSERT INTO `course` VALUES ('c006', 'C#', 't003');
INSERT INTO `course` VALUES ('c007', 'JavaScript', 't002');
INSERT INTO `course` VALUES ('c008', 'DIV+CSS', 't001');
INSERT INTO `course` VALUES ('c009', 'PHP', 't003');
INSERT INTO `course` VALUES ('c010', 'EJB3.0', 't002');

-- ----------------------------
-- Table structure for sc
-- ----------------------------
DROP TABLE IF EXISTS `sc`;
CREATE TABLE `sc`  (
  `s_no` varchar(10) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `c_no` varchar(10) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `score` float NULL DEFAULT NULL,
  PRIMARY KEY (`s_no`, `c_no`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sc
-- ----------------------------
INSERT INTO `sc` VALUES ('s001', 'c001', 78.9);
INSERT INTO `sc` VALUES ('s001', 'c002', 82.9);
INSERT INTO `sc` VALUES ('s001', 'c003', 59);
INSERT INTO `sc` VALUES ('s002', 'c001', 80.9);
INSERT INTO `sc` VALUES ('s002', 'c002', 72.9);
INSERT INTO `sc` VALUES ('s003', 'c001', 81.9);
INSERT INTO `sc` VALUES ('s003', 'c002', 81.9);
INSERT INTO `sc` VALUES ('s004', 'c001', 60.9);

-- ----------------------------
-- Table structure for student
-- ----------------------------
DROP TABLE IF EXISTS `student`;
CREATE TABLE `student`  (
  `s_no` varchar(10) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `s_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
  `s_age` int NULL DEFAULT NULL,
  `s_sex` varchar(5) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
  PRIMARY KEY (`s_no`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of student
-- ----------------------------
INSERT INTO `student` VALUES ('s001', '张三', 23, '男');
INSERT INTO `student` VALUES ('s002', '李四', 23, '男');
INSERT INTO `student` VALUES ('s003', '吴鹏', 25, '男');
INSERT INTO `student` VALUES ('s004', '琴沁', 20, '女');
INSERT INTO `student` VALUES ('s005', '王丽', 20, '女');
INSERT INTO `student` VALUES ('s006', '李波', 21, '男');
INSERT INTO `student` VALUES ('s007', '刘玉', 21, '男');
INSERT INTO `student` VALUES ('s008', '萧蓉', 21, '女');
INSERT INTO `student` VALUES ('s009', '陈萧晓', 23, '女');
INSERT INTO `student` VALUES ('s010', '陈美', 22, '女');

-- ----------------------------
-- Table structure for teacher
-- ----------------------------
DROP TABLE IF EXISTS `teacher`;
CREATE TABLE `teacher`  (
  `t_no` varchar(10) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `t_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
  PRIMARY KEY (`t_no`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of teacher
-- ----------------------------
INSERT INTO `teacher` VALUES ('t001', '刘阳');
INSERT INTO `teacher` VALUES ('t002', '谌燕');
INSERT INTO `teacher` VALUES ('t003', '胡明星');

SET FOREIGN_KEY_CHECKS = 1;
