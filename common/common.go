package common

import uuid "github.com/satori/go.uuid"

func GetUUid() string {
	id := uuid.NewV4()
	ids := id.String()
	return ids
}
