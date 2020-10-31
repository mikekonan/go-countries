package country

import (
	"strings"
	"testing"
)

func TestAllSet(t *testing.T) {
	for key, country := range countryByName {
		if !key.IsSet() {
			t.FailNow()
		}

		if !country.name.IsSet() {
			t.FailNow()
		}

		if !country.alpha2.IsSet() {
			t.FailNow()
		}

		if !country.alpha3.IsSet() {
			t.FailNow()
		}
	}

	for key, country := range countryByAlpha3 {
		if !key.IsSet() {
			t.FailNow()
		}

		if !country.name.IsSet() {
			t.FailNow()
		}

		if !country.alpha2.IsSet() {
			t.FailNow()
		}

		if !country.alpha3.IsSet() {
			t.FailNow()
		}
	}

	for key, country := range countryByAlpha2 {
		if !key.IsSet() {
			t.FailNow()
		}

		if !country.name.IsSet() {
			t.FailNow()
		}

		if !country.alpha2.IsSet() {
			t.FailNow()
		}

		if !country.alpha3.IsSet() {
			t.FailNow()
		}
	}
}

func TestIsNotSet(t *testing.T) {
	if Name("").IsSet() {
		t.FailNow()
	}

	if Alpha2Code("").IsSet() {
		t.FailNow()
	}

	if Alpha3Code("").IsSet() {
		t.FailNow()
	}
}

func TestMappingIsCorrect(t *testing.T) {
	for key, country := range countryByName {
		if key != country.name {
			t.FailNow()
		}
	}

	for key, country := range countryByAlpha3 {
		if key != country.alpha3 {
			t.FailNow()
		}
	}

	for key, country := range countryByAlpha2 {
		if key != country.alpha2 {
			t.FailNow()
		}
	}
}

func TestMappingStringsCorrect(t *testing.T) {
	for key, country := range countryByName {
		if key.String() != country.name.String() {
			t.FailNow()
		}
	}

	for key, country := range countryByAlpha3 {
		if key.String() != country.alpha3.String() {
			t.FailNow()
		}
	}

	for key, country := range countryByAlpha2 {
		if key.String() != country.alpha2.String() {
			t.FailNow()
		}
	}
}

func TestMappingValueCorrect(t *testing.T) {
	for key, country := range countryByName {
		_, actual := key.Value()
		_, expected := country.name.Value()

		if actual != expected {
			t.FailNow()
		}
	}

	for key, country := range countryByAlpha3 {
		_, actual := key.Value()
		_, expected := country.name.Value()

		if actual != expected {
			t.FailNow()
		}
	}

	for key, country := range countryByAlpha2 {
		_, actual := key.Value()
		_, expected := country.name.Value()

		if actual != expected {
			t.FailNow()
		}
	}
}

func TestWrongNameValue(t *testing.T) {
	value, err := Name("a").Value()
	if err == nil || value != nil {
		t.FailNow()
	}
}

func TestWrongAlpha2Value(t *testing.T) {
	value, err := Alpha2Code("a").Value()
	if err == nil || value != nil {
		t.FailNow()
	}
}

func TestWrongAlpha3Value(t *testing.T) {
	value, err := Alpha3Code("a").Value()
	if err == nil || value != nil {
		t.FailNow()
	}
}

func TestNameValidate(t *testing.T) {
	if Name("a").Validate(nil) == nil {
		t.FailNow()
	}

	if NameBrazil.Validate(nil) != nil {
		t.FailNow()
	}
}

func TestAlpha2Validate(t *testing.T) {
	if Alpha2Code("a").Validate(nil) == nil {
		t.FailNow()
	}

	if Alpha2CM.Validate(nil) != nil {
		t.FailNow()
	}
}

func TestAlpha3Validate(t *testing.T) {
	if Alpha3Code("a").Validate(nil) == nil {
		t.FailNow()
	}

	if Alpha3ARG.Validate(nil) != nil {
		t.FailNow()
	}
}

func TestCountryTypes(t *testing.T) {
	for _, country := range countryByAlpha3 {
		if country.name != country.Name() {
			t.FailNow()
		}

		if country.alpha3 != country.Alpha3Code() {
			t.FailNow()
		}

		if country.alpha2 != country.Alpha2Code() {
			t.FailNow()
		}

		if string(country.name) != country.NameStr() {
			t.FailNow()
		}

		if string(country.alpha3) != country.Alpha3CodeStr() {
			t.FailNow()
		}

		if string(country.alpha2) != country.Alpha2CodeStr() {
			t.FailNow()
		}
	}
}

func TestAlpha2Lookup(t *testing.T) {
	if _, ok := ByAlpha2Code("a"); ok {
		t.FailNow()
	}

	if _, ok := ByAlpha2CodeStr("a"); ok {
		t.FailNow()
	}

	if _, err := ByAlpha2CodeErr("a"); err == nil {
		t.FailNow()
	}

	if _, err := ByAlpha2CodeStrErr("a"); err == nil {
		t.FailNow()
	}

	if _, ok := ByAlpha2Code(Canada.Alpha2Code()); !ok {
		t.FailNow()
	}

	if _, ok := ByAlpha2CodeStr(strings.ToLower(Canada.Alpha2Code().String())); !ok {
		t.FailNow()
	}

	if _, err := ByAlpha2CodeErr(Canada.Alpha2Code()); err != nil {
		t.FailNow()
	}

	if _, err := ByAlpha2CodeStrErr(strings.ToLower(Canada.Alpha2Code().String())); err != nil {
		t.FailNow()
	}
}

func TestAlpha3Lookup(t *testing.T) {
	if _, ok := ByAlpha3Code("a"); ok {
		t.FailNow()
	}

	if _, ok := ByAlpha3CodeStr("a"); ok {
		t.FailNow()
	}

	if _, err := ByAlpha3CodeErr("a"); err == nil {
		t.FailNow()
	}

	if _, err := ByAlpha3CodeStrErr("a"); err == nil {
		t.FailNow()
	}

	if _, ok := ByAlpha3Code(Canada.Alpha3Code()); !ok {
		t.FailNow()
	}

	if _, ok := ByAlpha3CodeStr(Canada.Alpha3Code().String()); !ok {
		t.FailNow()
	}

	if _, ok := ByAlpha3CodeStr(strings.ToLower(Canada.Alpha3Code().String())); !ok {
		t.FailNow()
	}

	if _, err := ByAlpha3CodeErr(Canada.Alpha3Code()); err != nil {
		t.FailNow()
	}

	if _, err := ByAlpha3CodeStrErr(strings.ToLower(Canada.Alpha3Code().String())); err != nil {
		t.FailNow()
	}
}

func TestNameLookup(t *testing.T) {
	if _, ok := ByName("a"); ok {
		t.FailNow()
	}

	if _, ok := ByNameStr("a"); ok {
		t.FailNow()
	}

	if _, err := ByNameErr("a"); err == nil {
		t.FailNow()
	}

	if _, err := ByNameStrErr("a"); err == nil {
		t.FailNow()
	}

	if _, ok := ByName(Canada.Name()); !ok {
		t.FailNow()
	}

	if _, ok := ByNameStr(Canada.Name().String()); !ok {
		t.FailNow()
	}

	if _, err := ByNameErr(Canada.Name()); err != nil {
		t.FailNow()
	}

	if _, err := ByNameStrErr(Canada.Name().String()); err != nil {
		t.FailNow()
	}
}
