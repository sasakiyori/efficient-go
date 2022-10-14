package string_test

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

type StringConcatenationTestData struct {
	ByteValue      byte
	RuneValue      rune
	ByteArrayValue []byte
	StringValue    string
}

var (
	td *StringConcatenationTestData
)

// TODO: import different datas from files
func initTestData() {
	td = &StringConcatenationTestData{
		ByteValue:      'b',
		RuneValue:      'r',
		ByteArrayValue: []byte("ByteArrayValue"),
		StringValue:    "StringValue",
	}
}

func TestMain(m *testing.M) {
	initTestData()
	code := m.Run()
	os.Exit(code)
}

func BenchmarkStringBuilderInCommonUsage(b *testing.B) {
	var result string
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		_, _ = sb.WriteString(td.StringValue)
		_, _ = sb.WriteRune(td.RuneValue)
		_, _ = sb.Write(td.ByteArrayValue)
		_ = sb.WriteByte(td.ByteValue)
		result = sb.String()
	}
	_ = result
}

func BenchmarkStringBuilderInFmtf(b *testing.B) {
	var result string
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		_, _ = fmt.Fprintf(&sb, "%s%c%s%c", td.StringValue, td.RuneValue, td.ByteArrayValue, td.ByteValue)
		result = sb.String()
	}
	_ = result
}

func BenchmarkStringBuilderInFmt(b *testing.B) {
	var result string
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		_, _ = fmt.Fprint(&sb, td.StringValue, td.RuneValue, td.ByteArrayValue, td.ByteValue)
		result = sb.String()
	}
	_ = result
}

func BenchmarkBytesBuffer(b *testing.B) {
	var result string
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		_, _ = buf.WriteString(td.StringValue)
		_, _ = buf.WriteRune(td.RuneValue)
		_, _ = buf.Write(td.ByteArrayValue)
		_ = buf.WriteByte(td.ByteValue)
		result = buf.String()
	}
	_ = result
}
