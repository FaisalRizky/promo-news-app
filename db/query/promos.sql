-- name: CreatePromos :one
INSERT INTO promos (
  promo_name,
  store_id ,
  promo_code,
  promo_description,
  quantity,
  start_at,
  expired_at,
  is_active,
  created_by
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: GetPromos :one
SELECT * FROM promos
WHERE id = $1 LIMIT 1;

-- name: ListPromos :many
SELECT * FROM promos
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdatePromos :one
UPDATE promos
SET  
  promo_name =  $2,
  store_id  =  $3,
  promo_code =  $4,
  promo_description = $5, 
  quantity =  $6,
  start_at =  $7,
  expired_at =  $8,
  is_active =  $9,
  created_by =  $10
WHERE id = $1
RETURNING *;

-- name: ToogleActivePromos :one
UPDATE promos
SET is_active = $2
WHERE id = $1
RETURNING *;