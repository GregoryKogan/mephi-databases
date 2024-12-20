--------------------------------------------------------
-- Library
--------------------------------------------------------

CREATE SCHEMA IF NOT EXISTS "library";

--------------------------------------------------------
--  DDL for Table authors
--------------------------------------------------------

  CREATE TABLE "library"."authors" 
   (	
    "a_id" SERIAL primary KEY, 
	"a_name" VARCHAR(150)
   );
--------------------------------------------------------
--  DDL for Table books
--------------------------------------------------------

  CREATE TABLE "library"."books" 
   (	
    "b_id" SERIAL primary KEY, 
	"b_name" VARCHAR(150), 
	"b_year" SMALLINT, 
	"b_quantity" SMALLINT
   );
--------------------------------------------------------
--  DDL for Table genres
--------------------------------------------------------

  CREATE TABLE "library"."genres" 
   (	
    "g_id" SERIAL primary KEY, 
	"g_name" VARCHAR(150)
   );
--------------------------------------------------------
--  DDL for Table m2m_books_authors
--------------------------------------------------------

  CREATE TABLE "library"."m2m_books_authors" 
   (	
    "b_id" INTEGER, 
	"a_id" INTEGER
   );
--------------------------------------------------------
--  DDL for Table m2m_books_genres
--------------------------------------------------------

  CREATE TABLE "library"."m2m_books_genres" 
   (	
    "b_id" INTEGER, 
	"g_id" INTEGER
   );
--------------------------------------------------------
--  DDL for Table subscribers
--------------------------------------------------------

  CREATE TABLE "library"."subscribers" 
   (	
    "s_id" SERIAL primary KEY, 
	"s_name" VARCHAR(150)
   );
--------------------------------------------------------
--  DDL for Table subscriptions
--------------------------------------------------------

  CREATE TABLE "library"."subscriptions" 
   (	
    "sb_id" SERIAL primary KEY, 
	"sb_subscriber" INTEGER, 
	"sb_book" INTEGER, 
	"sb_start" DATE, 
	"sb_finish" DATE, 
	"sb_is_active" CHAR(1)
   );
  
--------------------------------------------------------
--  Constraints for Table genres
--------------------------------------------------------

  ALTER TABLE "library"."genres" ALTER COLUMN "g_name" SET NOT NULL ;

--------------------------------------------------------
--  Constraints for Table books
--------------------------------------------------------


  ALTER TABLE "library"."books" ALTER COLUMN "b_name" SET NOT NULL ;
  ALTER TABLE "library"."books" ALTER COLUMN "b_year" SET NOT NULL ;
  ALTER TABLE "library"."books" ALTER COLUMN "b_quantity" SET NOT NULL ;

--------------------------------------------------------
--  Constraints for Table m2m_books_genres
--------------------------------------------------------

  ALTER TABLE "library"."m2m_books_genres" ADD CONSTRAINT "PK_m2m_books_genres" PRIMARY KEY ("b_id", "g_id");

--------------------------------------------------------
--  Constraints for Table m2m_books_authors
--------------------------------------------------------

  ALTER TABLE "library"."m2m_books_authors" ADD CONSTRAINT "PK_m2m_books_authors" PRIMARY KEY ("b_id", "a_id");

--------------------------------------------------------
--  Constraints for Table subscribers
--------------------------------------------------------

  ALTER TABLE "library"."subscribers" ALTER COLUMN "s_name" SET NOT NULL ;

--------------------------------------------------------
--  Constraints for Table subscriptions
--------------------------------------------------------

  ALTER TABLE "library"."subscriptions" ALTER COLUMN "sb_subscriber" SET NOT NULL ;
  ALTER TABLE "library"."subscriptions" ALTER COLUMN "sb_book" SET NOT NULL ;
  ALTER TABLE "library"."subscriptions" ALTER COLUMN "sb_finish" SET NOT NULL ;
  ALTER TABLE "library"."subscriptions" ALTER COLUMN "sb_is_active" SET NOT NULL ;
  ALTER TABLE "library"."subscriptions" ADD CONSTRAINT "check_enum" CHECK ("sb_is_active" IN ('Y', 'N'));
 
--------------------------------------------------------
--  Constraints for Table authors
--------------------------------------------------------

  ALTER TABLE "library"."authors" ALTER COLUMN "a_name" SET NOT NULL ;

--------------------------------------------------------
--  Ref Constraints for Table m2m_books_authors
--------------------------------------------------------

  ALTER TABLE "library"."m2m_books_authors" ADD CONSTRAINT "FK_m2m_books_authors_authors" FOREIGN KEY ("a_id")
	  REFERENCES "library"."authors" ("a_id") ON DELETE CASCADE;
  ALTER TABLE "library"."m2m_books_authors" ADD CONSTRAINT "FK_m2m_books_authors_books" FOREIGN KEY ("b_id")
	  REFERENCES "library"."books" ("b_id") ON DELETE CASCADE;
--------------------------------------------------------
--  Ref Constraints for Table m2m_books_genres
--------------------------------------------------------

  ALTER TABLE "library"."m2m_books_genres" ADD CONSTRAINT "FK_m2m_books_genres_books" FOREIGN KEY ("b_id")
	  REFERENCES "library"."books" ("b_id");
  ALTER TABLE "library"."m2m_books_genres" ADD CONSTRAINT "FK_m2m_books_genres_genres" FOREIGN KEY ("g_id")
	  REFERENCES "library"."genres" ("g_id");
--------------------------------------------------------
--  Ref Constraints for Table subscriptions
--------------------------------------------------------

  ALTER TABLE "library"."subscriptions" ADD CONSTRAINT "FK_subscriptions_books" FOREIGN KEY ("sb_book")
	  REFERENCES "library"."books" ("b_id") ON DELETE CASCADE;
  ALTER TABLE "library"."subscriptions" ADD CONSTRAINT "FK_subscriptions_subscribers" FOREIGN KEY ("sb_subscriber")
	  REFERENCES "library"."subscribers" ("s_id") ON DELETE CASCADE;
