package validation.email

default valid := false

valid {
	254 >= count(input)
    [local, domain] := split(input, "@")
    64 >= count(local)
    regex.match(`^(?:[a-z]|\d)+(?:[-_.](?:[a-z]|\d)+)*$`, local)
    253 >= count(domain)
    regex.match(`^((?:[a-z\d]\-?)+)((?:\.(?:(?:(?:\-?[a-z\d]+)+\-?)|\-))*)(\.(?:\-?[a-z\d])+)$`, domain)
}
