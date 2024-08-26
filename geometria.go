package geometria

type Triangulo struct {
	Base   float64
	Altura float64
}
type Cuadrado struct {
	Lado float64
}
type Pentagono struct {
	Lado    float64
	Apotema float64
}

func (t *Triangulo) Area() float64 {
	return t.Base * t.Altura / 2
}
func (c *Cuadrado) Area() float64 {
	return c.Lado * c.Lado
}
func (p *Pentagono) Area() float64 {
	per := p.perimetro()
	return per * p.Apotema / 2
}
func (p *Pentagono) perimetro() float64 {
	return p.Lado * 5
}

func Figura(lados string) string {
	if lados == "3" {
		return "Triangulo"
	} else if lados == "4" {
		return "Cuadrado"
	} else if lados == "5" {
		return "Pentagono"
	} else {
		return "Valor no valido"
	}
}
