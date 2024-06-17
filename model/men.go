package model

type Men struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	vivo       bool
}

func (h *Men) Respirar()    { h.respirando = true }
func (h *Men) Pensar()      { h.comiendo = true }
func (h *Men) Comer()       { h.pensando = true }
func (h *Men) Sexo() string { return "Hombre" }
