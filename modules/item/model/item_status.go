package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type ItemStatus int

const (
	ItemStatusDoing ItemStatus = iota
	ItemStatusDone
	ItemStatusDeleted
)

var allItemStatus = [3]string{"DOING", "DONE", "DELETED"}

func (item *ItemStatus) String() string {
	return allItemStatus[*item]
}
func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}

	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}
func (item *ItemStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")

	itemValue, err := parseStringToItemStatus(str)
	if err != nil {
		return err
	}

	*item = itemValue
	return nil
}

func parseStringToItemStatus(s string) (ItemStatus, error) {
	for i := range allItemStatus {
		if allItemStatus[i] == s {
			return ItemStatus(i), nil
		}
	}

	return ItemStatus(0), errors.New("invalid status string")
}

func (item *ItemStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return fmt.Errorf("fail to scan data from sql: %s", value)
	}

	v, err := parseStringToItemStatus(string(bytes))
	if err != nil {
		return fmt.Errorf("fail to scan data from sql: %s", value)
	}

	*item = v
	return nil
}
func (item *ItemStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}

	return item.String(), nil
}
