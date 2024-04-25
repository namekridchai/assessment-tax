package incomeTaxAllowance

func GetKrcp() float64 {

	var krcp float64
	row := db.QueryRow("SELECT k_receipt FROM public.allowance_master")
	row.Scan(&krcp)
	return krcp

}
