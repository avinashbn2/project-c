package repository

import "database/sql"

type ResourceRepo struct {
	db *sql.DB
}

func NewResourceRepo(db *sql.DB) *ResourceRepo {
	return &ResourceRepo{
		db: db,
	}
}
//FindAll : Select all from the table
func (rp *ResourceRepo) FindAll() (models.Resources, error) {
	query := `Select * from resource_item`
	var rs Resources
	err := rp.db.Select(rs, query)
	if err != nil {
		return nil, err
	}
	return rs, nil
}