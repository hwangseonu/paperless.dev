package mail

import "bytes"

func VerifyCodeTemplate(code string) (string, error) {
	var buf bytes.Buffer
	if err := verifyCodeTmpl.Execute(&buf, code); err != nil {
		return "", err
	}
	return buf.String(), nil
}
