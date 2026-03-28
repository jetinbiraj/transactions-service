package accounts

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func validateDocumentNumber(doc string) error {
	doc = strings.TrimSpace(doc)
	if doc == "" {
		return errors.New("document_number is required")
	}

	// TODO: Assumption made that document_number will be of length 11, change the logic if that's not true
	if len(doc) != 11 {
		return errors.New("document_number must be 11 digits")
	}

	// TODO: Assumption made that document_number will only contain digits, change the logic if that's not true
	matched, _ := regexp.MatchString(`^\d+$`, doc)
	if !matched {
		return errors.New("document_number must contain only digits")
	}

	return nil
}

func validateAccountId(accountId string) error {
	if accountId == "" {
		return errors.New("account_id is required")
	}

	id, err := strconv.ParseInt(accountId, 10, 64)
	if err != nil || id <= 0 {
		return errors.New("invalid account_id")
	}

	return nil
}
