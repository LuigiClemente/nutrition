package utils

import "strconv"

func ParseAndValidateID(idStr string) (int, error) {
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
