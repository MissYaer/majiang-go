package game

/*
*
1.游戏开始 先创建所有麻将
2.初始化所有用户
3.
*/
func Start() string {
	create_pai()

	g_json := create_init_user()
	/*for _, v := range user_nan.Pai {
		println(v)
	}*/
	return g_json
}

func UserMopai() string {
	g_json := user_mopai()
	return g_json
}
