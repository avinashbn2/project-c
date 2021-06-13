CREATE INDEX titleIndex ON resource_item USING GIN (to_tsvector('english', title));
CREATE INDEX excerptIndex ON resource_item USING GIN (to_tsvector('english', excerpt));
