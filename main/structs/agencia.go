package structs

type Agencia struct {
	modelo string
	autos  []Car
}

func NewAgencia(nombreModelo string) Agencia {
	return Agencia{
		modelo: nombreModelo,
	}
}

func (a *Agencia) AgregarAuto(autoNuevo Car) {
	a.autos = append(a.autos, autoNuevo)
}

func (a *Agencia) BorrarAuto(auto Car) {
	//
}
