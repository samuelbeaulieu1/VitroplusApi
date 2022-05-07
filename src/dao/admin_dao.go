package dao

type AdminDao struct {
	Dao
}

func NewAdminDao() *AdminDao {
	dao := &AdminDao{}
	connectDb(&dao.Dao)

	return dao
}
