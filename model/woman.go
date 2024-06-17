package model

type Woman struct {
	Men // Hereda de hombre

}

func (m *Woman) Respirar()    { m.respirando = true }
func (m *Woman) Pensar()      { m.comiendo = true }
func (m *Woman) Comer()       { m.pensando = true }
func (m *Woman) Sexo() string { return "Mujer" }
