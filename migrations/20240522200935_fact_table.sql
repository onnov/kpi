-- +goose Up
-- +goose StatementBegin
CREATE TABLE `kpi`.`fact`
(
    `id`                      INT          NOT NULL AUTO_INCREMENT,
    `period_start`            DATE         NOT NULL,
    `period_end`              DATE         NOT NULL,
    `period_key`              VARCHAR(255) NOT NULL,
    `indicator_to_mo_id`      INT          NOT NULL,
    `indicator_to_mo_fact_id` INT          NOT NULL,
    `value`                   INT          NOT NULL,
    `fact_time`               DATE         NOT NULL,
    `is_plan`                 TINYINT(1)   NOT NULL,
    `auth_user_id`            INT          NOT NULL,
    `comment`                 TEXT         NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `kpi`.`fact`;
-- +goose StatementEnd
