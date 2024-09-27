package script_creator

import "testing"

func TestScriptCreator_GetScript(t *testing.T) {
	type args struct {
		data [][]int
	}

	ar := args{
		data: [][]int{
			{8, 7, 1, 9, 4, 6, 2, 5, 3},
			{3, 6, 4, 1, 2, 5, 9, 8, 7},
			{9, 2, 5, 3, 7, 8, 1, 6, 4},
			{2, 4, 8, 7, 1, 3, 6, 9, 5},
			{7, 5, 9, 8, 6, 2, 3, 4, 1},
			{6, 1, 3, 4, 5, 9, 7, 2, 8},
			{5, 9, 7, 6, 3, 4, 8, 1, 2},
			{1, 8, 2, 5, 9, 7, 4, 3, 6},
			{4, 3, 6, 2, 8, 1, 5, 7, 9},
		},
	}

	expectedScript := "input tap 90 505; input tap 889 2011; input tap 202 505; input tap 773 2011; input tap 314 505; input tap 77 2011; input tap 426 505; input tap 1005 2011; input tap 538 505; input tap 425 2011; input tap 650 505; input tap 657 2011; input tap 762 505; input tap 193 2011; input tap 874 505; input tap 541 2011; input tap 986 505; input tap 309 2011; input tap 130 619; input tap 309 2011; input tap 242 619; input tap 657 2011; input tap 354 619; input tap 425 2011; input tap 466 619; input tap 77 2011; input tap 578 619; input tap 193 2011; input tap 690 619; input tap 541 2011; input tap 802 619; input tap 1005 2011; input tap 914 619; input tap 889 2011; input tap 1026 619; input tap 773 2011; input tap 130 733; input tap 1005 2011; input tap 242 733; input tap 193 2011; input tap 354 733; input tap 541 2011; input tap 466 733; input tap 309 2011; input tap 578 733; input tap 773 2011; input tap 690 733; input tap 889 2011; input tap 802 733; input tap 77 2011; input tap 914 733; input tap 657 2011; input tap 1026 733; input tap 425 2011; input tap 130 847; input tap 193 2011; input tap 242 847; input tap 425 2011; input tap 354 847; input tap 889 2011; input tap 466 847; input tap 773 2011; input tap 578 847; input tap 77 2011; input tap 690 847; input tap 309 2011; input tap 802 847; input tap 657 2011; input tap 914 847; input tap 1005 2011; input tap 1026 847; input tap 541 2011; input tap 130 961; input tap 773 2011; input tap 242 961; input tap 541 2011; input tap 354 961; input tap 1005 2011; input tap 466 961; input tap 889 2011; input tap 578 961; input tap 657 2011; input tap 690 961; input tap 193 2011; input tap 802 961; input tap 309 2011; input tap 914 961; input tap 425 2011; input tap 1026 961; input tap 77 2011; input tap 130 1075; input tap 657 2011; input tap 242 1075; input tap 77 2011; input tap 354 1075; input tap 309 2011; input tap 466 1075; input tap 425 2011; input tap 578 1075; input tap 541 2011; input tap 690 1075; input tap 1005 2011; input tap 802 1075; input tap 773 2011; input tap 914 1075; input tap 193 2011; input tap 1026 1075; input tap 889 2011; input tap 130 1189; input tap 541 2011; input tap 242 1189; input tap 1005 2011; input tap 354 1189; input tap 773 2011; input tap 466 1189; input tap 657 2011; input tap 578 1189; input tap 309 2011; input tap 690 1189; input tap 425 2011; input tap 802 1189; input tap 889 2011; input tap 914 1189; input tap 77 2011; input tap 1026 1189; input tap 193 2011; input tap 130 1303; input tap 77 2011; input tap 242 1303; input tap 889 2011; input tap 354 1303; input tap 193 2011; input tap 466 1303; input tap 541 2011; input tap 578 1303; input tap 1005 2011; input tap 690 1303; input tap 773 2011; input tap 802 1303; input tap 425 2011; input tap 914 1303; input tap 309 2011; input tap 1026 1303; input tap 657 2011; input tap 130 1417; input tap 425 2011; input tap 242 1417; input tap 309 2011; input tap 354 1417; input tap 657 2011; input tap 466 1417; input tap 193 2011; input tap 578 1417; input tap 889 2011; input tap 690 1417; input tap 77 2011; input tap 802 1417; input tap 541 2011; input tap 914 1417; input tap 773 2011; input tap 1026 1417; input tap 1005 2011; "

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test_1",
			args: ar,
			want: expectedScript,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			creator := &ScriptCreator{}
			if got := creator.GetScript(tt.args.data); got != tt.want {
				t.Errorf("GetScript() = %v, want %v", got, tt.want)
			}
		})
	}
}
