package version

import (
	"reflect"
	"testing"
)

func TestVersion_String(t *testing.T) {
	tests := []struct {
		name string
		v    *Version
		want string
	}{
		{
			name: "test string",
			v: &Version{
				major:      1,
				minor:      2,
				patch:      3,
				identifier: "alpha2.0",
				metadata:   "20190715-a3dc5s",
			},
			want: "v1.2.3-alpha2.0+20190715-a3dc5s",
		},
		{
			name: "test string pretty 1",
			v: &Version{
				major:      1,
				minor:      2,
				patch:      3,
				identifier: "alpha2.0",
				metadata:   "20190715-a3dc5s",
				PrettyPrint: true,
			},
			want: "v1.2.3 (alpha2.0) (2019-07-15 a3dc5s)",
		},
		{
			name: "test string pretty 2",
			v: &Version{
				major:      1,
				minor:      2,
				patch:      3,
				identifier: "alpha2.0",
				metadata:   "a3dc5s-20190715",
				PrettyPrint: true,
			},
			want: "v1.2.3 (alpha2.0) (2019-07-15 a3dc5s)",
		},
		{
			name: "test string pretty 3",
			v: &Version{
				major:      1,
				minor:      2,
				patch:      3,
				identifier: "alpha2.0",
				metadata:   "a3dc5s",
				PrettyPrint: true,
			},
			want: "v1.2.3 (alpha2.0) (a3dc5s)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.String(); got != tt.want {
				t.Errorf("Version.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewVersion(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    *Version
		wantErr bool
	}{
		{
			name: "test constructor 1",
			args: args{
				v: "1.2.3-alpha2.0+20190715-a3dc5s",
			},
			want: &Version{
				major:      1,
				minor:      2,
				patch:      3,
				identifier: "alpha2.0",
				metadata:   "20190715-a3dc5s",
			},
		},
		{
			name: "test constructor 2",
			args: args{
				v: "1.2",
			},
			want: &Version{
				major: 1,
				minor: 2,
			},
		},
		{
			name: "test constructor 3",
			args: args{
				v: "a.2",
			},
			wantErr: true,
		},
		{
			name: "test constructor 4",
			args: args{
				"z5-1",
			},
			wantErr: true,
		},
		{
			name: "test constructor 5",
			args: args{
				"3_5_1",
			},
			wantErr: true,
		},
		{
			name: "test constructor 6",
			args: args{
				"3-5+1",
			},
			want: &Version{
				major:      3,
				identifier: "5",
				metadata:   "1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewVersion(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
