CREATE TABLE "link" (
	"tinyurl" VARCHAR(10),
	"fullurl" VARCHAR(4000),
	PRIMARY KEY ("tinyurl")
);

CREATE UNIQUE INDEX "idxfullurl"
on "link" ("fullurl")
