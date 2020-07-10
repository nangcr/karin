package main

type clanApiResponse struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Ts   int64       `json:"ts"`
	Full int64       `json:"full"`
}

type Line struct {
	Rank           int64  `json:"rank"`
	Damage         int64  `json:"damage"`
	ClanName       string `json:"clan_name"`
	MemberNum      int64  `json:"member_num"`
	LeaderName     string `json:"leader_name"`
	LeaderViewerID int64  `json:"leader_viewer_id"`
}

type clanApiLine struct {
	clanApiResponse
	Data []Line `json:"data"`
}
