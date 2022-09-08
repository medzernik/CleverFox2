package Info

import (
	"reflect"
	"testing"
)

func TestUserID_ToUserMention1(t *testing.T) {
	tests := []struct {
		name string
		self UserID
		want *UserID
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.self.ToUserMention(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUserMention() = %v, want %v", got, tt.want)
			}
		})
	}
}
