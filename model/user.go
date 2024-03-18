package model

type (
	MstUser struct {
		ID        string `json:"id,omitempty"`
		Name      string `json:"name,omitempty"`
		Email     string `json:"email,omitempty"`
		Password  string `json:"password,omitempty"`
		CreatedAt int    `json:"created_at,omitempty"`
		UpdatedAt int    `json:"updated_at,omitempty"`
		DeletedAt int    `json:"deleted_at,omitempty"`
	}
	GetListUserRequest struct {
		Page  int32
		Limit int32
	}
	GetListUserResponse struct {
		Message string
		Data    []MstUser
		Total   int64
	}
)
