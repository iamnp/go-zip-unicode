package zip_unicode_name

import (
	"archive/zip"
	"encoding/binary"
	"errors"
	"hash/crc32"
)

const (
	_UNICODE_PATH_EXTRA_FIELD uint16 = 0x7075
)

func read8bits(extra []byte, pos int) (uint8, int, error) {
	if pos >= len(extra) {
		return 0, 0, errors.New("byte array is too short")
	}
	return extra[pos], pos + 1, nil
}

func read16bits(extra []byte, pos int) (uint16, int, error) {
	if pos+2 > len(extra) {
		return 0, 0, errors.New("byte array is too short")
	}
	return binary.LittleEndian.Uint16(extra[pos : pos+2]), pos + 2, nil
}

func read32bits(extra []byte, pos int) (uint32, int, error) {
	if pos+4 > len(extra) {
		return 0, 0, errors.New("byte array is too short")
	}
	return binary.LittleEndian.Uint32(extra[pos : pos+4]), pos + 4, nil
}

func ParseUnicodeFileName(f *zip.File) string {
	return parseUnicodeFileName(f.Extra, f.Name)
}

// https://pkware.cachefly.net/webdocs/casestudies/APPNOTE.TXT
// 4.6.9 -Info-ZIP Unicode Path Extra Field (0x7075):
func parseUnicodeFileName(extra []byte, name string) string {
	i := 0
	for i < len(extra) {
		var flag uint16
		var err error
		flag, i, err = read16bits(extra, i)
		if err != nil {
			break
		}

		if flag == _UNICODE_PATH_EXTRA_FIELD {

			var tsize uint16
			tsize, i, err = read16bits(extra, i)
			if err != nil {
				continue
			}

			var version uint8
			version, i, err = read8bits(extra, i)
			if err != nil {
				continue
			}
			if version != 1 {
				continue
			}

			var nameCRC32 uint32
			nameCRC32, i, err = read32bits(extra, i)
			if err != nil {
				continue
			}

			unicodeBytesCount := int(tsize) - 1 - 4
			if unicodeBytesCount < 0 {
				continue
			}
			if i+unicodeBytesCount > len(extra) {
				continue
			}

			unicodeBytes := extra[i : i+unicodeBytesCount]
			if crc32.ChecksumIEEE([]byte(name)) != nameCRC32 {
				continue
			}

			return string(unicodeBytes)
		} else {
			i -= 1
		}
	}

	return ""
}
