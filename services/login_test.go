package services

import (
	"testing"
)

func TestFastLogin(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FastLogin()
		})
	}
}

// 这个单元测试依赖本地网络环境， 没办法测试, 是我水平不够， 水平够了也是能测试的
func TestFirstLogin(t *testing.T) {
	type args struct {
		userId   string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "用户殷梁禹", args: args{userId: "202114600207", password: "qwefsdd435fgsdH@H@"}, want: "ok", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if userOnline() {
				LogoutCampusNetWork()
			}
			got, err := FirstLogin(tt.args.userId, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("FirstLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FirstLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_formatCampusNetService(t *testing.T) {
	type args struct {
		service string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := formatCampusNetService(tt.args.service)
			if (err != nil) != tt.wantErr {
				t.Errorf("formatCampusNetService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("formatCampusNetService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUserIndex(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUserIndex(); got != tt.want {
				t.Errorf("GetUserIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogoutCampusNetWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogoutCampusNetWork()
		})
	}
}

func Test_userOnline(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{name: "s", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := userOnline(); got != tt.want {
				t.Errorf("userOnline() = %v, want %v", got, tt.want)
			}
		})
	}
}
