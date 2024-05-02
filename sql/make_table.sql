create table if not exists graph
(
    id       integer generated always as identity
        constraint graph_pk
            primary key,
    name     varchar not null,
    graph_id varchar not null
);

create table if not exists node
(
    id           integer generated always as identity
        constraint nodes_pk
            primary key,
    name         varchar not null,
    node_id      varchar not null,
    ref_graph_id integer not null
        constraint node_graph_id_fk
            references graph
);

create table if not exists edge
(
    id               integer generated always as identity
        constraint edge_pk
            primary key,
    edge_id          varchar          not null,
    "from"           varchar          not null,
    ref_from_node_id integer          not null
        constraint edge_from_node_id_fk
            references node,
    "to"             varchar          not null,
    ref_to_node_id   integer          not null
        constraint edge_to_node_id_fk
            references node,
    cost             double precision not null,
    ref_graph_id     integer          not null
        constraint edge_graph_id_fk
            references graph
);
