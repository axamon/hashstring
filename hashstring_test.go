// Copyright 2019 Alberto Bregliano. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hashstring

import (
	"testing"
)

func TestMd5Sum(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "md5sum 1", args: args{s: "pippo"}, want: "f583ca884c1d93458fb61ed137ff44f6"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5Sum(tt.args.s); got != tt.want {
				t.Errorf("Md5Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSha256Sum(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "sha256sum 1", args: args{s: "pippo"}, want: "44aa336af4cb14a879432e53dd6571c7fa9bccafb75f488259262d6ea3a4d91b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sha256Sum(tt.args.s); got != tt.want {
				t.Errorf("Sha256Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSha512Sum(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "Sha512Sum 1", args: args{s: "pippo"}, want: "b7811c6592757524f0930c6517f612453f7ce22d09f0a6badbeca6610880ab71d118d2e9fa720aaa637756d853880c66d5043397a680ef1375e84ec07cd81beb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sha512Sum(tt.args.s); got != tt.want {
				t.Errorf("Sha512Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
