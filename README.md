[![Go Report Card](https://goreportcard.com/badge/github.com/mikekonan/go-countries)](https://goreportcard.com/report/github.com/mikekonan/go-countries)
# go-countries
Lightweight lookup over ISO-3166 codes. Each unit implements `driver.Valuer`, `ozzo validation.Validate`, `Stringer`, `json.Unmarshaler` interfaces. 

This library has been created with the purpose to facilitate search and transfer of ISO-3166 country codes.

Created with a codegen - check generator branch.

# Installation
```go get github.com/mikekonan/go-countries```

# Usage:
	//1. use in your structs
	type User struct {
		Name    string             `json:"name" db:"name"`
		Country country.Alpha2Code `json:"country" db:"country"`
	}

	func main() {
		user := User{}
		//2. use in your wire
		json.Unmarshal([]byte(`{"name":"name", "country": "ca"}`), &user)

		//3. check is set
		user.Country.IsSet() //check user country is provided

		//4. validate using ozzo-validation
		if err := validation.ValidateStruct(&user, validation.Field(&user.Country, validation.Required, user.Country)); err != nil {
			log.Fatal(err)
		}

		//5. lookup by alpha2, alpha3, country name
		if userCountry, ok := country.ByAlpha2Code(user.Country); ok {
			fmt.Printf("country name - '%s', alpha-2 - '%s', alpha-3 - '%s'", userCountry.Name(), userCountry.Alpha2Code(), userCountry.Alpha3Code())
		}

		//6. store in db
		fmt.Println(user.Country.Value()) //prints 'CA'

		//7. use specific countries
		fmt.Println(country.Canada.Alpha2Code())
	}
    

# API
### Lookup:
    country.ByNameStrErr()
    country.ByNameErr()
    country.ByNameStr()
    country.ByName()
    
    country.ByAlpha2CodeStrErr()
    country.ByAlpha2CodeErr()
    country.ByAlpha2CodeStr()
    country.ByAlpha2Code()
    
    country.ByAlpha3CodeStrErr()
    country.ByAlpha3CodeErr()
    country.ByAlpha3CodeStr()
    country.ByAlpha3Code()

### Constants:
    ...
    country.Canada
    country.Alpha2CA
    country.Alpha3CAN
    country.NameCanada
    ...

### Types:
    type Country struct {
	    name   Name
	    alpha2 Alpha2Code
	    alpha3 Alpha3Code
    }
    
    func (c Country) Name() Name             { return c.name }
    func (c Country) Alpha2Code() Alpha2Code { return c.alpha2 }
    func (c Country) Alpha3Code() Alpha3Code { return c.alpha3 }
    func (c Country) NameStr() string        { return c.name.String() }
    func (c Country) Alpha2CodeStr() string  { return c.alpha2.String() }
    func (c Country) Alpha3CodeStr() string  { return c.alpha3.String() }
