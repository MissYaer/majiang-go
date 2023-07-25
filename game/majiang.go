package game

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

// 固定牌
var Pai = [34]string{
	"wan1", "wan2", "wan3", "wan4", "wan5", "wan6", "wan7", "wan8", "wan9",
	"tong1", "tong2", "tong3", "tong4", "tong5", "tong6", "tong7", "tong8", "tong9",
	"tiao1", "tiao2", "tiao3", "tiao4", "tiao5", "tiao6", "tiao7", "tiao8", "tiao9",
	"dong", "nan", "xi", "bei", "zhong", "fa", "bai",
}

var zong_pai = []string{} // 剩余的总牌

// json 返回数据
type GameJson struct {
	Dong
	Nan
	Xi
	Bei
	Status   string `json:"_status"`
	Types    string `json:"type"`
	Location string `json:"location"`
}

type MoPaiJson struct {
	UserName string `json:"user_name"`
	Pai      string `json:"pai"`
	GameType string `json:"type"`
}

type Dong struct {
	Nameid string   `json:"dong_name"`
	Pai    []string `json:"dong_pai"`
}

type Nan struct {
	Nameid string   `json:"nan_name"`
	Pai    []string `json:"nan_pai"`
}

type Xi struct {
	Nameid string   `json:"xi_name"`
	Pai    []string `json:"xi_pai"`
}

type Bei struct {
	Nameid string   `json:"bei_name"`
	Pai    []string `json:"bei_pai"`
}

// 当前摸牌人,默认东家
var current_mopai_user string = "dong"

var user_dong Dong
var user_nan Nan
var user_xi Xi
var user_bei Bei

// 初始化所有牌
func create_pai() {
	for _, v := range Pai {
		i := 0
		for i < 4 {
			zong_pai = append(zong_pai, v)
			i++
		}
	}
}

func create_init_user() (g_jsons string) {
	user_dong.Nameid = "wangyang"
	user_dong.Pai = get_init_pai(13)

	user_nan.Nameid = "chenhui"
	user_nan.Pai = get_init_pai(13)

	user_xi.Nameid = "xi"
	user_xi.Pai = get_init_pai(13)

	user_bei.Nameid = "bei"
	user_bei.Pai = get_init_pai(13)

	fmt.Println("%T", user_dong)

	g := GameJson{Status: "success", Types: "kaiju", Location: "1", Dong: user_dong, Nan: user_nan, Xi: user_xi, Bei: user_bei}
	g_json, _ := json.Marshal(g)
	return string(g_json)
}

func user_mopai() (g_jsons string) {
	mopai := Mopai(current_mopai_user)
	g := MoPaiJson{UserName: current_mopai_user, Pai: mopai, GameType: "zimo"}
	g_json, _ := json.Marshal(g)
	return string(g_json)
}

// 初始化发牌
func get_init_pai(start_number int) []string {
	fapai := []string{}
	i := 0
	for i < start_number {
		pai := get_random_pai()
		fapai = append(fapai, pai)
		i++
	}
	return fapai
}

// 随机获取一张牌
func get_random_pai() string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(zong_pai))

	value := zong_pai[index]
	index_after := index + 1
	zong_pai = append(zong_pai[:index], zong_pai[index_after:]...)

	return value
}
