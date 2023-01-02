-- name: CreateOperationalTime :one
INSERT INTO operational_time (
  opening_time,
  closing_time,
  operational_days,
  off_days,
  is_active
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetOperationalTime :one
SELECT * FROM operational_time
WHERE id = $1 LIMIT 1;

-- name: ListOperationalTime :many
SELECT * FROM operational_time
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateOperationalTime :one
UPDATE operational_time
SET  
  opening_time=$2,
  closing_time=$3,
  operational_days=$4,
  off_days=$5,
  is_active=$6
WHERE id = $1
RETURNING *;

-- name: ToogleActiveOperationalTime :one
UPDATE operational_time
SET is_active = $2
WHERE id = $1
RETURNING *;