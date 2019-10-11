create table `matches`
(
  id bigint primary key not null auto_increment,
  status int,
  players int
);

create table `rounds`
(
  id bigint not null auto_increment,
  match_id bigint not null,
  `index` int not null,
  status int not null,
  team varchar(255) not null,
  state text,
  error text,

  created_at datetime(3) not null,
  updated_at datetime(3) not null,

  primary key (id),
  unique by_team_status (team,`index`)
);

create table `events`
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