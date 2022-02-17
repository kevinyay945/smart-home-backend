CREATE TABLE public.commands
(
    uuid      uuid NULL DEFAULT uuid_generate_v4(),
    create_at timestamptz NULL DEFAULT now(),
    update_at timestamptz NULL DEFAULT now(),
    url       varchar NULL
);
CREATE TABLE public.requests
(
    uuid      uuid        NOT NULL,
    create_at timestamptz NOT NULL,
    update_at timestamptz NOT NULL,
    "name"    varchar     NOT NULL
);
