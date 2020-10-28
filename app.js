const fs = require('fs');

const buff = fs.readFileSync(0);

const codes = JSON.parse(buff.toString());

const countryByCountry =
    codes.map(country => `"${country['English short name']}": {
		country: "${country['English short name']}",
		alpha2:  "${country['Alpha-2 code']}",
		alpha3:  "${country['Alpha-3 code']}",
    }`);

const countryByAlpha2 = codes.map(country => `"${country['Alpha-2 code']}": {
		country: "${country['English short name']}",
		alpha2:  "${country['Alpha-2 code']}",
		alpha3:  "${country['Alpha-3 code']}",
    }`);

const countryByAlpha3 = codes.map(country => `"${country['Alpha-3 code']}": {
		country: "${country['English short name']}",
		alpha2:  "${country['Alpha-2 code']}",
		alpha3:  "${country['Alpha-3 code']}",
    }`);

let template = `// This code is auto-generated; DO NOT EDIT.
package country
 
var countryByCountry = map[Country]country{
	${countryByCountry},
}

var countryByAlpha2 = map[Alpha2Code]country{
	${countryByAlpha2},
}

var countryByAlpha3 = map[Alpha3Code]country{
	${countryByAlpha3},
}
`

console.log(template);