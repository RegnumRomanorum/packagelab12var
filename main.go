package main

import (
	"packagelab12var/mobiletariff"

	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
)

func main() {
	callsCost, err := mobiletariff.CallCost(300, 5)
	if err != nil {
		color.Red("Ошибка при расчете стоимости звонков: %v", err)
		return
	}

	internetCost, err := mobiletariff.InternetCost(50, 100)
	if err != nil {
		color.Red("Ошибка при расчете стоимости интернета: %v", err)
		return
	}

	total := callsCost + internetCost
	err = mobiletariff.ApplyPackageDiscount(&total, 10)
	if err != nil {
		color.Red("Ошибка при применении скидки: %v", err)
		return
	}

	report, err := mobiletariff.FormatTariffReport("Сагымбаев Ж.Н.", callsCost, internetCost, total)
	if err != nil {
		color.Red("Ошибка при формировании отчета: %v", err)
		return
	}

	color.Cyan("Отчёт по тарифу:")
	color.Yellow(report)
	color.Green("Итоговая сумма после скидки: %s", humanize.Ftoa(total))
}
