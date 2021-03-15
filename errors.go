package readenv

type errKeyNotInEnvData string

type errFindingTwoRegexSubmatches string

func (e errKeyNotInEnvData) Error() string {
	return "key " + string(e) + " not found in EnvData"
}

func (e errFindingTwoRegexSubmatches) Error() string {
	return "error finding two submatches in line " + string(e)
}
