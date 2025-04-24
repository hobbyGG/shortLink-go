package mysqlse

import "testing"

func TestClient_Get(t *testing.T) {
	tests := []struct {
		name    string
		c       Client
		want    uint64
		wantErr bool
	}{
		{name: "正常示例", c: *NewClient("root:123@tcp(127.0.0.1:13306)/slDB?charset=utf8mb4&parseTime=True&loc=Local"), want: 13, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
