package main

type remark string

func (r *remark) getRunes() []rune {
	runeSlice := make([]rune, len(*(*string)(r)))
	for _, r := range *(*string)(r) {
		runeSlice = append(runeSlice, r)
	}
	return runeSlice
}

func (r *remark) isAllCaps() bool {
	runes := r.getRunes()
	for _, r := range runes {
		if isLetter(r) && !isUpperCase(r) {
			return false
		}
	}
	return true
}

func isUpperCase(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isLetter(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
}

func isWhiteSpace(r rune) bool {
	return r == '\t' || r == '\n' || r == '\r'
}

func main() {
	//r := remark("JoeBoy")
}
