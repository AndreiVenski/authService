package repository

const (
	writeRefreshTokenRecordQuery = `INSERT INTO refresh_tokens (user_id, refresh_token_id, hashed_token, expires, ip_addr) 
									VALUES ($1, $2, $3, $4, $5) `

	getRefreshTokenRecordQuery = `SELECT user_id, refresh_token_id, hashed_token, expires, ip_addr
								  FROM refresh_tokens 
								  WHERE refresh_token_id = $1`

	updateRefreshTokenIDQuery = `UPDATE refresh_tokens
									SET refresh_token_id = $1
									WHERE refresh_token_id = $2
									`

	getUserQuery = `SELECT user_id, name, email
					FROM users
					WHERE user_id = $1`
)
