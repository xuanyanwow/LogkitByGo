CREATE TABLE `siam_api_log`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `project_id` int(11) NOT NULL COMMENT '所属项目id',
  `api_full` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'api路径 = api类目.\"/\".api方法',
  `api_category` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'api类目',
  `api_method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'api方法',
  `api_param` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL COMMENT 'api参数',
  `api_response` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL COMMENT 'api响应',
  `is_success` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '成功|失败 1|0',
  `consume_time` int(10) NOT NULL COMMENT '消耗时间 单位ms',
  `user_from` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '用户来源，可以填入ip、城市名、调用账号等类型',
  `user_identify` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '用户标识，比如可以用订单号，结合来源，就可以定位请求',
  `create_date` int(10) NOT NULL COMMENT '记录日期  YYYYddmm',
  `create_time` datetime(0) NOT NULL COMMENT '记录时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `用户搜索`(`user_from`, `user_identify`) USING BTREE,
  INDEX `日期`(`create_date`, `create_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;