-- name: LoginUser :one
SELECT uid, name, password FROM users WHERE email = $1;

-- name: SignupUser :one
INSERT INTO users (name,email,password) VALUES ($1,$2,$3) RETURNING uid ;

-- name: EmailTaken :one
SELECT EXISTS ( SELECT 1 FROM users WHERE email = $1 ) AS found;
