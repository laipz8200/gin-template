package entity

type User struct {
	id     string
	amount float64
}

func (u User) ID() string { return u.id }

func (u User) Amount() float64 { return u.amount }

func (u *User) AddAmount(amount float64) {
	u.amount += amount
}

func (u *User) ReduceAmount(amount float64) {
	u.amount -= amount
}

func NewUser(id string, amount float64) User {
	return User{
		id:     id,
		amount: amount,
	}
}
