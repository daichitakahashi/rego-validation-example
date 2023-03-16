package policy

import (
	"context"
	"testing"
)

var ctx = context.Background()

func TestValidEmail(t *testing.T) {
	type testCase struct {
		email string
		valid bool
	}
	runTest := func(t *testing.T, tc testCase) {
		t.Helper()
		valid, err := ValidEmail(ctx, tc.email)
		if err != nil {
			t.Fatal(err)
		}
		if valid != tc.valid {
			t.Fatalf("unexpected validity: want=%t got=%t (email=%s)", tc.valid, valid, tc.email)
		}
	}

	testCases := []testCase{
		{
			email: "hoge@example.com",
			valid: true,
		}, {
			email: "hoge@fuga@example.com",
			valid: false,
		}, {
			email: "hoge@example",
			valid: false,
		},
		// ...
	}

	for _, tc := range testCases {
		t.Run(tc.email, func(t *testing.T) {
			runTest(t, tc)
		})
	}
}

func TestValidDomain(t *testing.T) {
	type testCase struct {
		domain string
		valid  bool
	}
	runTest := func(t *testing.T, tc testCase) {
		t.Helper()
		valid, err := ValidDomain(ctx, tc.domain)
		if err != nil {
			t.Fatal(err)
		}
		if valid != tc.valid {
			t.Fatalf("unexpected validity: want=%t got=%t (domain=%s)", tc.valid, valid, tc.domain)
		}
	}

	testCases := []testCase{
		{
			domain: "example.com",
			valid:  true,
		}, {
			domain: "example",
			valid:  false,
		},
		// ...
	}

	for _, tc := range testCases {
		t.Run(tc.domain, func(t *testing.T) {
			runTest(t, tc)
		})
	}
}
