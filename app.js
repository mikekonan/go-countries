const fs = require('fs');
const transliterate = require('transliteration');

let codes = JSON.parse(fs.readFileSync(0, 'utf8'));

codes = codes.map(code => {
  let result = {};
  result.a2 = {
    key: `Alpha2${code['Alpha-2 code']}`,
    value: code['Alpha-2 code']
  }

              result.a3 = {
    key: `Alpha3${code['Alpha-3 code']}`,
    value: code['Alpha-3 code']
  }

                          result.name = {value: code['English short name']}

  let enumCountryKey = code['English short name']

      if (enumCountryKey === 'Korea (the Republic of)') {
    enumCountryKey = 'South Korea'
  }

  switch (enumCountryKey) {
    case 'Korea (the Democratic People\'s Republic of)': {
      enumCountryKey = 'North Korea'
      break
    }
    case 'Korea (the Republic of)': {
      enumCountryKey = 'South Korea'
      break
    }
    case 'Virgin Islands (British)': {
      enumCountryKey = 'British Virgin Islands'
      break
    }
    case 'Virgin Islands (U.S.)': {
      enumCountryKey = 'US Virgin Islands'
      break
    }
    case 'Palestine, State of': {
      enumCountryKey = 'Palestine'
      break
    }
    case 'Congo (the Democratic Republic of the)': {
      enumCountryKey = 'Democratic Republic of the Congo'
      break
    }
  }

  enumCountryKey = transliterate.transliterate(enumCountryKey)
  enumCountryKey = enumCountryKey.replace(',', '')
  enumCountryKey = enumCountryKey.replace('.', '')
  enumCountryKey = enumCountryKey.replace('\'', '')
  enumCountryKey = enumCountryKey.replace('-', '')
  enumCountryKey = enumCountryKey.replace('*', '')
  enumCountryKey = enumCountryKey.replace(/( ?)\(.*\)( ?)/, '')
  enumCountryKey = enumCountryKey.replace(/( ?)\[.*\]( ?)/, '')
  enumCountryKey = enumCountryKey.split(' ')
                       .map(str => `${str[0].toUpperCase()}${str.slice(1)}`)
                       .join('')

  result.name.key = `Name${enumCountryKey}`
  result.key = enumCountryKey

  return result
})

const countriesTemplate = `package country

const (
${
    codes
        .map(
            code =>
                `//${code.name.key} represents '${code.name.value}' country name
${code.name.key} = Name("${code.name.value}")`)
        .join('\n')}
) 
`;

const alpha2Template = `package country

const (
${
    codes
        .map(
            code => `//${code.a2.key} represents '${
                code.a2.value}' country alpha-2 code
    ${code.a2.key} = Alpha2Code("${code.a2.value}")`)
        .join('\n')}
) 
`;

const alpha3Template = `package country

const (
${
    codes
        .map(
            code => `//${code.a3.key} represents '${
                code.a3.value}' country alpha-3 code
    ${code.a3.key} = Alpha3Code("${code.a3.value}")`)
        .join('\n')}
) 
`;

const entitiesTemplate =
    `package country

var (
${
        codes
            .map(code => `//${code.key} represents '${code.name.value}' country
            ${code.key} = Country{
		name: ${code.name.key},
		alpha2:  ${code.a2.key},
		alpha3:  ${code.a3.key},
	}`).join('\n')}
) 
`;

const countryByCountryTemplate = `package country

var countryByName = map[Name]Country{
	${codes.map(code => `${code.name.key} : ${code.key}`).join(',\n')},
}

var countryByAlpha2 = map[Alpha2Code]Country{
	${codes.map(code => `${code.a2.key} : ${code.key}`).join(',\n')},
}

var countryByAlpha3 = map[Alpha3Code]Country{
	${codes.map(code => `${code.a3.key} : ${code.key}`).join(',\n')},
}
`;

fs.writeFileSync(
    'name_gen.go', countriesTemplate, {encoding: 'utf8', flag: 'w'});

fs.writeFileSync(
    'alpha2_gen.go', alpha2Template, {encoding: 'utf8', flag: 'w'});

fs.writeFileSync(
    'alpha3_gen.go', alpha3Template, {encoding: 'utf8', flag: 'w'});

fs.writeFileSync(
    'entities_gen.go', entitiesTemplate, {encoding: 'utf8', flag: 'w'});

fs.writeFileSync(
    'country_mapping_gen.go', countryByCountryTemplate,
    {encoding: 'utf8', flag: 'w'});
