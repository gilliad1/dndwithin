CREATE TABLE public.spells (
    id integer NOT NULL,
    name text,
    school character varying(25),
    sub_school character varying(25),
    source character varying(25),
    description text,
    effect text,
    range character varying(50),
    components character varying(25),
    area character varying(25),
    target character varying(100),
    duration character varying(25),
    saving_throw character varying(25),
    spell_resistance boolean DEFAULT false,
    is_30 boolean DEFAULT false,
    is_evil boolean DEFAULT false
);

ALTER TABLE public.spells OWNER TO "svc-dndwithin";
ALTER TABLE ONLY public.spells ADD CONSTRAINT spells_pkey PRIMARY KEY (id);
REVOKE CONNECT,TEMPORARY ON DATABASE dndwithin FROM PUBLIC;
REVOKE ALL ON DATABASE dndwithin FROM "svc-dndwithin";
GRANT ALL ON DATABASE dndwithin TO "svc-dndwithin" WITH GRANT OPTION;