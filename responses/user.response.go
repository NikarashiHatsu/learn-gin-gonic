package responses

import "time"

type UserResponse struct {
	ID        *int       `json:"id"`
	Name      *string    `json:"name"`
	Address   *string    `json:"address"`
	Email     *string    `json:"email"`
	BornDate  *time.Time `json:"born_date"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
