package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {

	if word != "" {
		var res string

		for _, r := range word {
			res = string(r) + res
		}
		return res
	}
	return ""
}