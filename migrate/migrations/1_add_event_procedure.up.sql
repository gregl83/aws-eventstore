CREATE PROCEDURE add_event(
  sourceId bigint(11),
  traceId binary(16),
  actorId binary(16),
  sourceType varchar(70),
  sourceVersion tinyint(2),
  eventName text(70),
  eventData text(2000)
)
  BEGIN
    DECLARE currentVersion tinyint(2);
	  DECLARE eventId bigint(11);
    DECLARE eventVersion tinyint(2);
    DECLARE EXIT HANDLER FOR SQLEXCEPTION
      BEGIN
        ROLLBACK;
        RESIGNAL;
      END;
    -- event source version must be tinyint using 0 for first event
    IF sourceVersion IS NULL THEN
      SIGNAL SQLSTATE '45000'
      SET MESSAGE_TEXT = 'Event Source version null while trying to add Event to the store';
    END IF;

    -- get event source version number using source id
    IF sourceId IS NOT NULL THEN
      SELECT `version`
      INTO currentVersion
      FROM `event_sources`
      WHERE `id` = sourceId;
      IF currentVersion IS NULL THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Event Source not found while trying to add Event to the store';
      END IF;
    END IF;

    START TRANSACTION;

    -- if current version is null this is a new event source (version is 0)
    IF currentVersion IS NULL THEN
    	SET currentVersion = 0;
    	INSERT INTO `event_sources` (`trace_id`, `actor_id`, `type`, `version`)
    	VALUES (traceId, actorId, sourceType, currentVersion);
      SET sourceId = LAST_INSERT_ID();
    END IF;

    -- sanity check event source version against current version
    IF sourceVersion != currentVersion THEN
      SIGNAL SQLSTATE '45000'
      SET MESSAGE_TEXT = 'Optimistic concurrency error while trying to add Event to store';
    END IF;

    -- insert new event in events table (consider handling event stream)
    INSERT INTO `events` (`trace_id`, `source_id`, `actor_id`, `name`, `version`, `data`)
    VALUES (traceId, sourceId, actorId, eventName, currentVersion, eventData);
	SET eventId = LAST_INSERT_ID();
	SET eventVersion = currentVersion;
    -- increment event source version number for next concurrency test
	SET sourceVersion = currentVersion + 1;
    UPDATE `event_sources`
    SET `version` = sourceVersion
    WHERE `id` = sourceId;

    COMMIT;

    SELECT traceId, sourceId, actorId, sourceVersion, eventId, eventVersion;
  END;