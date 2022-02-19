CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

ALTER TABLE public.commands
    ADD CONSTRAINT commands_pk PRIMARY KEY (uuid);
ALTER TABLE public.requests
    ADD CONSTRAINT requests_pk PRIMARY KEY (uuid);

ALTER TABLE public.requests
    ALTER COLUMN uuid SET DEFAULT uuid_generate_v4();
ALTER TABLE public.requests
    ALTER COLUMN create_at SET DEFAULT now();
ALTER TABLE public.requests
    ALTER COLUMN update_at SET DEFAULT now();

CREATE TABLE public.commands_requests
(
    uuid          uuid        NOT NULL DEFAULT uuid_generate_v4(),
    create_at     timestamptz NOT NULL DEFAULT now(),
    update_at     timestamptz NOT NULL DEFAULT now(),
    commands_uuid uuid        NOT NULL,
    requests_uuid uuid        NOT NULL,
    "order"       int8        NOT NULL,
    CONSTRAINT commands_requests_pk PRIMARY KEY (uuid),
    CONSTRAINT commands_requests_fk FOREIGN KEY (commands_uuid) REFERENCES public.commands (uuid) ON DELETE SET NULL ON UPDATE CASCADE,
    CONSTRAINT commands_requests_fk_1 FOREIGN KEY (requests_uuid) REFERENCES public.requests (uuid) ON DELETE SET NULL ON UPDATE CASCADE
);
