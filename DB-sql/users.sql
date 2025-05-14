create table public.users (
  id bigserial not null,
  username character varying(255) not null,
  email text null,
  email_verified_at timestamp without time zone null,
  password character varying(255) not null,
  remember_token character varying(100) null,
  created_at timestamp without time zone null,
  updated_at timestamp without time zone null,
  dni text null,
  telefono text null,
  direccion text null,
  tipo_sangre text null,
  role text null,
  constraint users_pkey primary key (id),
  constraint users_email_unique unique (email)
) TABLESPACE pg_default;