-- name: CreateCv :exec
insert into cvs(id, uploaded_by_id, file_id) values($1, $2, $3);

-- name: GetCv :one
select *
from cvs
where id = $1
limit 1
;

-- name: GetAllCvs :many
select *
from cvs
limit $1
offset $2
;

-- name: GetFeaturesByCvs :many
select *
from cvs_features f
where f.cv_id = $1
;
