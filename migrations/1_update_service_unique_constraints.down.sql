BEGIN;

ALTER TABLE "service"
DROP CONSTRAINT "service_name_key";

ALTER TABLE "service"
ADD CONSTRAINT "service_name_team_key" UNIQUE ("name", "team");

ALTER TABLE "incident"
DROP CONSTRAINT "incident_fk0";

ALTER TABLE "incident" ADD CONSTRAINT "incident_fk0" FOREIGN KEY ("service_id") REFERENCES "service"("id");

COMMIT;
