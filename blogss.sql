--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.15
-- Dumped by pg_dump version 9.6.15

-- Started on 2020-04-06 21:20:55

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
-- TOC entry 1 (class 3079 OID 12387)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 2156 (class 0 OID 0)
-- Dependencies: 1
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 185 (class 1259 OID 24695)
-- Name: articles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.articles (
    id bigint NOT NULL,
    title character varying(255) NOT NULL,
    content character varying NOT NULL,
    created_at timestamp with time zone NOT NULL,
    is_published boolean NOT NULL,
    published_at timestamp with time zone,
    updated_at timestamp with time zone,
    id_user integer NOT NULL
);


ALTER TABLE public.articles OWNER TO postgres;

--
-- TOC entry 186 (class 1259 OID 24701)
-- Name: articles_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.articles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.articles_id_seq OWNER TO postgres;

--
-- TOC entry 2157 (class 0 OID 0)
-- Dependencies: 186
-- Name: articles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.articles_id_seq OWNED BY public.articles.id;


--
-- TOC entry 190 (class 1259 OID 24720)
-- Name: messages; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.messages (
    id bigint NOT NULL,
    name character varying(50) NOT NULL,
    email character varying(100) NOT NULL,
    message text NOT NULL
);


ALTER TABLE public.messages OWNER TO postgres;

--
-- TOC entry 189 (class 1259 OID 24718)
-- Name: messages_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.messages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.messages_id_seq OWNER TO postgres;

--
-- TOC entry 2158 (class 0 OID 0)
-- Dependencies: 189
-- Name: messages_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.messages_id_seq OWNED BY public.messages.id;


--
-- TOC entry 187 (class 1259 OID 24703)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    is_admin boolean DEFAULT false NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 188 (class 1259 OID 24710)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 2159 (class 0 OID 0)
-- Dependencies: 188
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 2016 (class 2604 OID 24712)
-- Name: articles id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.articles ALTER COLUMN id SET DEFAULT nextval('public.articles_id_seq'::regclass);


--
-- TOC entry 2019 (class 2604 OID 24723)
-- Name: messages id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.messages ALTER COLUMN id SET DEFAULT nextval('public.messages_id_seq'::regclass);


--
-- TOC entry 2018 (class 2604 OID 24713)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 2143 (class 0 OID 24695)
-- Dependencies: 185
-- Data for Name: articles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.articles (id, title, content, created_at, is_published, published_at, updated_at, id_user) FROM stdin;
2	Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.	Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Purus sit amet volutpat consequat mauris nunc congue nisi. Cursus eget nunc scelerisque viverra mauris. Sed vulputate odio ut enim blandit volutpat maecenas volutpat. Dictumst vestibulum rhoncus est pellentesque elit ullamcorper. Ipsum suspendisse ultrices gravida dictum. Faucibus vitae aliquet nec ullamcorper. Quis vel eros donec ac odio tempor orci. Tristique senectus et netus et malesuada. Imperdiet proin fermentum leo vel orci porta non pulvinar neque. Ac tortor dignissim convallis aenean et tortor. Turpis egestas pretium aenean pharetra magna ac placerat vestibulum lectus. Sit amet aliquam id diam maecenas. Ornare massa eget egestas purus viverra accumsan in. Consectetur lorem donec massa sapien faucibus et molestie. Consequat id porta nibh venenatis cras sed felis eget velit.\n\nMattis enim ut tellus elementum. Euismod nisi porta lorem mollis. Nunc sed blandit libero volutpat sed cras ornare arcu. Ac odio tempor orci dapibus ultrices. Mattis ullamcorper velit sed ullamcorper. Tellus id interdum velit laoreet id donec ultrices tincidunt arcu. At augue eget arcu dictum varius duis at consectetur lorem. Eget dolor morbi non arcu. Id semper risus in hendrerit gravida. Diam vel quam elementum pulvinar etiam. Sit amet luctus venenatis lectus magna. Eget felis eget nunc lobortis mattis aliquam faucibus. Sed id semper risus in hendrerit gravida rutrum quisque. Pellentesque habitant morbi tristique senectus et netus. Lacus luctus accumsan tortor posuere ac ut consequat. Aliquet lectus proin nibh nisl.	2020-04-05 17:45:53.048905+07	t	2020-04-06 21:00:17.857335+07	\N	4
\.


--
-- TOC entry 2160 (class 0 OID 0)
-- Dependencies: 186
-- Name: articles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.articles_id_seq', 5, true);


--
-- TOC entry 2148 (class 0 OID 24720)
-- Dependencies: 190
-- Data for Name: messages; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.messages (id, name, email, message) FROM stdin;
1	Test	test@mail.com	Test
\.


--
-- TOC entry 2161 (class 0 OID 0)
-- Dependencies: 189
-- Name: messages_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.messages_id_seq', 1, true);


--
-- TOC entry 2145 (class 0 OID 24703)
-- Dependencies: 187
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, password, name, is_admin) FROM stdin;
4	teza	$2a$14$LKttdG3K.pBI.fdGnA1EC.gHqhmZMIHNojFAbBHQKUjkzNPoNupxu	Teza	t
\.


--
-- TOC entry 2162 (class 0 OID 0)
-- Dependencies: 188
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 5, true);


--
-- TOC entry 2021 (class 2606 OID 24715)
-- Name: articles articles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.articles
    ADD CONSTRAINT articles_pkey PRIMARY KEY (id);


--
-- TOC entry 2025 (class 2606 OID 24728)
-- Name: messages messages_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT messages_pkey PRIMARY KEY (id);


--
-- TOC entry 2023 (class 2606 OID 24717)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


-- Completed on 2020-04-06 21:20:55

--
-- PostgreSQL database dump complete
--

