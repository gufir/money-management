ALTER TABLE "sessions" DROP CONSTRAINT IF EXISTS "sessions_user_id_fkey";
DROP TABLE IF EXISTS "transaction";
DROP TABLE IF EXISTS "categories";
DROP TABLE IF EXISTS "reports";
DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "sessions";