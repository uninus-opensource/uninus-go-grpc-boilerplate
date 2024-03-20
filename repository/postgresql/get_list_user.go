package postgresql

import (
	"context"
	"github.com/uninus-opensource/uninus-go-grpc-boilerplate/model"
)

func (r *dbReadWriter) GetListUser(ctx context.Context, params model.GetListUserRequest) (*model.GetListUserResponse, error) {
	var users []model.MstUser

	const defaultLimit = 10 // Set default limit pagination
	if params.Limit == 0 {
		params.Limit = defaultLimit
	}
	if params.Page <= 0 {
		params.Page = 1
	}
	offset := (params.Page - 1) * params.Limit

	//Query get user with pagination
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, email, created_at, updated_at FROM mst_user LIMIT $1 OFFSET $2", params.Limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// maaping user data
	for rows.Next() {
		var user model.MstUser
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	//count total
	var total int64
	err = r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM mst_user").Scan(&total)
	if err != nil {
		return nil, err
	}
	response := &model.GetListUserResponse{
		Message: "List of users",
		Data:    users,
		Total:   total,
	}
	return response, nil
}
