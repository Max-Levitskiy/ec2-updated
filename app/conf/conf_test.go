package conf

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)
func Test_shouldInitConfig(t *testing.T) {
	_ = Init()
	tests := []struct {
		key string
		want interface{}
	}{
		{key: "app.tag", want: "myApp"},
		{key: "app.region", want: "eu-central-1"},
		{key: "app.ec2.required-amount", want: 3},
		{key: "app.ec2.class", want: "t2.micro"},
	}
	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			assert.Equal(t, viper.Get(tt.key), tt.want)
		})
	}
}

