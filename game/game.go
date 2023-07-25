package game

// 东家摸牌
func Mopai(user_type string) string {
	pai := get_random_pai()
	switch user_type {
	case "dong":
		user_dong.Pai = append(user_dong.Pai, pai)
		current_mopai_user = "nan"
	case "nan":
		user_nan.Pai = append(user_nan.Pai, pai)
		current_mopai_user = "xi"
	case "xi":
		user_xi.Pai = append(user_xi.Pai, pai)
		current_mopai_user = "bei"
	case "bei":
		user_bei.Pai = append(user_bei.Pai, pai)
		current_mopai_user = "dong"
	}
	return pai
}
