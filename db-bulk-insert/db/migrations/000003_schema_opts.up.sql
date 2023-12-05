CREATE MATERIALIZED VIEW IF NOT EXISTS public.view_common_user AS (
 SELECT e.name,
    e.short_description,
    e.classificacao_etaria,
    eo.starts_on,
    eo.starts_at,
    eo.ends_at,
    s.name AS space_name,
    s.location
   FROM events e
     JOIN event_occurrences eo ON e.id = eo.event
     JOIN spaces s ON eo.space = s.id
     LEFT JOIN projects p ON e.project = p.id )
WITH DATA;

CREATE INDEX IF NOT EXISTS idx_view_common_classificacao_etaria
    ON public.view_common_user USING btree (classificacao_etaria);

CREATE INDEX IF NOT EXISTS idx_view_common_nome_espaco
    ON public.view_common_user USING btree (name);

CREATE INDEX IF NOT EXISTS idx_view_common_nome_evento
    ON public.view_common_user USING btree (space_name);
    
CREATE INDEX IF NOT EXISTS idx_view_common_starts_on
    ON public.view_common_user USING btree (starts_on);


CREATE MATERIALIZED VIEW IF NOT EXISTS public.view_produtor
AS (
 SELECT e.name,
    e.short_description,
    e.classificacao_etaria,
    eo.starts_on,
    eo.starts_at,
    eo.ends_at,
    s.name AS space_name,
    s.location,
    s.short_description AS sd_space,
    s.telefone AS telefone_espaco,
    s.email AS email_espaco,
    s.horario AS horario_funcionamento,
    p.name AS project_name,
    p.short_description AS project_short_desc,
    p.registration_from,
    p.registration_to,
    a.name AS dono_evento,
    asp.name AS dono_espaco,
    ap.name AS dono_projeto,
    s.terms,
    spp.name AS parent_space,
    prp.name AS parent_project,
    agp.name AS parent_agent
   FROM events e
     JOIN event_occurrences eo ON e.id = eo.event
     JOIN spaces s ON eo.space = s.id
     LEFT JOIN projects p ON e.project = p.id
     LEFT JOIN agents a ON e.owner = a.id
     LEFT JOIN agents asp ON s.owner = asp.id
     LEFT JOIN agents ap ON p.owner = ap.id
     LEFT JOIN spaces spp ON s.parent = spp.id
     LEFT JOIN projects prp ON p.parent = prp.id
     LEFT JOIN agents agp ON a.parent = agp.id
)
WITH DATA;

CREATE INDEX IF NOT EXISTS idx_view_produtor_dono_espaco
    ON public.view_produtor USING btree (dono_espaco);

CREATE INDEX IF NOT EXISTS idx_view_produtor_dono_projeto
    ON public.view_produtor USING btree (dono_projeto);

CREATE INDEX IF NOT EXISTS idx_view_produtor_nome_projeto
    ON public.view_produtor USING btree (project_name);
