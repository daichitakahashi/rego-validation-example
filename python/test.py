#import os, sys
#sys.path.append(os.path.join(os.path.dirname(__file__), "deps"))

from opa_wasm import OPAPolicy
policy = OPAPolicy("../artifact/policy.wasm")

print("[validation/email]")
emails = [
    "hogehoge@example.com",
    "hogehogeexample.com",
    "hoge@hoge@example.com",
    "hogehoge@examplecom"
]
for email in emails:
    print(
        email,
        policy.evaluate(email, "validation/email")[0]["result"]
    )

print()
print("validation/domain")
domains = [
    "example.com",
    "examplecom",
    ".example.com"
]
for domain in domains:
    print(
        domain,
        policy.evaluate(domain, "validation/domain")[0]["result"]
    )
