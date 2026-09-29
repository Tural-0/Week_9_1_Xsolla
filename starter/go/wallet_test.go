package wallet

import "testing"

func TestCurrencySymbolKnown(t *testing.T) {
	if CurrencySymbol("USD") != "$" {
		t.Fatal("want $")
	}
}

func TestCurrencySymbolUnknown(t *testing.T) {
	if CurrencySymbol("XYZ") != "XYZ" {
		t.Fatal("want XYZ")
	}
}

func TestFormatAmount(t *testing.T) {
	if FormatAmount(1234) != "12.34" {
		t.Fatal("want 12.34")
	}
	if FormatAmount(5) != "0.05" {
		t.Fatal("want 0.05")
	}
}

func TestParseAmountBasic(t *testing.T) {
	if n, _ := ParseAmount("12.34"); n != 1234 {
		t.Fatal("want 1234")
	}
	if n, _ := ParseAmount("12"); n != 1200 {
		t.Fatal("want 1200")
	}
}
