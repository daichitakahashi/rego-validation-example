package validation.email

import future.keywords

test_valid_email if {
    valid with input as "hoge@example.com"
}

test_invalid_email_with_too_many_at if {
    not valid with input as "hoge@fuga@example.com"
}

test_invalid_email_with_invalid_domain if {
    not valid with input as "hoge@example"
}
