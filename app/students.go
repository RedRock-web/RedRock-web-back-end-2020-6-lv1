package app

import (
	"RedRock-web-back-end-2020-6-lv1/database"
	"encoding/json"
	"errors"
	"strconv"
)

// 获取学生信息
func GetStudentInfo(id int) {
	var sr database.StudentRec
	url := "http://jwc.cqupt.edu.cn/data/json_StudentSearch.php?searchKey=" + strconv.Itoa(id) + "&NhGeWgGN=4EAbxgvThC3E000aXv_5SOXXZa3yAlzUptju_FeS7VRUxHsX6VpWA42OEjc.KTkUE1oKDhnZW68XuW7ZPEy6F0pXvlzHX7ubWSv683GvN_ncVDiF067VF9YtJumUR7gtCw8E3APAkmHyXSjTMUgYbCw3i8AyH9MBRBIvcMM7w1mS2FyfX88Ephx8pHuiD7n5Ho4v7qiOzJg_cetQZSVSDz2dHF.E7SrOOuNfyCYvqbB4PrmD2OXNRq0BJOGNLaP0k_gUkr1mcjo4AmTsgEgKg5wx7UcgqVvyJ2mOQS_R78uAGcWhs93JtGOXilQVy.LHJXpO1eGCc.n.w6Icl2B7_80hw4X5U58JIPnGkNs86HKuqcX8N2IcufzRfe8M25ESAqrZ2Q5g7Kbg5utvHOnTg7rCYd3_uCjh72rWAmrqRtGoPppt_swcVABv8VItAcQFtxUQ"
	if err := json.Unmarshal([]byte(GetBody(url)), &sr); err != nil {
		errors.New("unmarashal json failed!")
	}
	database.G_db.Create(&database.Student{
		Grade:      sr.ReturnData[0].Nj,
		Name:       sr.ReturnData[0].Xm,
		StudentId:  sr.ReturnData[0].Xh,
		Gender:     sr.ReturnData[0].Xb,
		CalssId:    sr.ReturnData[0].Bj,
		Department: sr.ReturnData[0].Yxm,
		Profession: sr.ReturnData[0].Zym,
		Status:     sr.ReturnData[0].Xjzt,
		Nation:     sr.ReturnData[0].Mz,
	})
	ch2 <- 1
}

// 判断学生是否存在
func StudentIsExist(id int) bool {
	var sr database.StudentRec

	url := "http://jwc.cqupt.edu.cn/data/json_StudentSearch.php?searchKey=" + strconv.Itoa(id) + "&NhGeWgGN=4EAbxgvThC3E000aXv_5SOXXZa3yAlzUptju_FeS7VRUxHsX6VpWA42OEjc.KTkUE1oKDhnZW68XuW7ZPEy6F0pXvlzHX7ubWSv683GvN_ncVDiF067VF9YtJumUR7gtCw8E3APAkmHyXSjTMUgYbCw3i8AyH9MBRBIvcMM7w1mS2FyfX88Ephx8pHuiD7n5Ho4v7qiOzJg_cetQZSVSDz2dHF.E7SrOOuNfyCYvqbB4PrmD2OXNRq0BJOGNLaP0k_gUkr1mcjo4AmTsgEgKg5wx7UcgqVvyJ2mOQS_R78uAGcWhs93JtGOXilQVy.LHJXpO1eGCc.n.w6Icl2B7_80hw4X5U58JIPnGkNs86HKuqcX8N2IcufzRfe8M25ESAqrZ2Q5g7Kbg5utvHOnTg7rCYd3_uCjh72rWAmrqRtGoPppt_swcVABv8VItAcQFtxUQ"
	if err := json.Unmarshal([]byte(GetBody(url)), &sr); err != nil {
		errors.New("unmarashal json failed!")
	}

	return len(sr.ReturnData) == 1
}

