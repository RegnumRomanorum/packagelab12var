// packagelab12var/mobiletariff/tariffs.go
// Пакет для расчета стоимости мобильного тарифа, включающий функции для расчета стоимости звонков, интернета, применения скидок и формирования отчетов.
package mobiletariff

import (
	"errors"
	"fmt"
)

// CallCost рассчитывает стоимость звонков на основе количества минут и цены за минуту.
func CallCost(minutes float64, pricePerMin float64) (float64, error) {
	if minutes < 0 {
		return 0, errors.New("Минуты не могут быть отрицательными")
	}
	if pricePerMin < 0 {
		return 0, errors.New("Цена не может быть отрицательной")
	}
	return minutes * pricePerMin, nil
}

// InternetCost рассчитывает стоимость интернета на основе количества гигабайт и цены за гигабайт.
func InternetCost(gb float64, pricePerGb float64) (float64, error) {
	if gb < 0 {
		return 0, errors.New("Гигабайты не могут быть отрицательными")
	}
	if pricePerGb < 0 {
		return 0, errors.New("Цена за гигабайт не может быть отрицательной")
	}
	return gb * pricePerGb, nil
}

// ApplyPackageDiscount применяет скидку к общей стоимости тарифа.
func ApplyPackageDiscount(total *float64, percent float64) error {
	if percent < 0 || percent > 100 {
		return errors.New("Процент скидки должен быть от 0 до 100")
	}
	*total = *total - (*total * percent / 100)
	return nil
}

// FormatTariffReport формирует отчет о тарифе для пользователя, включая стоимость звонков, интернета и итоговую стоимость.
func FormatTariffReport(user string, calls, internet, total float64) (string, error) {
	report := "Пользователь: " + user + "\n"
	report += "Стоимость звонков: " + fmt.Sprintf("%.2f", calls) + "\n"
	report += "Стоимость интернета: " + fmt.Sprintf("%.2f", internet) + "\n"
	report += "Итоговая стоимость: " + fmt.Sprintf("%.2f", total)
	return report, nil
}
