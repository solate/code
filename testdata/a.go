package testdata

//科室信息， reply时list
type DepartmentReq struct {
	Name string `json:"name"` //名称
	Age  string `json:"age"`  //年龄
}
type DepartmentReply struct {
	List []Department
}
type Department struct {
	Code        string `json:"code"`        //科室编码
	Name        string `json:"name"`        //名称
	ParentId    string `json:"parent_id"`   //父类科室ID
}



// 医生个人信息查询, reply时结构
type DoctorInfo struct {
	Name              string `json:"name"`//姓名
	Phone             string `json:"phone"`//电话
}
type MedicalCardReply struct {
	Name string `json:"name"`//姓名
	Sex string `json:"sex"` //性别
	IdCard string `json:"id_card"` //身份证
}