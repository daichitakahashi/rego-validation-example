package policy

import (
	"context"
	"embed"
	"fmt"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/loader"
	"github.com/open-policy-agent/opa/rego"
)

var (
	//go:embed bundle.tar.gz
	bundle embed.FS

	emailPolicy,
	domainPolicy rego.PreparedEvalQuery

	querySet = map[string]*rego.PreparedEvalQuery{
		"data.validation.email.valid":  &emailPolicy,
		"data.validation.domain.valid": &domainPolicy,
	}
)

func init() {
	b, err := loader.NewFileLoader().WithFS(bundle).AsBundle("bundle.tar.gz")
	if err != nil {
		panic(fmt.Sprintf("policy: failed to load bundle: %s", err))
	}

	ctx := context.Background()
	parsedBundle := rego.ParsedBundle("bundle.tar.gz", b)

	for query, ref := range querySet {
		*ref, err = rego.New(
			parsedBundle,
			rego.ParsedQuery(ast.MustParseBody(query)),
		).PrepareForEval(ctx)
		if err != nil {
			panic(fmt.Sprintf("failed to prepare query(%s): %s", query, err))
		}
	}
}

func eval(ctx context.Context, policy rego.PreparedEvalQuery, input any) (bool, error) {
	result, err := policy.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		return false, err
	}
	return result.Allowed(), nil
}

// ValidEmail validates email address.
func ValidEmail(ctx context.Context, email string) (bool, error) {
	return eval(ctx, emailPolicy, email)
}

// ValidDomain validates domain.
func ValidDomain(ctx context.Context, domain string) (bool, error) {
	return eval(ctx, domainPolicy, domain)
}
