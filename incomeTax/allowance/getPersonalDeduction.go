package incomeTaxAllowance

func GetPersonalDeduction() float64 {
	var personalDeduction float64
	row := db.QueryRow("SELECT personal FROM public.allowance_master")
	row.Scan(&personalDeduction)
	return personalDeduction

}
