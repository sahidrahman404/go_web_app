create table snippets(
    id bigserial primary key,
    title varchar(100) not null,
    content text not null,
    created timestamptz default current_timestamp,
    expires timestamptz not null,
);
