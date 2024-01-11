create table person (
    id bigserial not null,
    full_name text not null,
    height bigint,
    age bigint,
    primary key (id)
);

create table movie (
    id bigserial not null,
    title text,
    movie_year bigint,
    tagline text,
    duration bigint,
    rating float,
    budget bigint,
    gross bigint,
    primary key (id)
);

create table professions (
    id bigserial not null,
    movie_id bigint not null,
    person_id bigint not null,
    movie_role bigint not null,
    primary key (id),
    foreign key (movie_id) references movie (id) on delete cascade,
    foreign key (person_id) references person (id) on delete cascade
);