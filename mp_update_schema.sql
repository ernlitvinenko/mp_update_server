--
-- PostgreSQL database dump
--

-- Dumped from database version 17.2
-- Dumped by pg_dump version 17.3 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: add_version(character varying, character varying, character varying); Type: PROCEDURE; Schema: public; Owner: mp_update
--

CREATE PROCEDURE public.add_version(IN id_app character varying, IN ver_id character varying, IN app_link character varying)
    LANGUAGE plpgsql
    AS $$
    declare max_ver_code integer;
begin
    select into max_ver_code max(version_code) from version where app_id = id_app;
    insert into version (id, app_id, version_code, link) values (ver_id, id_app, max_ver_code + 1, app_link);
end;
$$;


ALTER PROCEDURE public.add_version(IN id_app character varying, IN ver_id character varying, IN app_link character varying) OWNER TO mp_update;

--
-- Name: add_version(character varying, character varying, character varying, character varying); Type: PROCEDURE; Schema: public; Owner: mp_update
--

CREATE PROCEDURE public.add_version(IN id_app character varying, IN ver_id character varying, IN app_link character varying, IN description character varying)
    LANGUAGE plpgsql
    AS $$
declare max_ver_code integer;
begin
    select into max_ver_code max(version_code) from version where app_id = id_app;
    insert into version (id, app_id, version_code, link, description) values (ver_id, id_app, max_ver_code + 1, app_link, description);
end;
$$;


ALTER PROCEDURE public.add_version(IN id_app character varying, IN ver_id character varying, IN app_link character varying, IN description character varying) OWNER TO mp_update;

--
-- Name: create_app(character varying, character varying, character varying); Type: PROCEDURE; Schema: public; Owner: mp_update
--

CREATE PROCEDURE public.create_app(IN id_app character varying, IN name character varying, IN ver_link character varying)
    LANGUAGE plpgsql
    AS $$
begin
    insert into application values (id_app, name);
    insert into version (app_id, link) values (id_app, ver_link);
end;
$$;


ALTER PROCEDURE public.create_app(IN id_app character varying, IN name character varying, IN ver_link character varying) OWNER TO mp_update;

--
-- Name: create_app(character varying, character varying, character varying, character varying); Type: PROCEDURE; Schema: public; Owner: mp_update
--

CREATE PROCEDURE public.create_app(IN id_app character varying, IN name character varying, IN ver_link character varying, IN description character varying)
    LANGUAGE plpgsql
    AS $$
begin
    insert into application(id, app_name, description) values (id_app, name, description);
    insert into version (app_id, link, description) values (id_app, ver_link, description);
end;
$$;


ALTER PROCEDURE public.create_app(IN id_app character varying, IN name character varying, IN ver_link character varying, IN description character varying) OWNER TO mp_update;

--
-- Name: create_app(character varying, character varying, character varying, character varying, character varying); Type: PROCEDURE; Schema: public; Owner: mp_update
--

CREATE PROCEDURE public.create_app(IN id_app character varying, IN name character varying, IN ver_link character varying, IN description character varying, IN version_id character varying)
    LANGUAGE plpgsql
    AS $$
begin
    insert into application(id, app_name, description) values (id_app, name, description);
    insert into version (id, app_id, link, description) values (version_id, id_app, ver_link, description);
end;
$$;


ALTER PROCEDURE public.create_app(IN id_app character varying, IN name character varying, IN ver_link character varying, IN description character varying, IN version_id character varying) OWNER TO mp_update;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: application; Type: TABLE; Schema: public; Owner: mp_update
--

CREATE TABLE public.application (
    id character varying(128) NOT NULL,
    app_name character varying(64) NOT NULL,
    description character varying(512)
);


ALTER TABLE public.application OWNER TO mp_update;

--
-- Name: profile; Type: TABLE; Schema: public; Owner: mp_update
--

CREATE TABLE public.profile (
    id integer NOT NULL,
    username character varying(64) NOT NULL,
    password character varying(64) NOT NULL
);


ALTER TABLE public.profile OWNER TO mp_update;

--
-- Name: profile_id_seq; Type: SEQUENCE; Schema: public; Owner: mp_update
--

CREATE SEQUENCE public.profile_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.profile_id_seq OWNER TO mp_update;

--
-- Name: profile_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mp_update
--

ALTER SEQUENCE public.profile_id_seq OWNED BY public.profile.id;


--
-- Name: version; Type: TABLE; Schema: public; Owner: mp_update
--

CREATE TABLE public.version (
    id character varying(32) DEFAULT '0.0.0'::character varying NOT NULL,
    app_id character varying(128) NOT NULL,
    version_code integer DEFAULT 0 NOT NULL,
    link character varying(2048) NOT NULL,
    description character varying(512) DEFAULT NULL::character varying
);


ALTER TABLE public.version OWNER TO mp_update;

--
-- Name: version_id_seq; Type: SEQUENCE; Schema: public; Owner: mp_update
--

CREATE SEQUENCE public.version_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.version_id_seq OWNER TO mp_update;

--
-- Name: version_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mp_update
--

ALTER SEQUENCE public.version_id_seq OWNED BY public.version.id;


--
-- Name: profile id; Type: DEFAULT; Schema: public; Owner: mp_update
--

ALTER TABLE ONLY public.profile ALTER COLUMN id SET DEFAULT nextval('public.profile_id_seq'::regclass);


--
-- Name: application application_pkey; Type: CONSTRAINT; Schema: public; Owner: mp_update
--

ALTER TABLE ONLY public.application
    ADD CONSTRAINT application_pkey PRIMARY KEY (id);


--
-- Name: profile profile_pkey; Type: CONSTRAINT; Schema: public; Owner: mp_update
--

ALTER TABLE ONLY public.profile
    ADD CONSTRAINT profile_pkey PRIMARY KEY (id);


--
-- Name: version version_pkey; Type: CONSTRAINT; Schema: public; Owner: mp_update
--

ALTER TABLE ONLY public.version
    ADD CONSTRAINT version_pkey PRIMARY KEY (id, app_id);


--
-- Name: version version_app_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: mp_update
--

ALTER TABLE ONLY public.version
    ADD CONSTRAINT version_app_id_fkey FOREIGN KEY (app_id) REFERENCES public.application(id);


--
-- PostgreSQL database dump complete
--

