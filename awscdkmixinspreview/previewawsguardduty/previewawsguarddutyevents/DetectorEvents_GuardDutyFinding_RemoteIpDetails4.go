package previewawsguarddutyevents


// Type definition for RemoteIpDetails_4.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   remoteIpDetails4 := &RemoteIpDetails4{
//   	City: &City4{
//   		CityName: []*string{
//   			jsii.String("cityName"),
//   		},
//   	},
//   	Country: &Country4{
//   		CountryName: []*string{
//   			jsii.String("countryName"),
//   		},
//   	},
//   	GeoLocation: &GeoLocation1{
//   		Lat: []*string{
//   			jsii.String("lat"),
//   		},
//   		Lon: []*string{
//   			jsii.String("lon"),
//   		},
//   	},
//   	IpAddressV4: []*string{
//   		jsii.String("ipAddressV4"),
//   	},
//   	Organization: &Organization4{
//   		Asn: []*string{
//   			jsii.String("asn"),
//   		},
//   		AsnOrg: []*string{
//   			jsii.String("asnOrg"),
//   		},
//   		Isp: []*string{
//   			jsii.String("isp"),
//   		},
//   		Org: []*string{
//   			jsii.String("org"),
//   		},
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_RemoteIpDetails4 struct {
	// city property.
	//
	// Specify an array of string values to match this event if the actual value of city is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	City *DetectorEvents_GuardDutyFinding_City4 `field:"optional" json:"city" yaml:"city"`
	// country property.
	//
	// Specify an array of string values to match this event if the actual value of country is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Country *DetectorEvents_GuardDutyFinding_Country4 `field:"optional" json:"country" yaml:"country"`
	// geoLocation property.
	//
	// Specify an array of string values to match this event if the actual value of geoLocation is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	GeoLocation *DetectorEvents_GuardDutyFinding_GeoLocation1 `field:"optional" json:"geoLocation" yaml:"geoLocation"`
	// ipAddressV4 property.
	//
	// Specify an array of string values to match this event if the actual value of ipAddressV4 is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	IpAddressV4 *[]*string `field:"optional" json:"ipAddressV4" yaml:"ipAddressV4"`
	// organization property.
	//
	// Specify an array of string values to match this event if the actual value of organization is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Organization *DetectorEvents_GuardDutyFinding_Organization4 `field:"optional" json:"organization" yaml:"organization"`
}

