package go_separate

import "testing"

func TestOneThousandCanBeFormatted(t *testing.T) {
	number := 1000.0
	expected := "1,000"
	actual := FormatLargeNumber(number)

	if expected != actual {
		t.Fatalf("Expected %f but got %f", expected, actual)
	}
}

func TestOneHundredThousandCanBeFormatted(t *testing.T) {
	number := 100000.0
	expected := "100,000"
	actual := FormatLargeNumber(number)

	if expected != actual {
		t.Fatalf("Expected %f but got %f", expected, actual)
	}
}

func TestTenThousandCanBeFormatted(t *testing.T) {
	number := 10000.0
	expected := "10,000"
	actual := FormatLargeNumber(number)

	if expected != actual {
		t.Fatalf("Expected %f but got %f", expected, actual)
	}
}

func TestOneMillionCanBeFormatted(t *testing.T) {
	number := 1000000.0
	expected := "1,000,000"
	actual := FormatLargeNumber(number)

	if expected != actual {
		t.Fatalf("Expected %f but got %f", expected, actual)
	}
}

func TestTenMillionCanBeFormatted(t *testing.T) {
	number := 10000000.0
	expected := "10,000,000"
	actual := FormatLargeNumber(number)

	if expected != actual {
		t.Fatalf("Expected %f but got %f", expected, actual)
	}
}

func TestOneHundredMillionCanBeFormatted(t *testing.T) {
	number := 100000000.0
	expected := "100,000,000"
	actual := FormatLargeNumber(number)

	if expected != actual {
		t.Fatalf("Expected %f but got %f", expected, actual)
	}
}

func TestAveryLargeNumberCanBeFormatted(t *testing.T) {
	number := 100000000000.0
	expected := "100,000,000,000"
	actual := FormatLargeNumber(number)

	if expected != actual {
		t.Fatalf("Expected %f but got %f", expected, actual)
	}
}