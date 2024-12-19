--------------------------------------------------------
-- Library
--------------------------------------------------------

CREATE SCHEMA "library";

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


Insert into "library"."authors" ("a_id","a_name") values (1,'Donald Knuth');
Insert into "library"."authors" ("a_id","a_name") values (2,'Isaac Asimov');
Insert into "library"."authors" ("a_id","a_name") values (3,'Dale Carnegie');
Insert into "library"."authors" ("a_id","a_name") values (4,'Lev Landau');
Insert into "library"."authors" ("a_id","a_name") values (5,'Evgeny Lifshitz');
Insert into "library"."authors" ("a_id","a_name") values (6,'Bjarne Stroustrup');
Insert into "library"."authors" ("a_id","a_name") values (7,'Alexander Pushkin');


Insert into "library"."books" ("b_id","b_name","b_year","b_quantity") values (1,'Eugene Onegin',1985,2);
Insert into "library"."books" ("b_id","b_name","b_year","b_quantity") values (2,'The Fishermen and the Golden Fish',1990,3);
Insert into "library"."books" ("b_id","b_name","b_year","b_quantity") values (3,'Foundation and Empire',2000,5);
Insert into "library"."books" ("b_id","b_name","b_year","b_quantity") values (4,'Programming Psychology',1998,1);
Insert into "library"."books" ("b_id","b_name","b_year","b_quantity") values (5,'The C++ Programming Language',1996,3);
Insert into "library"."books" ("b_id","b_name","b_year","b_quantity") values (6,'Course of Theoretical Physics',1981,12);
Insert into "library"."books" ("b_id","b_name","b_year","b_quantity") values (7,'The Art of Computer Programming',1993,7);

Insert into "library"."genres" ("g_id","g_name") values (5,'Classic');
Insert into "library"."genres" ("g_id","g_name") values (4,'Science');
Insert into "library"."genres" ("g_id","g_name") values (1,'Poetry');
Insert into "library"."genres" ("g_id","g_name") values (2,'Programming');
Insert into "library"."genres" ("g_id","g_name") values (3,'Psychology');
Insert into "library"."genres" ("g_id","g_name") values (6,'Science Fiction');

Insert into "library"."m2m_books_authors" ("b_id","a_id") values (1,7);
Insert into "library"."m2m_books_authors" ("b_id","a_id") values (2,7);
Insert into "library"."m2m_books_authors" ("b_id","a_id") values (3,2);
Insert into "library"."m2m_books_authors" ("b_id","a_id") values (4,3);
Insert into "library"."m2m_books_authors" ("b_id","a_id") values (4,6);
Insert into "library"."m2m_books_authors" ("b_id","a_id") values (5,6);
Insert into "library"."m2m_books_authors" ("b_id","a_id") values (6,4);
Insert into "library"."m2m_books_authors" ("b_id","a_id") values (6,5);
Insert into "library"."m2m_books_authors" ("b_id","a_id") values (7,1);

Insert into "library"."m2m_books_genres" ("b_id","g_id") values (1,1);
Insert into "library"."m2m_books_genres" ("b_id","g_id") values (1,5);
Insert into "library"."m2m_books_genres" ("b_id","g_id") values (2,1);
Insert into "library"."m2m_books_genres" ("b_id","g_id") values (2,5);
Insert into "library"."m2m_books_genres" ("b_id","g_id") values (3,6);
Insert into "library"."m2m_books_genres" ("b_id","g_id") values (4,2);
Insert into "library"."m2m_books_genres" ("b_id","g_id") values (4,3);
Insert into "library"."m2m_books_genres" ("b_id","g_id") values (5,2);
Insert into "library"."m2m_books_genres" ("b_id","g_id") values (6,5);
Insert into "library"."m2m_books_genres" ("b_id","g_id") values (7,2);
Insert into "library"."m2m_books_genres" ("b_id","g_id") values (7,5);

Insert into "library"."subscribers" ("s_id","s_name") values (1,'Ivanov I.I.');
Insert into "library"."subscribers" ("s_id","s_name") values (2,'Petrov P.P.');
Insert into "library"."subscribers" ("s_id","s_name") values (3,'Sidorov S.S.');
Insert into "library"."subscribers" ("s_id","s_name") values (4,'Sidorov S.S.');

Insert into "library"."subscriptions" ("sb_id","sb_subscriber","sb_book","sb_start","sb_finish","sb_is_active") values (100,1,3,to_date('12-JAN-11','DD-MON-YY'),to_date('12-FEB-11','DD-MON-YY'),'N');
Insert into "library"."subscriptions" ("sb_id","sb_subscriber","sb_book","sb_start","sb_finish","sb_is_active") values (2,1,1,to_date('12-JAN-11','DD-MON-YY'),to_date('12-FEB-11','DD-MON-YY'),'N');
Insert into "library"."subscriptions" ("sb_id","sb_subscriber","sb_book","sb_start","sb_finish","sb_is_active") values (3,3,3,to_date('17-MAY-12','DD-MON-YY'),to_date('17-JUL-12','DD-MON-YY'),'Y');
Insert into "library"."subscriptions" ("sb_id","sb_subscriber","sb_book","sb_start","sb_finish","sb_is_active") values (42,1,2,to_date('11-JUN-12','DD-MON-YY'),to_date('11-AUG-12','DD-MON-YY'),'N');
Insert into "library"."subscriptions" ("sb_id","sb_subscriber","sb_book","sb_start","sb_finish","sb_is_active") values (57,4,5,to_date('11-JUN-12','DD-MON-YY'),to_date('11-AUG-12','DD-MON-YY'),'N');
Insert into "library"."subscriptions" ("sb_id","sb_subscriber","sb_book","sb_start","sb_finish","sb_is_active") values (61,1,7,to_date('03-AUG-14','DD-MON-YY'),to_date('03-OCT-14','DD-MON-YY'),'N');
Insert into "library"."subscriptions" ("sb_id","sb_subscriber","sb_book","sb_start","sb_finish","sb_is_active") values (62,3,5,to_date('03-AUG-14','DD-MON-YY'),to_date('03-OCT-14','DD-MON-YY'),'Y');
Insert into "library"."subscriptions" ("sb_id","sb_subscriber","sb_book","sb_start","sb_finish","sb_is_active") values (86,3,1,to_date('03-AUG-14','DD-MON-YY'),to_date('03-SEP-14','DD-MON-YY'),'Y');
Insert into "library"."subscriptions" ("sb_id","sb_subscriber","sb_book","sb_start","sb_finish","sb_is_active") values (91,4,1,to_date('07-OCT-15','DD-MON-YY'),to_date('07-MAR-15','DD-MON-YY'),'Y');
Insert into "library"."subscriptions" ("sb_id","sb_subscriber","sb_book","sb_start","sb_finish","sb_is_active") values (95,1,4,to_date('07-OCT-15','DD-MON-YY'),to_date('07-NOV-15','DD-MON-YY'),'N');
Insert into "library"."subscriptions" ("sb_id","sb_subscriber","sb_book","sb_start","sb_finish","sb_is_active") values (99,4,4,to_date('08-OCT-15','DD-MON-YY'),to_date('08-NOV-25','DD-MON-YY'),'Y');


  
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
