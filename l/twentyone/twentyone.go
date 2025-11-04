package twentyone

import "fmt"

type Notifier interface {
	Notify(msg string)
}

type Old struct{}

func (p *Old) Print(text string) {
	fmt.Println("old:", text)
}

type PrinterAdapter struct {
	printer *Old
}

func NewPrinterAdapter(p *Old) Notifier {
	return &PrinterAdapter{printer: p}
}

func (a *PrinterAdapter) Notify(msg string) {
	a.printer.Print(msg)
}

func twentyOneTask() {
	var notifier = NewPrinterAdapter(&Old{})
	notifier.Notify("адаптер работает")
}
