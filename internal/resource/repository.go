package resource

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

type ResourceResponse struct {
	Resources Resources `json:"resources"`
	Total     int       `json:"total" db:"total"`
}

func NewRepo(database *sqlx.DB) *repository {
	return &repository{
		db: database,
	}
}

func (rp *repository) getByUserLikes(qp *QueryWithUser) (*ResourceResponse, error) {
	query := fmt.Sprintf(`select title, url, author, image,excerpt,created_at, resource_item.updated_at  ,  true as like_by_user,  (select count(*) from
	user_resource where resource_item.id=user_resource.rid) as likes  from resource_item inner join user_resource on
	user_resource.rid = resource_item.id  and user_resource.uid=%d  order by %s %s limit %d offset %d`, qp.user.ID, qp.sortParam, qp.sortOrder, qp.limit, qp.offset)
	var rs Resources
	err := rp.db.Select(&rs, query)
	if err != nil {
		log.Fatal("Err", err)
		return nil, err

	}
	ctQuery := `Select count(*)  from resource_item `

	response := &ResourceResponse{}
	var rsp int
	err = rp.db.Get(&rsp, ctQuery)
	if err != nil {
		fmt.Println("errror ", err)
		log.Fatal("ERORR", err)
		return nil, err
	}
	response.Resources = rs
	response.Total = rsp / qp.limit
	return response, nil
}

func (rp *repository) getByTrending(qp *QueryWithUser) (*ResourceResponse, error) {
	query := fmt.Sprintf(`Select title, url, author, image,excerpt,created_at, resource_item.updated_at, (select count(*) from user_resource where resource_item.id=user_resource.rid )
as likes   from resource_item inner join user_resource on user_resource.rid = resource_item.id and user_resource.updated_at > now() - interval '1' day  order by likes desc Limit
%d offset %d`, qp.limit, qp.offset)
	var rs Resources
	err := rp.db.Select(&rs, query)
	if err != nil {
		log.Fatal("Err", err)
		return nil, err

	}
	ctQuery := `Select count(*)  from resource_item `

	response := &ResourceResponse{}
	var rsp int
	err = rp.db.Get(&rsp, ctQuery)
	if err != nil {
		fmt.Println("errror ", err)
		log.Fatal("ERORR", err)
		return nil, err
	}
	response.Resources = rs
	response.Total = rsp / qp.limit
	return response, nil
}

func (rp *repository) get(qp *QueryWithUser) (*ResourceResponse, error) {

	query := fmt.Sprintf(`Select *, exists( select * from user_resource ur
					 where  ur.rid=resource_item.id and ur.uid= %d
					) as like_by_user,
					(select count(*) from user_resource where resource_item.id=user_resource.rid) as likes   from resource_item order by %s %s limit %d offset %d`, qp.user.ID, qp.sortParam, qp.sortOrder, qp.limit, qp.offset)
	var rs Resources
	err := rp.db.Select(&rs, query)
	if err != nil {
		log.Fatal("Err", err)
		return nil, err

	}
	ctQuery := `Select count(*)  from resource_item `

	response := &ResourceResponse{}
	var rsp int
	err = rp.db.Get(&rsp, ctQuery)
	if err != nil {
		fmt.Println("errror ", err)
		log.Fatal("ERORR", err)
		return nil, err
	}
	response.Resources = rs
	response.Total = rsp / qp.limit
	return response, nil
}

//Add : Post Resource data to table
func (rp *repository) create(ri *ResourceItem) (*ResourceItem, error) {

	query := `INSERT INTO resource_item( title, url, tag,sitename, author, image, excerpt, created_at, updated_at)
VALUES(:title, :url, :tag,:sitename, :author, :image, :excerpt, :created_at, :updated_at) returning id`
	rows, err := rp.db.NamedQuery(query, ri)
	var id int
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		rows.Scan(&id)
	}
	ri.ID = id
	return ri, nil
}

//FindByID : Select  from the table by id
func (rp *repository) getByID(id int) (*Resources, error) {

	query := fmt.Sprintf("Select * from resource_item where id=%d", id)
	var rs Resources
	err := rp.db.Select(&rs, query)
	if err != nil {
		log.Fatal(err)
		return nil, err

	}
	return &rs, nil
}

func (rp *repository) update(resource *ResourceItem) (*ResourceItem, error) {

	query := `UPDATE resource_item SET title=:title, url=:url, tag=:tag,updated_at=:updated_at  WHERE id=:id`
	_, err := rp.db.NamedExec(query, resource)
	if err != nil {
		log.Fatal(err, "UPDATE")
		return nil, err
	}
	return resource, nil
}
