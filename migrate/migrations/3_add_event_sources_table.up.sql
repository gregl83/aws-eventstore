CREATE TABLE event_sources(
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  trace_id BINARY(16) NOT NULL,
  actor_id BINARY(16) NOT NULL,
  type TINYTEXT NOT NULL,
  version INT UNSIGNED NOT NULL,
  created DATETIME NOT NULL DEFAULT NOW(),
  updated DATETIME NOT NULL DEFAULT NOW() ON UPDATE NOW(),
  PRIMARY KEY (id),
  CONSTRAINT FOREIGN KEY (id)
  REFERENCES events(soucre_id)
)ENGINE=InnoDB;

CREATE TRIGGER event_sources_del BEFORE DELETE ON event_sources FOR EACH ROW
BEGIN
  SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = 'DELETE not allowed on event_sources table';
END;