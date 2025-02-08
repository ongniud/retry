package retry

import (
	"reflect"
	"testing"
	"time"

	json "github.com/json-iterator/go"
)

type DelayArgs struct {
	Loop int
	BackOff BackOff
}

type BackOff struct {
	Delay time.Duration
	Max time.Duration
}

func Test_ExponentialDelay(t *testing.T) {
	tests := []struct {
		name string
		args DelayArgs
		want []time.Duration
	}{
		{
			name: "normal",
			args: DelayArgs{
				Loop: 10,
				BackOff: BackOff{
					Delay: 1 * time.Second,
					Max: 10 * time.Second,
				},
			},
			want: []time.Duration{
				1 * time.Second,
				2 * time.Second,
				4 * time.Second,
				8 * time.Second,
				10 * time.Second,
				10 * time.Second,
				10 * time.Second,
				10 * time.Second,
				10 * time.Second,
				10 * time.Second,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := ExponentialDelay(tt.args.BackOff.Delay, tt.args.BackOff.Max)
			var res []time.Duration
			for i:=1; i<=tt.args.Loop;i++ {
				res = append(res, fn(uint32(i)))
			}
			if !reflect.DeepEqual(res, tt.want) {
				gotData, _ := json.MarshalToString(res)
				wantData, _ := json.MarshalToString(tt.want)
				t.Errorf("Test case `%s` failed.\n Got: %s\n Want: %s\n", tt.name, gotData, wantData)
			}
		})
	}
}

func Test_ConstDelay(t *testing.T) {
	tests := []struct {
		name string
		args DelayArgs
		want []time.Duration
	}{
		{
			name: "negative",
			args: DelayArgs{
				Loop: 5,
				BackOff: BackOff{
					Delay: 5 * time.Second,
				},
			},
			want: []time.Duration{
				5 * time.Second,
				5 * time.Second,
				5 * time.Second,
				5 * time.Second,
				5 * time.Second,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := ConstDelay(tt.args.BackOff.Delay)
			var res []time.Duration
			for i:=1; i<=tt.args.Loop;i++ {
				res = append(res, fn(uint32(i)))
			}
			if !reflect.DeepEqual(res, tt.want) {
				gotData, _ := json.MarshalToString(res)
				wantData, _ := json.MarshalToString(tt.want)
				t.Errorf("Test case `%s` failed.\n Got: %s\n Want: %s\n", tt.name, gotData, wantData)
			}
		})
	}
}

func Test_NoDelay(t *testing.T) {
	tests := []struct {
		name string
		args DelayArgs
		want []time.Duration
	}{
		{
			name: "negative",
			args: DelayArgs{
				Loop: 5,
			},
			want: []time.Duration{
				0 * time.Second,
				0 * time.Second,
				0 * time.Second,
				0 * time.Second,
				0 * time.Second,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := NoDelay()
			var res []time.Duration
			for i:=1; i<=tt.args.Loop;i++ {
				res = append(res, fn(uint32(i)))
			}
			if !reflect.DeepEqual(res, tt.want) {
				gotData, _ := json.MarshalToString(res)
				wantData, _ := json.MarshalToString(tt.want)
				t.Errorf("Test case `%s` failed.\n Got: %s\n Want: %s\n", tt.name, gotData, wantData)
			}
		})
	}
}
