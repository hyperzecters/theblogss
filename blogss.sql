PGDMP                         x            blogss    9.6.16    9.6.16     [           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                       false            \           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                       false            ]           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                       false            ^           1262    16447    blogss    DATABASE     �   CREATE DATABASE blogss WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'English_Indonesia.1252' LC_CTYPE = 'English_Indonesia.1252';
    DROP DATABASE blogss;
             postgres    false                        2615    2200    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
             postgres    false            _           0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                  postgres    false    3                        3079    12387    plpgsql 	   EXTENSION     ?   CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;
    DROP EXTENSION plpgsql;
                  false            `           0    0    EXTENSION plpgsql    COMMENT     @   COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';
                       false    1            �            1259    16470    articles    TABLE     S  CREATE TABLE public.articles (
    id bigint NOT NULL,
    title character varying(255) NOT NULL,
    content character varying NOT NULL,
    created_at timestamp with time zone NOT NULL,
    is_published boolean NOT NULL,
    published_at timestamp with time zone,
    updated_at timestamp with time zone,
    id_user integer NOT NULL
);
    DROP TABLE public.articles;
       public         postgres    false    3            �            1259    16468    articles_id_seq    SEQUENCE     x   CREATE SEQUENCE public.articles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.articles_id_seq;
       public       postgres    false    3    188            a           0    0    articles_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.articles_id_seq OWNED BY public.articles.id;
            public       postgres    false    187            �            1259    16450    users    TABLE     �   CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    is_admin boolean DEFAULT false NOT NULL
);
    DROP TABLE public.users;
       public         postgres    false    3            �            1259    16448    users_id_seq    SEQUENCE     u   CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.users_id_seq;
       public       postgres    false    186    3            b           0    0    users_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
            public       postgres    false    185            �           2604    16473    articles id    DEFAULT     j   ALTER TABLE ONLY public.articles ALTER COLUMN id SET DEFAULT nextval('public.articles_id_seq'::regclass);
 :   ALTER TABLE public.articles ALTER COLUMN id DROP DEFAULT;
       public       postgres    false    187    188    188            �           2604    16453    users id    DEFAULT     d   ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);
 7   ALTER TABLE public.users ALTER COLUMN id DROP DEFAULT;
       public       postgres    false    185    186    186            X          0    16470    articles 
   TABLE DATA                     public       postgres    false    188   �       c           0    0    articles_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.articles_id_seq', 1, false);
            public       postgres    false    187            V          0    16450    users 
   TABLE DATA                     public       postgres    false    186   �       d           0    0    users_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.users_id_seq', 4, true);
            public       postgres    false    185            �           2606    16478    articles articles_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.articles
    ADD CONSTRAINT articles_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.articles DROP CONSTRAINT articles_pkey;
       public         postgres    false    188    188            �           2606    16458    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public         postgres    false    186    186            X   
   x���          V   �   x���v
Q���W((M��L�+-N-*V��L�Q 1�sSu
����b~fq|bJnf��B��O�k������zIjU�:�V1JT14Qq�s�K)��p�2�L-6+�ʈ��K�5���1���JJu)H�2��,
.)t�/v�M2���SRT��i��� t�4     