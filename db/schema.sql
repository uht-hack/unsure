create table `match`
(
  id bigint primary key,
  status int
);

create table `round`
(
  id bigint primary key,
  status int,
  match_id bigint, -- foreign key
  data text
);
