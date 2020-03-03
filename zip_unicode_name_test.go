package zip_unicode_name

import "testing"

func TestParse(t *testing.T) {
	extra := []byte{10, 0, 32, 0, 0, 0, 0, 0, 1, 0, 24, 0, 151, 246, 70, 6, 82, 93, 211, 1, 134, 233, 233, 145, 226, 121, 211, 1, 134, 233, 233, 145, 226, 121, 211, 1, 117, 112, 39, 0, 1, 98, 3, 52, 221, 232, 167, 129, 232, 166, 139, 232, 180, 165, 230, 149, 151, 228, 188, 160, 229, 130, 179, 229, 133, 173, 228, 185, 166, 229, 133, 173, 230, 155, 184, 46, 120, 108, 102}
	if parseUnicodeFileName(extra, "jhtvcptytw.xlf") != "见見败敗传傳六书六書.xlf" {
		t.Fail()
	}
}

func TestWrongCRC32(t *testing.T) {
	extra := []byte{10, 0, 32, 0, 0, 0, 0, 0, 1, 0, 24, 0, 151, 246, 70, 6, 82, 93, 211, 1, 134, 233, 233, 145, 226, 121, 211, 1, 134, 233, 233, 145, 226, 121, 211, 1, 117, 112, 39, 0, 1, 98, 3, 52, 221, 232, 167, 129, 232, 166, 139, 232, 180, 165, 230, 149, 151, 228, 188, 160, 229, 130, 179, 229, 133, 173, 228, 185, 166, 229, 133, 173, 230, 155, 184, 46, 120, 108, 102}

	if parseUnicodeFileName(extra, "") != "" {
		t.Fail()
	}
	if parseUnicodeFileName(extra, "somename") != "" {
		t.Fail()
	}
}

func TestManyUpSequences(t *testing.T) {
	extra := []byte{117, 112, 10, 0, 32, 0, 0, 0, 117, 112, 0, 0, 1, 0, 24, 0, 151, 246, 117, 112, 70, 6, 82, 93, 211, 1, 134, 117, 112, 233, 233, 145, 226, 121, 211, 1, 134, 233, 233, 145, 226, 121, 211, 1, 117, 112, 39, 0, 1, 98, 3, 52, 221, 232, 167, 129, 232, 166, 139, 232, 180, 165, 230, 149, 151, 228, 188, 160, 229, 130, 179, 229, 133, 173, 228, 185, 166, 229, 133, 173, 230, 155, 184, 46, 120, 108, 102}
	if parseUnicodeFileName(extra, "jhtvcptytw.xlf") != "见見败敗传傳六书六書.xlf" {
		t.Fail()
	}
}

func TestReadUInt8(t *testing.T) {
	bytes := []byte{117}
	v, _, err := read8bits(bytes, 0)
	if v != 117 {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}
}

func TestReadUInt8Fail(t *testing.T) {
	bytes := []byte{}
	_, _, err := read8bits(bytes, 0)
	if err == nil {
		t.Fail()
	}
}

func TestReadUInt16(t *testing.T) {
	bytes := []byte{117, 112}
	v, _, err := read16bits(bytes, 0)
	if v != 0x7075 {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}
}

func TestReadUIn16Fail(t *testing.T) {
	bytes := []byte{117}
	_, _, err := read16bits(bytes, 0)
	if err == nil {
		t.Fail()
	}
}
