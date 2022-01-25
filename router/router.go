package router

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"raffle/dao"
	"raffle/model"
	"time"
)

func Lottery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	award := r.FormValue("name")
	//var id int
	data, _ := dao.CheckAward(award)
	ret := new(model.JsonResult)
	msg, _ := json.Marshal(model.JsonResult{Code: 400, Msg: "该奖项次数已抽完！"})
	//if award == "七点一刻老老实实奖" {
	//	if data.ID < 1 {
	//		for i:=0;i<5;i++ {
	//			for true {
	//				rand.Seed(time.Now().UnixNano())
	//				id = rand.Intn(11) + 1 //随机生成一个1到11的数字
	//				data, _ = dao.CheckId(id)
	//				if data.Award == "" {
	//					dao.SaveData(id, award)
	//					data,_=dao.CheckId(id)
	//					break
	//				}
	//			}
	//			ret.Code=200
	//			ret.Msg=""
	//			ret.Data = append(ret.Data, model.Data{
	//				ID:     data.ID,
	//				Member: data.Member,
	//				Award:  award,
	//			})
	//		}
	//		msg, _ = json.Marshal(ret)
	//		io.WriteString(w, string(msg))
	//	} else {
	//		//t := template.Must(template.ParseFiles("views/prompt.html"))
	//		//t.Execute(w, nil)
	//		ret.Code=200
	//		ret.Msg="该奖项次数已抽完！"
	//		msg, _ = json.Marshal(ret)
	//		io.WriteString(w, string(msg))
	//	}
	//
	//} else {
	if data.ID < 1 {
		for true {
			rand.Seed(time.Now().UnixNano())
			id := rand.Intn(11) + 1 //随机生成一个1到11的数字
			data, _ = dao.CheckId(id)
			if data.Award == "" {
				dao.SaveData(id, award)
				data, _ = dao.CheckId(id)
				break
			}
		}
		ret.Code = 200
		ret.Msg = ""
		ret.Data = append(ret.Data, model.Data{
			ID:     data.ID,
			Member: data.Member,
			Award:  award,
		})
		msg, _ = json.Marshal(ret)
		io.WriteString(w, string(msg))
	} else {
		//t := template.Must(template.ParseFiles("views/prompt.html"))
		//t.Execute(w, nil)
		ret.Code = 200
		ret.Msg = "该奖项次数已抽完！"
		msg, _ = json.Marshal(ret)
		io.WriteString(w, string(msg))
	}
}

//}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	msg, _ := json.Marshal(model.JsonResult{Code: 400, Msg: "重置失败"})
	var id int
	for id = 1; id < 12; id++ {
		dao.DeleteData(id)
	}
	msg, _ = json.Marshal(model.JsonResult{Code: 200, Msg: "重置成功"})
	io.WriteString(w, string(msg))

}
