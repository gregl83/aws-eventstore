-- todo - prevent updates and deletes using mysql trigger or similar

CREATE TABLE event_sources(
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  type TINYTEXT NOT NULL,
  version INT UNSIGNED NOT NULL,
  created DATETIME NOT NULL DEFAULT NOW(),
  updated DATETIME NOT NULL DEFAULT NOW() ON UPDATE NOW(),
  PRIMARY KEY (id),
  CONSTRAINT FOREIGN KEY (id)
  REFERENCES events(soucre_id)
  ON UPDATE CASCADE
  ON DELETE RESTRICT
)ENGINE=InnoDB;