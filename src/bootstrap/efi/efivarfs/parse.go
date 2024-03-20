package efivarfs

import (
	"encoding"
	"fmt"
	"io"
	"os"
)

const (
	efivars = "/sys/firmware/efi/efivars"
)

func ParseVar[T encoding.BinaryUnmarshaler](name string, guid string) (*T, error) {
	f, err := os.Open(fmt.Sprintf("%s/%s-%s", efivars, name, guid))
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var t T
	if err = t.UnmarshalBinary(data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return &t, nil
}
