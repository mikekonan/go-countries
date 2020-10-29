// This code is auto-generated; DO NOT EDIT.
package country

import (
	"database/sql/driver"
	"fmt"
)

type Country string

func (c Country) Value() (value driver.Value, err error) {
	if err = c.Validate(nil); err != nil {
		return nil, err
	}

	return c, nil
}

func (c Country) Validate(_ interface{}) (err error) {
	_, err = ByCountryErr(c)

	return nil
}

func (c Country) IsSet() bool {
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

type country struct {
	country Country
	alpha2  Alpha2Code
	alpha3  Alpha3Code
}

func (c country) Country() Country       { return c.country }
func (c country) Alpha2Code() Alpha2Code { return c.alpha2 }
func (c country) Alpha3Code() Alpha3Code { return c.alpha3 }

func ByAlpha3Code(code Alpha3Code) (result country, ok bool) {
	result, ok = countryByAlpha3[code]
	return
}

func ByAlpha3CodeStr(code string) (result country, ok bool) {
	return ByAlpha3Code(Alpha3Code(code))
}

func ByAlpha3CodeErr(code Alpha3Code) (result country, err error) {
	var ok bool
	result, ok = ByAlpha3Code(code)
	if !ok {
		err = fmt.Errorf("'%s' is not valid ISO-3166-alpha3 code", code)
	}

	return
}

func ByAlpha3CodeStrErr(code string) (result country, err error) {
	return ByAlpha3CodeErr(Alpha3Code(code))
}

func ByAlpha2Code(code Alpha2Code) (result country, ok bool) {
	result, ok = countryByAlpha2[code]
	return
}

func ByAlpha2CodeStr(code string) (result country, ok bool) {
	return ByAlpha2Code(Alpha2Code(code))
}

func ByAlpha2CodeErr(code Alpha2Code) (result country, err error) {
	var ok bool
	result, ok = ByAlpha2Code(code)
	if !ok {
		err = fmt.Errorf("'%s' is not valid ISO-3166-alpha2 code", code)
	}

	return
}

func ByAlpha2CodeStrErr(code string) (result country, err error) {
	return ByAlpha2CodeErr(Alpha2Code(code))
}

func ByCountry(country Country) (result country, ok bool) {
	result, ok = countryByCountry[country]
	return
}

func ByCountryStr(country string) (result country, ok bool) {
	return ByCountry(Country(country))
}

func ByCountryErr(country Country) (result country, err error) {
	var ok bool
	result, ok = ByCountry(country)
	if !ok {
		err = fmt.Errorf("'%s' is not valid ISO-3166 country name", country)
	}

	return
}

func ByCountryStrErr(country string) (result country, err error) {
	return ByCountryErr(Country(country))
}
