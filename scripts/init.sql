CREATE TABLE public.livros (
    id bigserial NOT NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL,
    titulo text NULL,
    categoria text NULL,
    autor text NULL,
    sinopse text NULL,
    CONSTRAINT livros_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_livros_deleted_at ON public.livros USING btree (deleted_at);
