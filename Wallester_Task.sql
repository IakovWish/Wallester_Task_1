--
-- PostgreSQL database dump
--

-- Dumped from database version 14.4
-- Dumped by pg_dump version 14.4

-- Started on 2022-08-14 12:18:33

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE "Wallester_Task";
--
-- TOC entry 3409 (class 1262 OID 16448)
-- Name: Wallester_Task; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE "Wallester_Task" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'Russian_Russia.1251';


ALTER DATABASE "Wallester_Task" OWNER TO postgres;

\connect "Wallester_Task"

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 2 (class 3079 OID 16460)
-- Name: citext; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;


--
-- TOC entry 3410 (class 0 OID 0)
-- Dependencies: 2
-- Name: EXTENSION citext; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION citext IS 'data type for case-insensitive character strings';


--
-- TOC entry 872 (class 1247 OID 16566)
-- Name: email; Type: DOMAIN; Schema: public; Owner: postgres
--

CREATE DOMAIN public.email AS public.citext
	CONSTRAINT email_check CHECK ((VALUE OPERATOR(public.~) '^[a-zA-Z0-9.!#$%&''*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$'::public.citext));


ALTER DOMAIN public.email OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 210 (class 1259 OID 16596)
-- Name: customers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.customers (
    id integer NOT NULL,
    first_name character varying(100) NOT NULL,
    last_name character varying(100) NOT NULL,
    birth_date date NOT NULL,
    gender character varying(6) NOT NULL,
    e_mail public.email NOT NULL,
    address character varying(200),
    CONSTRAINT customers_birth_date_check CHECK (((date_part('year'::text, age((birth_date)::timestamp with time zone)) >= (18)::double precision) AND (date_part('year'::text, age((birth_date)::timestamp with time zone)) < (60)::double precision))),
    CONSTRAINT customers_gender_check CHECK ((((gender)::text = 'Male'::text) OR ((gender)::text = 'Female'::text)))
);


ALTER TABLE public.customers OWNER TO postgres;

--
-- TOC entry 211 (class 1259 OID 16607)
-- Name: customers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.customers ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.customers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 3402 (class 0 OID 16596)
-- Dependencies: 210
-- Data for Name: customers; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.customers OVERRIDING SYSTEM VALUE VALUES (20, 'Patricia', 'Ross', '2002-02-08', 'Female', 'dean26@yahoo.com', '3674 Ward Harbors');
INSERT INTO public.customers OVERRIDING SYSTEM VALUE VALUES (21, 'Patricia', 'Baker', '1977-02-26', 'Female', 'billy76@hotmail.com', '');
INSERT INTO public.customers OVERRIDING SYSTEM VALUE VALUES (22, 'Paul', 'Hunt', '1969-12-31', 'Male', 'frances.mohr@gmail.com', '459 Yvette Corner');
INSERT INTO public.customers OVERRIDING SYSTEM VALUE VALUES (23, 'Sandra', 'Willis', '1981-03-26', 'Female', 'DemonicFlag@yopmail.com', '7087 Stroman Square Apt. 011');
INSERT INTO public.customers OVERRIDING SYSTEM VALUE VALUES (24, 'Juan', 'Hill', '1982-12-12', 'Male', 'trent.mante@langosh.net', '');
INSERT INTO public.customers OVERRIDING SYSTEM VALUE VALUES (25, 'Carol', 'Lopez', '1989-07-05', 'Male', 'antonietta.feil@beahan.com', '');
INSERT INTO public.customers OVERRIDING SYSTEM VALUE VALUES (26, 'Joanne', 'Brown', '1983-01-18', 'Female', 'cheathcote@yahoo.com', '6794 Ethel Mountains Suite 558');
INSERT INTO public.customers OVERRIDING SYSTEM VALUE VALUES (27, 'Albert', 'Evans', '1979-04-15', 'Male', 'dora.bauch@yahoo.com', '98344 Amari Knoll');
INSERT INTO public.customers OVERRIDING SYSTEM VALUE VALUES (28, 'Ben', 'Adams', '1972-07-24', 'Male', 'vbogan@senger.com', '7129 Lakin Mountains Apt. 272');
INSERT INTO public.customers OVERRIDING SYSTEM VALUE VALUES (29, 'Ernest', 'Baker', '1986-04-20', 'Male', 'brionna.durgan@gmail.com', '');
INSERT INTO public.customers OVERRIDING SYSTEM VALUE VALUES (30, 'Mary', 'Lambert', '1962-09-29', 'Female', 'zachariah.grant@gmail.com', '870 Mikel Brook');
INSERT INTO public.customers OVERRIDING SYSTEM VALUE VALUES (31, 'Alice', 'McDaniel', '1980-10-14', 'Female', 'rice.gustave@rodriguez.com', '');


--
-- TOC entry 3411 (class 0 OID 0)
-- Dependencies: 211
-- Name: customers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.customers_id_seq', 31, true);


--
-- TOC entry 3262 (class 2606 OID 16604)
-- Name: customers customers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customers
    ADD CONSTRAINT customers_pkey PRIMARY KEY (id);


-- Completed on 2022-08-14 12:18:34

--
-- PostgreSQL database dump complete
--

