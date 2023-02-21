package user

type (
	sUser struct{}
)

func init() {

}

func New() *sUser {
	return &sUser{}
}
