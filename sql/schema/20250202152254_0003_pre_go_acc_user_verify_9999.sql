-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_go_acc_user_base_9999`;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE `pre_go_acc_user_base_9999`
(
    `user_id`          int          NOT NULL AUTO_INCREMENT,
    `user_account`     varchar(255) NOT NULL,
    `user_password`    varchar(255) NOT NULL,
    `user_salt`        varchar(255) NOT NULL,
    `user_login_time`  timestamp NULL DEFAULT NULL,
    `user_logout_time` timestamp NULL DEFAULT NULL,
    `user_login_ip`    varchar(45) DEFAULT NULL,
    `user_created_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `user_updated_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`user_id`),
    UNIQUE KEY `unique_user_account` (`user_account`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 4
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='pre_go_acc';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_go_acc_user_base_9999`;
-- +goose StatementEnd