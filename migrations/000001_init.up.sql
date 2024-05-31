create type statuses as enum('waiting', 'ok');

create table if not exists cvs (
  id text PRIMARY KEY,
  uploaded_by_id text not null,  
  file_id text not null,
  status statuses not null default 'waiting'
);

create table if not exists cvs_features (
  cv_id text not null,
  feature_id int not null,
  value text not null,
  FOREIGN KEY(cv_id) REFERENCES cvs(id)
);
