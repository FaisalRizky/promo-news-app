-- name: CreateUsers :one
INSERT INTO users (
  email,
  name,
  username,
  password,
  password_changed_at,
  phone_number,
  device_token,
  lang,
  avatar,
  user_level,
  is_active
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING *;

-- name: GetUsers :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUsers :one
UPDATE users
SET  
  email = $1,
  name= $2,
  username= $3,
  password= $4,
  password_changed_at= $5,
  phone_number= $6,
  device_token= $7,
  lang= $8,
  avatar= $9,
  user_level= $10,
  is_active= $11
RETURNING *;

-- name: ToogleActiveUsers :one
UPDATE users
SET is_active = $2
WHERE id = $1
RETURNING *;