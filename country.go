package country

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type Name string

func (name Name) Value() (value driver.Value, err error) {
	if _, err = ByNameErr(name); err != nil {
		return nil, err
	}

	return name.String(), nil
}

func (name Name) Validate(_ interface{}) (err error) {
	_, err = ByNameErr(name)

	return
}

func (name Name) IsSet() bool {
	return len(string(name)) > 0
}

func (name Name) String() string {
	return string(name)
}

type Alpha2Code string

func (code Alpha2Code) Value() (value driver.Value, err error) {
	if _, err = ByAlpha2CodeErr(code); err != nil {
		return nil, err
	}

	return code.String(), nil
}

func (code Alpha2Code) Validate(_ interface{}) (err error) {
	_, err = ByAlpha2CodeErr(code)

	return
}

func (code Alpha2Code) IsSet() bool {
	return len(string(code)) > 0
}

func (code Alpha2Code) String() string {
	return string(code)
}

type Alpha3Code string

func (code Alpha3Code) Value() (value driver.Value, err error) {
	if _, err = ByAlpha3CodeErr(code); err != nil {
		return nil, err
	}

	return code.String(), nil
}

func (code Alpha3Code) Validate(_ interface{}) (err error) {
	_, err = ByAlpha3CodeErr(code)

	return
}

func (code Alpha3Code) IsSet() bool {
	return len(string(code)) > 0
}

func (code Alpha3Code) String() string {
	return string(code)
}

type Country struct {
	name   Name
	alpha2 Alpha2Code
	alpha3 Alpha3Code
}

func (country Country) Name() Name             { return country.name }
func (country Country) Alpha2Code() Alpha2Code { return country.alpha2 }
func (country Country) Alpha3Code() Alpha3Code { return country.alpha3 }
func (country Country) NameStr() string        { return country.name.String() }
func (country Country) Alpha2CodeStr() string  { return country.alpha2.String() }
func (country Country) Alpha3CodeStr() string  { return country.alpha3.String() }

func ByAlpha3Code(code Alpha3Code) (result Country, ok bool) {
	result, ok = countryByAlpha3[code]
	return
}

func ByAlpha3CodeStr(code string) (result Country, ok bool) {
	return ByAlpha3Code(Alpha3Code(strings.ToUpper(code)))
}

func ByAlpha3CodeErr(code Alpha3Code) (result Country, err error) {
	var ok bool
	result, ok = ByAlpha3Code(code)
	if !ok {
		err = fmt.Errorf("'%s' is not valid ISO-3166-alpha3 code", code)
	}

	return
}

func ByAlpha3CodeStrErr(code string) (result Country, err error) {
	return ByAlpha3CodeErr(Alpha3Code(strings.ToUpper(code)))
}

func ByAlpha2Code(code Alpha2Code) (result Country, ok bool) {
	result, ok = countryByAlpha2[code]
	return
}

func ByAlpha2CodeStr(code string) (result Country, ok bool) {
	return ByAlpha2Code(Alpha2Code(strings.ToUpper(code)))
}

func ByAlpha2CodeErr(code Alpha2Code) (result Country, err error) {
	var ok bool
	result, ok = ByAlpha2Code(code)
	if !ok {
		err = fmt.Errorf("'%s' is not valid ISO-3166-alpha2 code", code)
	}

	return
}

func ByAlpha2CodeStrErr(code string) (result Country, err error) {
	return ByAlpha2CodeErr(Alpha2Code(strings.ToUpper(code)))
}

func ByName(country Name) (result Country, ok bool) {
	result, ok = countryByName[country]
	return
}

func ByNameStr(country string) (result Country, ok bool) {
	return ByName(Name(country))
}

func ByNameErr(country Name) (result Country, err error) {
	var ok bool
	result, ok = ByName(country)
	if !ok {
		err = fmt.Errorf("'%s' is not valid ISO-3166 Country name", country)
	}

	return
}

func ByNameStrErr(country string) (result Country, err error) {
	return ByNameErr(Name(country))
}
