package validation.domain

default valid := false

valid {
    253 >= count(input)
    regex.match(`^((?:[a-z\d]\-?)+)((?:\.(?:(?:(?:\-?[a-z\d]+)+\-?)|\-))*)(\.(?:\-?[a-z\d])+)$`, input)
}
