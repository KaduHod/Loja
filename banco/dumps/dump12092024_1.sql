--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4 (Debian 16.4-1.pgdg120+1)
-- Dumped by pg_dump version 16.4 (Debian 16.4-1.pgdg120+1)

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
-- Name: businesses_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.businesses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.businesses_id_seq OWNER TO root;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: businesses; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.businesses (
    id integer DEFAULT nextval('public.businesses_id_seq'::regclass) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    name character varying(255) NOT NULL,
    cnpj character varying(14) NOT NULL
);


ALTER TABLE public.businesses OWNER TO root;

--
-- Name: category; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.category (
    id integer NOT NULL,
    description character varying(50) NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.category OWNER TO root;

--
-- Name: category_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.category_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.category_id_seq OWNER TO root;

--
-- Name: category_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.category_id_seq OWNED BY public.category.id;


--
-- Name: file_category; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.file_category (
    id integer NOT NULL,
    file_id integer NOT NULL,
    category_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.file_category OWNER TO root;

--
-- Name: file_category_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.file_category_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.file_category_id_seq OWNER TO root;

--
-- Name: file_category_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.file_category_id_seq OWNED BY public.file_category.id;


--
-- Name: files; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.files (
    id integer NOT NULL,
    file_name character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    business_id integer,
    person_id integer,
    product_id integer,
    store_id integer,
    CONSTRAINT files_file_name_check CHECK ((length((file_name)::text) >= 4))
);


ALTER TABLE public.files OWNER TO root;

--
-- Name: files_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.files_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.files_id_seq OWNER TO root;

--
-- Name: files_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.files_id_seq OWNED BY public.files.id;


--
-- Name: person_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.person_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.person_id_seq OWNER TO root;

--
-- Name: person; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.person (
    id integer DEFAULT nextval('public.person_id_seq'::regclass) NOT NULL,
    name character varying(256),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    email character varying(255) NOT NULL
);


ALTER TABLE public.person OWNER TO root;

--
-- Name: product_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.product_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.product_id_seq OWNER TO root;

--
-- Name: product; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.product (
    id integer DEFAULT nextval('public.product_id_seq'::regclass) NOT NULL,
    name character varying(256) NOT NULL,
    description text,
    file_prefix character varying(50),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.product OWNER TO root;

--
-- Name: purchase_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.purchase_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.purchase_id_seq OWNER TO root;

--
-- Name: purchase; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.purchase (
    id integer DEFAULT nextval('public.purchase_id_seq'::regclass) NOT NULL,
    product_id integer,
    person_id integer,
    file_prefix character varying(50),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.purchase OWNER TO root;

--
-- Name: purchase_status_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.purchase_status_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.purchase_status_id_seq OWNER TO root;

--
-- Name: purchase_status; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.purchase_status (
    id integer DEFAULT nextval('public.purchase_status_id_seq'::regclass) NOT NULL,
    status character varying(50) NOT NULL,
    status_description text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.purchase_status OWNER TO root;

--
-- Name: purchase_status_mapping_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.purchase_status_mapping_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.purchase_status_mapping_id_seq OWNER TO root;

--
-- Name: purchase_status_mapping; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.purchase_status_mapping (
    id integer DEFAULT nextval('public.purchase_status_mapping_id_seq'::regclass) NOT NULL,
    purchase_id integer,
    status_id integer,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.purchase_status_mapping OWNER TO root;

--
-- Name: store; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.store (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    business_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.store OWNER TO root;

--
-- Name: store_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.store_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.store_id_seq OWNER TO root;

--
-- Name: store_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.store_id_seq OWNED BY public.store.id;


--
-- Name: store_product; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.store_product (
    id integer NOT NULL,
    store_id integer NOT NULL,
    product_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.store_product OWNER TO root;

--
-- Name: store_product_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.store_product_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.store_product_id_seq OWNER TO root;

--
-- Name: store_product_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.store_product_id_seq OWNED BY public.store_product.id;


--
-- Name: user_businesses_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.user_businesses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.user_businesses_id_seq OWNER TO root;

--
-- Name: user_businesses; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.user_businesses (
    person_id integer NOT NULL,
    business_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    id integer DEFAULT nextval('public.user_businesses_id_seq'::regclass) NOT NULL
);


ALTER TABLE public.user_businesses OWNER TO root;

--
-- Name: category id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.category ALTER COLUMN id SET DEFAULT nextval('public.category_id_seq'::regclass);


--
-- Name: file_category id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.file_category ALTER COLUMN id SET DEFAULT nextval('public.file_category_id_seq'::regclass);


--
-- Name: files id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.files ALTER COLUMN id SET DEFAULT nextval('public.files_id_seq'::regclass);


--
-- Name: store id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.store ALTER COLUMN id SET DEFAULT nextval('public.store_id_seq'::regclass);


--
-- Name: store_product id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.store_product ALTER COLUMN id SET DEFAULT nextval('public.store_product_id_seq'::regclass);


--
-- Data for Name: businesses; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.businesses (id, created_at, updated_at, name, cnpj) FROM stdin;
1	2024-09-09 23:24:51.134439	2024-09-09 23:24:51.134439	TESTE_BUSINESS	12345678912345
27	2024-09-11 02:40:42.480996	2024-09-11 02:40:42.480996	TESTE_CREATE_BUSINESS	12345678912346
\.


--
-- Data for Name: category; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.category (id, description, created_at) FROM stdin;
1	person	2024-09-11 19:15:24.287927
2	store	2024-09-11 19:15:24.287927
3	product	2024-09-11 19:15:24.287927
4	business	2024-09-11 19:15:24.287927
\.


--
-- Data for Name: file_category; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.file_category (id, file_id, category_id, created_at, updated_at) FROM stdin;
26	39	4	2024-09-12 11:26:58.476323	2024-09-12 11:26:58.476323
\.


--
-- Data for Name: files; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.files (id, file_name, created_at, updated_at, business_id, person_id, product_id, store_id) FROM stdin;
39	__test__.jpeg	2024-09-12 11:26:49.917448	2024-09-12 11:26:49.917448	1	\N	\N	\N
\.


--
-- Data for Name: person; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.person (id, name, created_at, updated_at, email) FROM stdin;
6	root	2024-09-09 22:25:41.88473	2024-09-09 22:25:41.88473	root@mail.com
\.


--
-- Data for Name: product; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.product (id, name, description, file_prefix, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: purchase; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.purchase (id, product_id, person_id, file_prefix, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: purchase_status; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.purchase_status (id, status, status_description, created_at, updated_at) FROM stdin;
1	Pending	The order has been received but not yet processed.	2024-09-08 15:26:58.237135	2024-09-08 15:26:58.237135
2	Confirmed	The order has been confirmed and is awaiting preparation or shipping.	2024-09-08 15:26:58.237135	2024-09-08 15:26:58.237135
3	Processing	The order is being processed.	2024-09-08 15:26:58.237135	2024-09-08 15:26:58.237135
4	Shipped	The order has been shipped and is on its way to the customer.	2024-09-08 15:26:58.237135	2024-09-08 15:26:58.237135
5	Delivered	The order has been delivered to the customer.	2024-09-08 15:26:58.237135	2024-09-08 15:26:58.237135
6	Completed	The order has been completed and the transaction is final.	2024-09-08 15:26:58.237135	2024-09-08 15:26:58.237135
7	Cancelled	The order has been cancelled.	2024-09-08 15:26:58.237135	2024-09-08 15:26:58.237135
8	Returned	The order has been returned by the customer.	2024-09-08 15:26:58.237135	2024-09-08 15:26:58.237135
9	Refunded	The order has been refunded.	2024-09-08 15:26:58.237135	2024-09-08 15:26:58.237135
10	Failed	The order processing failed.	2024-09-08 15:26:58.237135	2024-09-08 15:26:58.237135
\.


--
-- Data for Name: purchase_status_mapping; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.purchase_status_mapping (id, purchase_id, status_id, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: store; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.store (id, name, business_id, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: store_product; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.store_product (id, store_id, product_id, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: user_businesses; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.user_businesses (person_id, business_id, created_at, updated_at, id) FROM stdin;
6	1	2024-09-10 22:08:10.39547	2024-09-10 22:08:10.39547	2
6	1	2024-09-11 01:15:30.302676	2024-09-11 01:15:30.302676	3
6	1	2024-09-11 01:15:55.679107	2024-09-11 01:15:55.679107	4
6	1	2024-09-11 02:37:09.863775	2024-09-11 02:37:09.863775	5
6	1	2024-09-11 02:37:28.943564	2024-09-11 02:37:28.943564	6
6	1	2024-09-11 02:37:53.954092	2024-09-11 02:37:53.954092	7
6	1	2024-09-11 02:38:02.117829	2024-09-11 02:38:02.117829	8
6	1	2024-09-11 02:39:32.547148	2024-09-11 02:39:32.547148	9
6	1	2024-09-11 02:40:42.482619	2024-09-11 02:40:42.482619	10
\.


--
-- Name: businesses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.businesses_id_seq', 27, true);


--
-- Name: category_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.category_id_seq', 4, true);


--
-- Name: file_category_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.file_category_id_seq', 26, true);


--
-- Name: files_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.files_id_seq', 39, true);


--
-- Name: person_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.person_id_seq', 57, true);


--
-- Name: product_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.product_id_seq', 1, false);


--
-- Name: purchase_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.purchase_id_seq', 1, false);


--
-- Name: purchase_status_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.purchase_status_id_seq', 10, true);


--
-- Name: purchase_status_mapping_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.purchase_status_mapping_id_seq', 1, false);


--
-- Name: store_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.store_id_seq', 1, false);


--
-- Name: store_product_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.store_product_id_seq', 1, false);


--
-- Name: user_businesses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.user_businesses_id_seq', 10, true);


--
-- Name: businesses businesses_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.businesses
    ADD CONSTRAINT businesses_pkey PRIMARY KEY (id);


--
-- Name: category category_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.category
    ADD CONSTRAINT category_pkey PRIMARY KEY (id);


--
-- Name: file_category file_category_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.file_category
    ADD CONSTRAINT file_category_pkey PRIMARY KEY (id);


--
-- Name: files files_file_name_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.files
    ADD CONSTRAINT files_file_name_key UNIQUE (file_name);


--
-- Name: files files_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.files
    ADD CONSTRAINT files_pkey PRIMARY KEY (id);


--
-- Name: person person_email_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.person
    ADD CONSTRAINT person_email_key UNIQUE (email);


--
-- Name: person person_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.person
    ADD CONSTRAINT person_pkey PRIMARY KEY (id);


--
-- Name: product product_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product
    ADD CONSTRAINT product_pkey PRIMARY KEY (id);


--
-- Name: purchase purchase_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.purchase
    ADD CONSTRAINT purchase_pkey PRIMARY KEY (id);


--
-- Name: purchase_status_mapping purchase_status_mapping_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.purchase_status_mapping
    ADD CONSTRAINT purchase_status_mapping_pkey PRIMARY KEY (id);


--
-- Name: purchase_status purchase_status_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.purchase_status
    ADD CONSTRAINT purchase_status_pkey PRIMARY KEY (id);


--
-- Name: store store_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.store
    ADD CONSTRAINT store_pkey PRIMARY KEY (id);


--
-- Name: store_product store_product_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.store_product
    ADD CONSTRAINT store_product_pkey PRIMARY KEY (id);


--
-- Name: store_product store_product_unique; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.store_product
    ADD CONSTRAINT store_product_unique UNIQUE (product_id, store_id);


--
-- Name: businesses unique_cnpj; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.businesses
    ADD CONSTRAINT unique_cnpj UNIQUE (cnpj);


--
-- Name: file_category unique_file_category; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.file_category
    ADD CONSTRAINT unique_file_category UNIQUE (file_id, category_id);


--
-- Name: user_businesses user_businesses_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.user_businesses
    ADD CONSTRAINT user_businesses_pkey PRIMARY KEY (id);


--
-- Name: files fk_business; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.files
    ADD CONSTRAINT fk_business FOREIGN KEY (business_id) REFERENCES public.businesses(id);


--
-- Name: file_category fk_category; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.file_category
    ADD CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES public.category(id) ON DELETE CASCADE;


--
-- Name: file_category fk_file; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.file_category
    ADD CONSTRAINT fk_file FOREIGN KEY (file_id) REFERENCES public.files(id) ON DELETE CASCADE;


--
-- Name: store fk_negocio; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.store
    ADD CONSTRAINT fk_negocio FOREIGN KEY (business_id) REFERENCES public.businesses(id);


--
-- Name: files fk_person; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.files
    ADD CONSTRAINT fk_person FOREIGN KEY (person_id) REFERENCES public.person(id);


--
-- Name: files fk_product; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.files
    ADD CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES public.product(id);


--
-- Name: files fk_store; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.files
    ADD CONSTRAINT fk_store FOREIGN KEY (store_id) REFERENCES public.store(id);


--
-- Name: purchase purchase_person_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.purchase
    ADD CONSTRAINT purchase_person_id_fkey FOREIGN KEY (person_id) REFERENCES public.person(id);


--
-- Name: purchase purchase_product_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.purchase
    ADD CONSTRAINT purchase_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.product(id);


--
-- Name: purchase_status_mapping purchase_status_mapping_purchase_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.purchase_status_mapping
    ADD CONSTRAINT purchase_status_mapping_purchase_id_fkey FOREIGN KEY (purchase_id) REFERENCES public.purchase(id);


--
-- Name: purchase_status_mapping purchase_status_mapping_status_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.purchase_status_mapping
    ADD CONSTRAINT purchase_status_mapping_status_id_fkey FOREIGN KEY (status_id) REFERENCES public.purchase_status(id);


--
-- Name: user_businesses user_businesses_business_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.user_businesses
    ADD CONSTRAINT user_businesses_business_id_fkey FOREIGN KEY (business_id) REFERENCES public.businesses(id) ON DELETE CASCADE;


--
-- Name: user_businesses user_businesses_person_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.user_businesses
    ADD CONSTRAINT user_businesses_person_id_fkey FOREIGN KEY (person_id) REFERENCES public.person(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

