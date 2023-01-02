-- name: CreateStores :one
INSERT INTO stores (
  name,
  address,
  description,
  phone_number,
  operational_id,
  is_active
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetStores :one
SELECT * FROM stores
WHERE id = $1 LIMIT 1;

-- name: ListStores :many
SELECT * FROM stores
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateStores :one
UPDATE stores
SET name = $2,
    address = $3,
    description= $4,
    phone_number = $5,
    operational_id = $6,
    is_active = $7
WHERE id = $1
RETURNING *;

-- name: ToogleActiveStores :one
UPDATE stores
SET is_active = $2
WHERE id = $1
RETURNING *;