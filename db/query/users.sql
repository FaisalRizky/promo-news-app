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

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUsers :one
UPDATE users
SET  
  email = $2,
  name= $3,
  username= $4,
  password= $5,
  password_changed_at= $6,
  phone_number= $7,
  device_token= $8,
  lang= $9,
  avatar= $10,
  user_level= $11,
  is_active= $12
WHERE id = $1
RETURNING *;

-- name: ToogleActiveUsers :one
UPDATE users
SET is_active = $2
WHERE id = $1
RETURNING *;