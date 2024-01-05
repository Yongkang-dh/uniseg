package uniseg

import "testing"

func TestEmojiCount(t *testing.T) {

	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{sentence: "this is a test comment with emoji 👨‍❤️‍👨, 👨‍👨‍👦, 👨🏻‍❤️‍👨🏻, 🇫🇲, 🇬🇲"},
			want: 5,
		},
		{
			name: "",
			args: args{sentence: "this is a test without emoji"},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EmojiCount(tt.args.sentence); got != tt.want {
				t.Errorf("emoji nums = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContinuousEmojiCount(t *testing.T) {

	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "split emojis",
			args: args{sentence: "this is a test comment with split emoji 👨‍❤️‍👨,👨‍👨‍👦, 👨🏻‍❤️‍👨🏻, 🇫🇲, 🇬🇲"},
			want: 1,
		},
		{
			name: "continuous emojis without blank",
			args: args{sentence: "👨‍❤️‍👨👨‍👨‍👦👨🏻‍❤️‍👨🏻🇫🇲🇬🇲"},
			want: 5,
		},
		{
			name: "continuous emojis with blank",
			args: args{sentence: "👨‍❤️‍👨 👨‍👨‍👦 👨🏻‍❤️‍👨🏻 🇫🇲 🇬🇲"},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContinuousEmojiCount(tt.args.sentence); got != tt.want {
				t.Errorf("emoji nums = %v, want %v", got, tt.want)
			}
		})
	}
}
