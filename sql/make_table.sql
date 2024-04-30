create table graph
(
    id       integer not null
        constraint graph_pk
            primary key,
    name     varchar not null,
    graph_id varchar not null
);

create table node
(
    id           integer not null
        constraint nodes_pk
            primary key,
    name         varchar not null,
    node_id      varchar not null,
    ref_graph_id integer not null
);

create table edge
(
    id               integer          not null
        constraint edge_pk
            primary key,
    edge_id          varchar          not null,
    "from"           varchar          not null,
    ref_from_node_id integer          not null,
    "to"             varchar          not null,
    ref_to_node_id   integer          not null,
    cost             double precision not null
);
