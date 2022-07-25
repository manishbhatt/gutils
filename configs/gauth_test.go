package configs

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"testing"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/tasks/v1"
)

func Test_getClient(t *testing.T) {
	home, _ := os.UserHomeDir()
	b, _ := ioutil.ReadFile(fmt.Sprintf("%s/.got/credentials.json", home))
	// If modifying these scopes, delete your previously saved token.json.
	config, _ := google.ConfigFromJSON(b, tasks.TasksReadonlyScope)

	type args struct {
		config *oauth2.Config
	}

	tests := []struct {
		name    string
		args    args
		want    *http.Client
		wantErr bool
	}{
		{
			name:    "basic",
			args:    args{config: config},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetClient(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("getClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
