import { readFileSync } from "fs"
import { loadPolicy } from "@open-policy-agent/opa-wasm";

const policyWasm = readFileSync("../artifact/policy.wasm"); // fetch("/path/to/wasm").then(res => res.arrayBuffer())
const policy = await loadPolicy(policyWasm);

console.log("[validation/email]");
[
    "hogehoge@example.com",
    "hogehogeexample.com",
    "hoge@hoge@example.com",
    "hogehoge@examplecom"
].forEach((email) => {
    console.log(
        email,
        policy.evaluate(email, "validation/email")[0].result
    );
});

console.log()
console.log("[validation/domain]");
[
    "example.com",
    "examplecom",
    ".example.com"
].forEach((domain) => {
    console.log(
        domain,
        policy.evaluate(domain, "validation/domain")[0].result
    );
});
