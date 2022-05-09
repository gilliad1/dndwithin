CREATE TABLE IF NOT EXISTS public.spells (
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

CREATE TABLE IF NOT EXISTS public.characters (
    id integer NOT NULL,
    name text,
    hitPoints integer,
    speed integer,
    ac integer,
    touchAC integer,
    flatFootAC integer,
    classes text,
    strength integer,
    dexterity integer,
    constitution integer,
    intelligence integer,
    wisdom integer,
    charisma integer,
    initiativeMod integer,
    fortitudeSave integer,
    reflexSave integer,
    willSave integer,
    attackHandheld integer,
    attackMissile integer,
    grappleCheck integer,
    carryingLight integer,
    carryingMedium integer,
    carryingHeavy integer,
    carryLiftOverhead integer,
    carryLiftOffGrouhd integer,
    carryPushDrag integer,
    languages text,
    feats text,
    skills text,
    background text,
    description text,
    mostRecent boolean DEFAULT true
)
ALTER TABLE public.characters OWNER TO "svc-dndwithin";
ALTER TABLE ONLY public.characters ADD CONSTRAINT characters_pkey PRIMARY KEY (id);
