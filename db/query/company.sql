-- name: GetCompany :one
SELECT * FROM company
WHERE id = $1 LIMIT 1;

-- name: ListCompanies :many
SELECT * FROM company
ORDER BY name
LIMIT $1
OFFSET $2
;

-- name: CreateCompany :one
INSERT INTO company (
  name, 
  description,
  amount,
  registered,
  type
) VALUES (
  $1, $2,$3,$4,$5
)
RETURNING *;

-- name: UpdateCompany :one
UPDATE company
  set name = $2,
  description = $3,
  amount = $4,
  registered = $5,
  type = $6
WHERE id = $1
RETURNING *;



-- name: DeleteCompany :exec
DELETE FROM company
WHERE id = $1;