create table `match`
(
  id bigint primary key,
  status int
);

create table `match_events`
(
    id bigint not null auto_increment,
    foreign_id bigint not null,
    timestamp datetime not null,
    type int not null,

    primary key (id)
);

create table `round`
(
  id bigint primary key,
  status int,
  match_id bigint, -- foreign key
  data text
);

create table `round_events`
(
  id bigint not null auto_increment,
  foreign_id bigint not null,
  timestamp datetime not null,
  type int not null,

  primary key (id)
);

create table `uht_cursors`
(
    id varchar(255) not null,
    last_event_id bigint not null default 0,
    updated_at datetime not null,

    primary key (id)
)