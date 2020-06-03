package logger2

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// MaskCard masks card number and cvv if exists
func MaskCard(s string) string {
	m := map[string]interface{}{}

	// Check if string is JSON and mask card
	if err := json.Unmarshal([]byte(s), &m); err == nil {
		for k, v := range m {
			if value, ok := v.(map[string]interface{}); ok {
				var (
					b, _ = json.Marshal(value)
					m2   = map[string]interface{}{}
				)

				_ = json.Unmarshal([]byte(MaskCard(string(b))), &m2)
				m[k] = m2
				continue
			}

			switch strings.ToLower(k) {
			case "cvv",
				"credit_card_cvv",
				"cardcvc",
				"securitycode":
				if value, ok := v.(string); ok {
					m[k] = mask(value, len(value))
				}
			case "number",
				"cardnumber",
				"cardnum",
				"cardno",
				"accountnumber",
				"card_number",
				"card_no":
				if value, ok := v.(string); ok {
					// m[k] = mask(value, 4)
					m[k] = maskFirstMLastN(value, 6, 4)
				} else if value, ok := v.(float64); ok {
					// m[k] = mask(strconv.Itoa(int(value)), 4)
					m[k] = maskFirstMLastN(strconv.Itoa(int(value)), 6, 4)
				}
			case "email":
				if value, ok := v.(string); ok {
					m[k] = maskEmail(value)
				}
			}

		}

		b, _ := json.Marshal(m)

		return string(b)
	}

	// Check if string is URL encoded and does not contain `<`, and mask card
	if values, err := url.ParseQuery(s); !strings.Contains(s, "<") && strings.Contains(s, "=") && err == nil {
		newValues := url.Values{}
		for k, v := range values {
			switch strings.ToLower(k) {
			case "cvv", "credit_card_cvv", "cardcvc", "securitycode":
				newValues[k] = []string{mask(v[0], len(v[0]))}
			case "number", "cardnumber", "cardnum", "cardno", "accountnumber", "card_no":
				// newValues[k] = []string{mask(v[0], 4)}
				newValues[k] = []string{maskFirstMLastN(v[0], 6, 4)}
			case "email":
				newValues[k] = []string{maskEmail(v[0])}
			default:
				newValues[k] = v
			}
		}

		var (
			buf  strings.Builder
			keys = make([]string, 0, len(newValues))
		)
		for k := range newValues {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			vs := newValues[k]
			for _, v := range vs {
				if buf.Len() > 0 {
					buf.WriteByte('&')
				}
				buf.WriteString(k)
				buf.WriteByte('=')
				buf.WriteString(v)
			}
		}

		return buf.String()
	}

	// Check if string is XML and mask card
	r := regexp.MustCompile(`(?i)<(number|cardnumber|cardnum|cardno|accountnumber)>(\d+)<\/(number|cardnumber|cardnum|cardno|accountnumber)>`)
	if m := r.FindStringSubmatch(s); len(m) == 4 {
		s = r.ReplaceAllString(s, fmt.Sprintf("<%s>%s</%s>", m[1], maskFirstMLastN(m[2], 6, 4), m[3]))
	}

	r = regexp.MustCompile(`(?i)<(cvv|securitycode|cvNumber)>(\d{3,4})<\/(cvv|securitycode|cvNumber)>`)
	if m := r.FindStringSubmatch(s); len(m) == 4 {
		s = r.ReplaceAllString(s, fmt.Sprintf("<%s>%s</%s>", m[1], mask(m[2], len(m[2])), m[3]))
	}

	return s
}

func mask(str string, size int) (response string) {
	response = str

	if len(str) == size && size >= 0 {
		response = strings.Repeat("*", size)
		return
	}

	intSize := len(str) - size
	if intSize >= 0 {
		response = fmt.Sprintf("%s%s", strings.Repeat("*", intSize), str[intSize:])
		return
	}

	return
}

// maskFirstMLastN ...
func maskFirstMLastN(str string, M, N int) (response string) {
	if len(str) < M {
		M = len(str)
	}

	if M+N > len(str) {
		N = 0
	}
	firstM := str[:M]
	lastN := str[len(str)-N:]
	nib := len(str) - M - N
	valInB := strings.Repeat("*", nib)

	return fmt.Sprintf("%s%s%s", firstM, valInB, lastN)
}

// masks email
// hello@gmail.com -> hel**@gmail.com
func maskEmail(email string) (response string) {
	var (
		user   string
		domain string
	)
	split := strings.Split(email, "@")
	user = split[0]
	domain = split[1]

	if lengthUser := len(user); lengthUser >= 0 {
		repeater := lengthUser - 3
		user = fmt.Sprintf("%s%s", user[0:3], strings.Repeat("*", repeater))
	}
	return fmt.Sprintf("%s@%s", user, domain)
}
