--
-- PostgreSQL database dump
--

-- Dumped from database version 13.3 (Debian 13.3-1.pgdg100+1)
-- Dumped by pg_dump version 13.3

-- Started on 2021-07-04 16:29:46 WITA

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
-- TOC entry 3 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO postgres;

--
-- TOC entry 2944 (class 0 OID 0)
-- Dependencies: 3
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 200 (class 1259 OID 16397)
-- Name: account; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.account (
    account_number character varying,
    customer_number character varying,
    balance bigint
);


ALTER TABLE public.account OWNER TO postgres;

--
-- TOC entry 201 (class 1259 OID 16403)
-- Name: customers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.customers (
    customer_number character varying NOT NULL,
    name character varying
);


ALTER TABLE public.customers OWNER TO postgres;

--
-- TOC entry 2937 (class 0 OID 16397)
-- Dependencies: 200
-- Data for Name: account; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.account (account_number, customer_number, balance) FROM stdin;
550001	1001	94000
550002	1002	56000
\.


--
-- TOC entry 2938 (class 0 OID 16403)
-- Dependencies: 201
-- Data for Name: customers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.customers (customer_number, name) FROM stdin;
1001	Bob Martin
1002	Linux Trovalds
\.


-- Completed on 2021-07-04 16:29:47 WITA

--
-- PostgreSQL database dump complete
--

