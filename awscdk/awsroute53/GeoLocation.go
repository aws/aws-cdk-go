package awsroute53

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Routing based on geographical location.
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
type GeoLocation interface {
	ContinentCode() Continent
	CountryCode() *string
	SubdivisionCode() *string
}

// The jsii proxy struct for GeoLocation
type jsiiProxy_GeoLocation struct {
	_ byte // padding
}

func (j *jsiiProxy_GeoLocation) ContinentCode() Continent {
	var returns Continent
	_jsii_.Get(
		j,
		"continentCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeoLocation) CountryCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeoLocation) SubdivisionCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdivisionCode",
		&returns,
	)
	return returns
}


// Geolocation resource record based on continent code.
//
// Returns: Continent-based geolocation record.
func GeoLocation_Continent(continentCode Continent) GeoLocation {
	_init_.Initialize()

	if err := validateGeoLocation_ContinentParameters(continentCode); err != nil {
		panic(err)
	}
	var returns GeoLocation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.GeoLocation",
		"continent",
		[]interface{}{continentCode},
		&returns,
	)

	return returns
}

// Geolocation resource record based on country code.
//
// Returns: Country-based geolocation record.
// See: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
//
func GeoLocation_Country(countryCode *string) GeoLocation {
	_init_.Initialize()

	if err := validateGeoLocation_CountryParameters(countryCode); err != nil {
		panic(err)
	}
	var returns GeoLocation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.GeoLocation",
		"country",
		[]interface{}{countryCode},
		&returns,
	)

	return returns
}

// Default (wildcard) routing record if no specific geolocation record is found.
//
// Returns: Wildcard routing record.
func GeoLocation_Default() GeoLocation {
	_init_.Initialize()

	var returns GeoLocation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.GeoLocation",
		"default",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Geolocation resource record based on subdivision code (e.g. state of the United States).
// See: https://docs.aws.amazon.com/Route53/latest/APIReference/API_GeoLocation.html#Route53-Type-GeoLocation-SubdivisionCode
//
func GeoLocation_Subdivision(subdivisionCode *string, countryCode *string) GeoLocation {
	_init_.Initialize()

	if err := validateGeoLocation_SubdivisionParameters(subdivisionCode); err != nil {
		panic(err)
	}
	var returns GeoLocation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.GeoLocation",
		"subdivision",
		[]interface{}{subdivisionCode, countryCode},
		&returns,
	)

	return returns
}

