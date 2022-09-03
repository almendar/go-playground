package adt

type Account interface {
	seal()
}

type Regular struct{}

func (Regular) seal() {}

type Premium struct{}

func (Premium) seal() {}

type Family struct {
	users []Regular
}

func (Family) seal() {}

type UsersProcessor[A any] interface {
	onUser(Regular) A
	onPremiumuser(Premium) A
	onFamilyAccount(Family) A
}

func Fold[A any](
	acc Account,
	up UsersProcessor[A],
) A {
	switch t := acc.(type) {
	case Regular:
		return up.onUser(t)
	case Premium:
		return up.onPremiumuser(t)
	case Family:
		return up.onFamilyAccount(t)
	default:
		var a A
		return a
	}
}

type PaymentCalcualtor struct{}

func (it PaymentCalcualtor) onUser(u Regular) int {
	return 25
}
func (it PaymentCalcualtor) onPremiumuser(pu Premium) int {
	return 43
}
func (it PaymentCalcualtor) onFamilyAccount(fa Family) int {
	users := len(fa.users)
	switch users {
	case 2:
		return 39
	case 3:
		return 49
	default:
		return 55
	}
}
