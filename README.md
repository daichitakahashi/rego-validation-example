# rego-validation-example

Examples for validation using [Policy Language Rego](https://www.openpolicyagent.org/docs/latest/policy-language/).

## Policies
* `policy/email.rego`: validate email address
  * `policy/email_test.rego`: test code
* `policy/domain.rego`: validate domain part of email address
  * `policy/domain_test.rego`: test code
* `artifact/policy.wasm`: bundled WebAssembly(`make build`)

## Examples
### Go
Go example handles ".rego" files directly using [github.com/open-policy-agent/opa](https://github.com/open-policy-agent/opa).
```shell
$ cd go
$ go test -v .
=== RUN   TestValidEmail
=== RUN   TestValidEmail/hoge@example.com
=== RUN   TestValidEmail/hoge@fuga@example.com
=== RUN   TestValidEmail/hoge@example
--- PASS: TestValidEmail (0.00s)
    --- PASS: TestValidEmail/hoge@example.com (0.00s)
    --- PASS: TestValidEmail/hoge@fuga@example.com (0.00s)
    --- PASS: TestValidEmail/hoge@example (0.00s)
=== RUN   TestValidDomain
=== RUN   TestValidDomain/example.com
=== RUN   TestValidDomain/example
--- PASS: TestValidDomain (0.00s)
    --- PASS: TestValidDomain/example.com (0.00s)
    --- PASS: TestValidDomain/example (0.00s)
PASS
ok      rego-validation-example/go      0.365s
```

### JavaScript(node.js)
Javascript example uses "policy.wasm" with [@open-policy-agent/opa-wasm](https://github.com/open-policy-agent/npm-opa-wasm).
```shell
$ cd node
$ npm install
$ node test.mjs
[validation/email]
hogehoge@example.com { valid: true }
hogehogeexample.com { valid: false }
hoge@hoge@example.com { valid: false }
hogehoge@examplecom { valid: false }

[validation/domain]
example.com { valid: true }
examplecom { valid: false }
.example.com { valid: false }
```

### Python(3.10)
Python example uses "policy.wasm" too, with [opa-wasm](https://pypi.org/project/opa-wasm/).
```shell
$ cd python
$ pip3 install -r requirements.txt
$ python3 test.py
[validation/email]
hogehoge@example.com {'valid': True}
hogehogeexample.com {'valid': False}
hoge@hoge@example.com {'valid': False}
hogehoge@examplecom {'valid': False}

[validation/domain]
example.com {'valid': True}
examplecom {'valid': False}
.example.com {'valid': False}
```
