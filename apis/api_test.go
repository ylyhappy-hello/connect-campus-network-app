package apis

import (
	"testing"
)

// 这个单元测试依赖本地网络环境， 没办法测试
func TestGetQueryString(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{name: testing.CoverMode(), want: "wlanuserip=7d44873d9e9a46a1ef0c11db72554ad5&wlanacname=cee24f8cac61aac974e74b503c3a1669&ssid=&nasip=8618724f96768bd61edb375c4c21ccc0&snmpagentip=&mac=f95d27987c9488911905c9ac9062e704&t=wireless-v2&url=ee87b1634742a905d3bcaa94ab6f72ecb6810fb1e1bcd8cb&apmac=&nasid=cee24f8cac61aac974e74b503c3a1669&vid=5d55317a73fbd6fb&port=991a6a4279d5c31d&nasportid=85280d0376cea3455469dda3beba88f5f9c8893793420b7c9481fa7ceeb292ca065abdfb8031177b", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetQueryString()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetQueryString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetQueryString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCampusNetService(t *testing.T) {
	type args struct {
		userId      string
		queryString string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: testing.CoverMode(), args: args{userId: "202114600207", queryString: "wlanuserip=7d44873d9e9a46a1ef0c11db72554ad5&wlanacname=cee24f8cac61aac974e74b503c3a1669&ssid=&nasip=8618724f96768bd61edb375c4c21ccc0&snmpagentip=&mac=f95d27987c9488911905c9ac9062e704&t=wireless-v2&url=ee87b1634742a905d3bcaa94ab6f72ecb6810fb1e1bcd8cb&apmac=&nasid=cee24f8cac61aac974e74b503c3a1669&vid=5d55317a73fbd6fb&port=991a6a4279d5c31d&nasportid=85280d0376cea3455469dda3beba88f5f9c8893793420b7c9481fa7ceeb292ca065abdfb8031177b"}, want: "校园网@联通专线", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCampusNetService(tt.args.userId, tt.args.queryString)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCampusNetService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetCampusNetService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnectCampusNetWork(t *testing.T) {
	type args struct {
		cl *CampusLogin
	}
	tests := []struct {
		name    string
		args    args
		want    ConnectCampusNetWorkResp
		wantErr bool
	}{
		{name: testing.CoverMode(), args: args{cl: &CampusLogin{UserId: "202114600207", Password: "qwefsdd435fgsdH@H@", Service: "联通专线", QueryString: "wlanuserip=7d44873d9e9a46a1ef0c11db72554ad5&wlanacname=cee24f8cac61aac974e74b503c3a1669&ssid=&nasip=8618724f96768bd61edb375c4c21ccc0&snmpagentip=&mac=f95d27987c9488911905c9ac9062e704&t=wireless-v2&url=ee87b1634742a905d3bcaa94ab6f72ecb6810fb1e1bcd8cb&apmac=&nasid=cee24f8cac61aac974e74b503c3a1669&vid=5d55317a73fbd6fb&port=991a6a4279d5c31d&nasportid=85280d0376cea3455469dda3beba88f5f9c8893793420b7c9481fa7ceeb292ca065abdfb8031177b"}},
			want: ConnectCampusNetWorkResp{UserIndex: "38363138373234663936373638626436316564623337356334633231636363305f31302e3133312e3138322e3133385f323032313134363030323037",Result: "success", Message: "",Forwordurl: nil, KeepaliveInterval: 0,ValidCodeURL: ""}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConnectCampusNetWork(tt.args.cl)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConnectCampusNetWork() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConnectCampusNetWork() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogoutCampusNetWork(t *testing.T) {
	type args struct {
		userIndex string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: testing.CoverMode(), args: args{userIndex: "38363138373234663936373638626436316564623337356334633231636363305f31302e3133312e3138322e3133385f323032313134363030323037"},want: `{"result":"success","message":"下线成功！"}`,wantErr: false},
		{name: testing.CoverMode(), args: args{userIndex: "sss"},want: `{"result":"fail","message":"用户已不在线"}`,wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LogoutCampusNetWork(tt.args.userIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("LogoutCampusNetWork() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LogoutCampusNetWork() = %v, want %v", got, tt.want)
			}
		})
	}
}
