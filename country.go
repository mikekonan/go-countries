// This code is auto-generated; DO NOT EDIT.
package country

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type Name string

func (c Name) Value() (value driver.Value, err error) {
	if err = c.Validate(nil); err != nil {
		return nil, err
	}

	return c, nil
}

func (c Name) Validate(_ interface{}) (err error) {
	_, err = ByCountryErr(c)

	return nil
}

func (c Name) IsSet() bool {
	return len(string(c)) > 0
}

type Alpha2Code string

func (c Alpha2Code) Value() (value driver.Value, err error) {
	if err = c.Validate(nil); err != nil {
		return nil, err
	}

	return c, nil
}

func (c Alpha2Code) Validate(_ interface{}) (err error) {
	_, err = ByAlpha2CodeErr(c)

	return
}

func (c Alpha2Code) IsSet() bool {
	return len(string(c)) > 0
}

type Alpha3Code string

func (c Alpha3Code) Value() (value driver.Value, err error) {
	if err = c.Validate(nil); err != nil {
		return nil, err
	}

	return c, nil
}

func (c Alpha3Code) Validate(_ interface{}) (err error) {
	_, err = ByAlpha3CodeErr(c)

	return
}

func (c Alpha3Code) IsSet() bool {
	return len(string(c)) > 0
}

type Country struct {
	country Name
	alpha2  Alpha2Code
	alpha3  Alpha3Code
}

func (c Country) Country() Name          { return c.country }
func (c Country) Alpha2Code() Alpha2Code { return c.alpha2 }
func (c Country) Alpha3Code() Alpha3Code { return c.alpha3 }

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

func ByCountry(country Name) (result Country, ok bool) {
	result, ok = countryByCountry[country]
	return
}

func ByCountryStr(country string) (result Country, ok bool) {
	return ByCountry(Name(strings.ToUpper(country)))
}

func ByCountryErr(country Name) (result Country, err error) {
	var ok bool
	result, ok = ByCountry(country)
	if !ok {
		err = fmt.Errorf("'%s' is not valid ISO-3166 Country name", country)
	}

	return
}

func ByCountryStrErr(country string) (result Country, err error) {
	return ByCountryErr(Name(strings.ToUpper(country)))
}
