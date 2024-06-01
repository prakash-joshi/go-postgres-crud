--
-- PostgreSQL database dump
--

-- Dumped from database version 16.2
-- Dumped by pg_dump version 16.2

-- Started on 2024-06-01 15:40:30


-- Database: stocksdb

-- DROP DATABASE IF EXISTS stocksdb;

CREATE DATABASE stocksdb
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'English_United States.1252'
    LC_CTYPE = 'English_United States.1252'
    LOCALE_PROVIDER = 'libc'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;


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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 216 (class 1259 OID 24960)
-- Name: stocks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.stocks (
    stockid integer NOT NULL,
    name text,
    price integer,
    company text
);


ALTER TABLE public.stocks OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 24959)
-- Name: stocks_stockid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.stocks_stockid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.stocks_stockid_seq OWNER TO postgres;

--
-- TOC entry 4839 (class 0 OID 0)
-- Dependencies: 215
-- Name: stocks_stockid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.stocks_stockid_seq OWNED BY public.stocks.stockid;


--
-- TOC entry 4688 (class 2604 OID 24963)
-- Name: stocks stockid; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stocks ALTER COLUMN stockid SET DEFAULT nextval('public.stocks_stockid_seq'::regclass);


--
-- TOC entry 4690 (class 2606 OID 24967)
-- Name: stocks stocks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stocks
    ADD CONSTRAINT stocks_pkey PRIMARY KEY (stockid);


-- Completed on 2024-06-01 15:40:31

--
-- PostgreSQL database dump complete
--

