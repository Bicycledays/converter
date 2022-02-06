package service

const firstPort = 44000

type Manager struct {
	Unoservers map[int]*Unoserver
	Reserve    int
}

func CreateManager(r int) *Manager {
	u := make(map[int]*Unoserver)
	p := firstPort

	for i := 0; i < r; i++ {
		u[p] = &Unoserver{
			port:   p,
			status: statusSleep,
		}
		p++
	}

	return &Manager{
		Unoservers: u,
		Reserve:    r,
	}
}
