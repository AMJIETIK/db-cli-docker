--
-- PostgreSQL database dump
--

-- Dumped from database version 15.12 (Homebrew)
-- Dumped by pg_dump version 15.12 (Homebrew)

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
-- Name: users; Type: TABLE; Schema: public; Owner: stormside7
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name text NOT NULL,
    email text NOT NULL,
    date_registered timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.users OWNER TO stormside7;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: stormside7
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO stormside7;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: stormside7
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: stormside7
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: stormside7
--

COPY public.users (id, name, email, date_registered) FROM stdin;
1	Vlad	vladpopov432@gmail.com	2025-04-23 20:14:19.562837
2	Ilya Sokolov	ilyaMJTV@gmail.com	2025-04-23 21:18:17.910995
4	Anna Ivanova	anna.ivanova@gmail.com	2024-01-12 00:00:00
5	John Smith	john.smith@gmail.com	2024-02-18 00:00:00
6	Maria Garcia	maria.garcia@gmail.com	2024-03-05 00:00:00
7	Liam Brown	liam.brown@gmail.com	2024-03-21 00:00:00
8	Sophia Wilson	sophia.wilson@gmail.com	2024-04-01 00:00:00
9	James Johnson	james.johnson@gmail.com	2024-04-10 00:00:00
10	Emily Davis	emily.davis@gmail.com	2024-04-12 00:00:00
11	Michael Miller	michael.miller@gmail.com	2024-04-14 00:00:00
14	Olivia Thomas	olivia.thomas@gmail.com	2024-04-18 00:00:00
15	Benjamin Taylor	benjamin.taylor@gmail.com	2024-04-20 00:00:00
16	Charlotte Moore	charlotte.moore@gmail.com	2024-04-21 00:00:00
17	Ethan Jackson	ethan.jackson@gmail.com	2024-04-22 00:00:00
20	Diana Karas	dianakar@gmail.com	2025-05-03 18:08:12.86209
21	maciej biernat	maciek111@gmail.com	2025-05-08 10:57:28.102072
\.


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: stormside7
--

SELECT pg_catalog.setval('public.users_id_seq', 26, true);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: stormside7
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: stormside7
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

