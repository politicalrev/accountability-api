-- +goose Up
INSERT INTO politicians (id, name, title, country) VALUES ('donald-trump', 'Donald Trump', 'President', 'US');

INSERT INTO sources (name, link) VALUES ('100 day plan', 'http://www.npr.org/2016/11/09/501451368/here-is-what-donald-trump-wants-to-do-in-his-first-100-days');

INSERT INTO promises (politician_id, name, details, category) VALUES ('donald-trump', 'Propose a Constitutional Amendment to impose term limits on all members of Congress', '', 'government');
INSERT INTO promise_status (promise_id, name, detail) SELECT currval('promises_id_seq'), 'not-started', 'Donald Trump outlined his plan prior to becoming president.';
INSERT INTO promise_sources SELECT pid, sid FROM currval('promises_id_seq') AS pid, currval('sources_id_seq') AS sid;

INSERT INTO promises (politician_id, name, details, category) VALUES ('donald-trump', 'Hiring freeze on all federal employees to reduce federal workforce through attrition (exempting military, public safety, and public health)', '', 'government');
INSERT INTO promise_status (promise_id, name, detail) SELECT currval('promises_id_seq'), 'not-started', 'Donald Trump outlined his plan prior to becoming president.';
INSERT INTO promise_sources SELECT pid, sid FROM currval('promises_id_seq') AS pid, currval('sources_id_seq') AS sid;

INSERT INTO promises (politician_id, name, details, category) VALUES ('donald-trump', 'A requirement that for every new federal regulation, two existing regulations must be eliminated', '', 'government');
INSERT INTO promise_status (promise_id, name, detail) SELECT currval('promises_id_seq'), 'not-started', 'Donald Trump outlined his plan prior to becoming president.';
INSERT INTO promise_sources SELECT pid, sid FROM currval('promises_id_seq') AS pid, currval('sources_id_seq') AS sid;

INSERT INTO promises (politician_id, name, details, category) VALUES ('donald-trump', 'A 5 year-ban on White House and Congressional officials becoming lobbyists after they leave government service', '', 'government');
INSERT INTO promise_status (promise_id, name, detail) SELECT currval('promises_id_seq'), 'not-started', 'Donald Trump outlined his plan prior to becoming president.';
INSERT INTO promise_sources SELECT pid, sid FROM currval('promises_id_seq') AS pid, currval('sources_id_seq') AS sid;

INSERT INTO promises (politician_id, name, details, category) VALUES ('donald-trump', 'A lifetime ban on White House officials lobbying on behalf of a foreign government', '', 'government');
INSERT INTO promise_status (promise_id, name, detail) SELECT currval('promises_id_seq'), 'not-started', 'Donald Trump outlined his plan prior to becoming president.';
INSERT INTO promise_sources SELECT pid, sid FROM currval('promises_id_seq') AS pid, currval('sources_id_seq') AS sid;

INSERT INTO promises (politician_id, name, details, category) VALUES ('donald-trump', 'A complete ban on foreign lobbyists raising money for American elections.', '', 'government');
INSERT INTO promise_status (promise_id, name, detail) SELECT currval('promises_id_seq'), 'not-started', 'Donald Trump outlined his plan prior to becoming president.';
INSERT INTO promise_sources SELECT pid, sid FROM currval('promises_id_seq') AS pid, currval('sources_id_seq') AS sid;

INSERT INTO promises (politician_id, name, details, category) VALUES ('donald-trump', 'Renegotiate NAFTA or withdraw from the deal under Article 2205', '', 'economy');
INSERT INTO promise_status (promise_id, name, detail) SELECT currval('promises_id_seq'), 'not-started', 'Donald Trump outlined his plan prior to becoming president.';
INSERT INTO promise_sources SELECT pid, sid FROM currval('promises_id_seq') AS pid, currval('sources_id_seq') AS sid;

INSERT INTO promises (politician_id, name, details, category) VALUES ('donald-trump', 'Withdrawal from the Trans-Pacific Partnership', '', 'economy');
INSERT INTO promise_status (promise_id, name, detail) SELECT currval('promises_id_seq'), 'not-started', 'Donald Trump outlined his plan prior to becoming president.';
INSERT INTO promise_sources SELECT pid, sid FROM currval('promises_id_seq') AS pid, currval('sources_id_seq') AS sid;

INSERT INTO promises (politician_id, name, details, category) VALUES ('donald-trump', 'Direct the Secretary of the Treasury to label China a currency manipulator', '', 'economy');
INSERT INTO promise_status (promise_id, name, detail) SELECT currval('promises_id_seq'), 'not-started', 'Donald Trump outlined his plan prior to becoming president.';
INSERT INTO promise_sources SELECT pid, sid FROM currval('promises_id_seq') AS pid, currval('sources_id_seq') AS sid;

INSERT INTO promises (politician_id, name, details, category) VALUES ('donald-trump', 'Direct the Secretary of Commerce and U.S. Trade Representative to identify all foreign trading abuses that unfairly impact American workers and direct them to use every tool under American and international law to end those abuses immediately', '', 'economy');
INSERT INTO promise_status (promise_id, name, detail) SELECT currval('promises_id_seq'), 'not-started', 'Donald Trump outlined his plan prior to becoming president.';
INSERT INTO promise_sources SELECT pid, sid FROM currval('promises_id_seq') AS pid, currval('sources_id_seq') AS sid;

INSERT INTO promises (politician_id, name, details, category) VALUES ('donald-trump', 'Lift the restrictions on the production of $50 trillion dollars'' worth of job-producing American energy reserves, including shale, oil, natural gas and clean coal.', '', 'climate');
INSERT INTO promise_status (promise_id, name, detail) SELECT currval('promises_id_seq'), 'not-started', 'Donald Trump outlined his plan prior to becoming president.';
INSERT INTO promise_sources SELECT pid, sid FROM currval('promises_id_seq') AS pid, currval('sources_id_seq') AS sid;

INSERT INTO promises (politician_id, name, details, category) VALUES ('donald-trump', 'Lift the Obama-Clinton roadblocks and allow vital energy infrastructure projects, like the Keystone Pipeline, to move forward', '', 'climate');
INSERT INTO promise_status (promise_id, name, detail) SELECT currval('promises_id_seq'), 'not-started', 'Donald Trump outlined his plan prior to becoming president.';
INSERT INTO promise_sources SELECT pid, sid FROM currval('promises_id_seq') AS pid, currval('sources_id_seq') AS sid;

INSERT INTO promises (politician_id, name, details, category) VALUES ('donald-trump', 'Cancel billions in payments to U.N. climate change programs and use the money to fix America''s water and environmental infrastructure', '', 'climate');
INSERT INTO promise_status (promise_id, name, detail) SELECT currval('promises_id_seq'), 'not-started', 'Donald Trump outlined his plan prior to becoming president.';
INSERT INTO promise_sources SELECT pid, sid FROM currval('promises_id_seq') AS pid, currval('sources_id_seq') AS sid;

INSERT INTO promises (politician_id, name, details, category) VALUES ('donald-trump', 'Cancel every unconstitutional executive action, memorandum and order issued by President Obama', '', 'government');
INSERT INTO promise_status (promise_id, name, detail) SELECT currval('promises_id_seq'), 'not-started', 'Donald Trump outlined his plan prior to becoming president.';
INSERT INTO promise_sources SELECT pid, sid FROM currval('promises_id_seq') AS pid, currval('sources_id_seq') AS sid;

-- +goose Down
TRUNCATE promise_sources CASCADE;
TRUNCATE promise_status CASCADE;
TRUNCATE promises CASCADE;
TRUNCATE politicians CASCADE;
