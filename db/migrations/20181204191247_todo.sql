
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE todos (
  id int(11) NOT NULL,
  name varchar(255),
  completed tinyint(1),
  due datetime,
  PRIMARY KEY(id)
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE todos;