ALTER TABLE public.commands DROP CONSTRAINT commands_pk;
ALTER TABLE public.requests DROP CONSTRAINT requests_pk;

ALTER TABLE public.requests ALTER COLUMN uuid DROP DEFAULT;
ALTER TABLE public.requests ALTER COLUMN create_at DROP DEFAULT;
ALTER TABLE public.requests ALTER COLUMN update_at DROP DEFAULT;

DROP TABLE IF EXISTS public.commands_requests;
