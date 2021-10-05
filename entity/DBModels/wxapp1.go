package DBModels

type StuMessage1 struct {
	Uid      int    `json:"uid" db:"uid"`
	Openid   string `json:"openid" db:"openid"`
	Position int    `json:"position" db:"position"`
}
