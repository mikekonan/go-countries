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

func (c Country) Validate(_ interface{}) error {
	if _, ok := ByCountry(c); !ok {
		return fmt.Errorf("'%s' is not valid ISO-3166 country name", c)
	}

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

func (c Alpha2Code) Validate(_ interface{}) error {
	if _, ok := ByAlpha2Code(c); !ok {
		return fmt.Errorf("'%s' is not valid ISO-3166-alpha2 code", c)
	}

	return nil
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

func (c Alpha3Code) Validate(_ interface{}) error {
	if _, ok := ByAlpha3Code(c); !ok {
		return fmt.Errorf("'%s' is not valid ISO-3166-alpha3 code", c)
	}

	return nil
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

func ByAlpha2Code(code Alpha2Code) (result country, ok bool) {
	result, ok = countryByAlpha2[code]
	return
}

func ByCountry(country Country) (result country, ok bool) {
	result, ok = countryByCountry[country]
	return
}
