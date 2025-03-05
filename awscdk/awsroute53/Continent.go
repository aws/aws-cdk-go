package awsroute53


// Continents for geolocation routing.
//
// Example:
//   var myZone hostedZone
//
//
//   // continent
//   // continent
//   route53.NewARecord(this, jsii.String("ARecordGeoLocationContinent"), &ARecordProps{
//   	Zone: myZone,
//   	Target: route53.RecordTarget_FromIpAddresses(jsii.String("1.2.3.0"), jsii.String("5.6.7.0")),
//   	GeoLocation: route53.GeoLocation_Continent(route53.Continent_EUROPE),
//   })
//
//   // country
//   // country
//   route53.NewARecord(this, jsii.String("ARecordGeoLocationCountry"), &ARecordProps{
//   	Zone: myZone,
//   	Target: route53.RecordTarget_*FromIpAddresses(jsii.String("1.2.3.1"), jsii.String("5.6.7.1")),
//   	GeoLocation: route53.GeoLocation_Country(jsii.String("DE")),
//   })
//
//   // subdivision
//   // subdivision
//   route53.NewARecord(this, jsii.String("ARecordGeoLocationSubDividion"), &ARecordProps{
//   	Zone: myZone,
//   	Target: route53.RecordTarget_*FromIpAddresses(jsii.String("1.2.3.2"), jsii.String("5.6.7.2")),
//   	GeoLocation: route53.GeoLocation_Subdivision(jsii.String("WA")),
//   })
//
//   // default (wildcard record if no specific record is found)
//   // default (wildcard record if no specific record is found)
//   route53.NewARecord(this, jsii.String("ARecordGeoLocationDefault"), &ARecordProps{
//   	Zone: myZone,
//   	Target: route53.RecordTarget_*FromIpAddresses(jsii.String("1.2.3.3"), jsii.String("5.6.7.3")),
//   	GeoLocation: route53.GeoLocation_Default(),
//   })
//
type Continent string

const (
	// Africa.
	Continent_AFRICA Continent = "AFRICA"
	// Antarctica.
	Continent_ANTARCTICA Continent = "ANTARCTICA"
	// Asia.
	Continent_ASIA Continent = "ASIA"
	// Europe.
	Continent_EUROPE Continent = "EUROPE"
	// Oceania.
	Continent_OCEANIA Continent = "OCEANIA"
	// North America.
	Continent_NORTH_AMERICA Continent = "NORTH_AMERICA"
	// South America.
	Continent_SOUTH_AMERICA Continent = "SOUTH_AMERICA"
)

