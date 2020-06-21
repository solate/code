package {{.Package}}

import (
{{- range $key, $value := .ImportList}}
  "{{$value}}"
{{- end}}
)


type {{.Method}}Req struct {
    PatientId string `form:"patientId" json:"patientId" binding:"required"` //患者id
    PatientName string `form:"patientName" json:"patientName" binding:"omitempty"` //患者姓名
{{- range $key, $value := .ImportList}}
  "{{$value}}"
{{- end}}
}

type {{.Method}}Reply struct {
	List []PatientInfo
}

//{{.Note}}
func (*{{.Struct}}) {{.Method}}(c *moduleman.Context, req *{{.Method}}Req, reply *{{.Method}}Reply) *moduleman.Error {

    //TODO 代码

    doctorInfoReq := &virhis.DoctorInfoReq{
        DoctorId : req.DoctorId,
        DoctorName : req.DoctorName,
    }
    doctorInfoReply := new(virhis.DoctorInfoReply)
    err := his.DoctorInfo(doctorInfoReq, doctorInfoReply)
    if err != nil {
        logger.Error("call DoctorInfo err, please remote", err)
        moduleman.NewError(std.AdminCodeServerError, err.Error())
    }
    reply.DoctorName = doctorInfoReply.DoctorName
    reply.DoctorAge = doctorInfoReply.DoctorAge

	return nil
}
