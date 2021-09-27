package testLogic

import "uapply_go/entity/ResponseModels"

func Pong(p *ResponseModels.Pong) (err error) {
	// mysql.Pong(...)
	p.Msg = "pong"
	return err
}
