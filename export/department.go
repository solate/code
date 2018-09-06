package basic

import (
	"encoding/json"
	"gitlab.dev.okapp.cc/med-his-group/medhis-platform-service/virhis"
	"io/ioutil"
	"net/url"
)

// 获取科室信息
func (s *VirHIS) Department(req *virhis.DepartmentReq, reply *virhis.DepartmentReply) (err error) {

	data := &url.Values{}
	data.Add("name", req.Name)
data.Add("age", req.Age)
	res, err := HttpGet(s.URL+URLDepartment, data)
	if err != nil {
		logger.Error("remote err: ", err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error("read body err: ", err)
		return
	}

	reply = new(virhis.DepartmentReply)
	err = json.Unmarshal(body, reply)
	if err != nil {
		logger.Error("unmarshal err: ", err)
		return
	}

	return
}