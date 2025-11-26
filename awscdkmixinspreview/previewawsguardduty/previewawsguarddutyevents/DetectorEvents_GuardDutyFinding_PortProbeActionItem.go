package previewawsguarddutyevents


// Type definition for PortProbeActionItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   portProbeActionItem := &PortProbeActionItem{
//   	LocalIpDetails: &LocalIpDetails1{
//   		IpAddressV4: []*string{
//   			jsii.String("ipAddressV4"),
//   		},
//   	},
//   	LocalPortDetails: &LocalPortDetails1{
//   		Port: []*string{
//   			jsii.String("port"),
//   		},
//   		PortName: []*string{
//   			jsii.String("portName"),
//   		},
//   	},
//   	RemoteIpDetails: &RemoteIpDetails4{
//   		City: &City4{
//   			CityName: []*string{
//   				jsii.String("cityName"),
//   			},
//   		},
//   		Country: &Country4{
//   			CountryName: []*string{
//   				jsii.String("countryName"),
//   			},
//   		},
//   		GeoLocation: &GeoLocation1{
//   			Lat: []*string{
//   				jsii.String("lat"),
//   			},
//   			Lon: []*string{
//   				jsii.String("lon"),
//   			},
//   		},
//   		IpAddressV4: []*string{
//   			jsii.String("ipAddressV4"),
//   		},
//   		Organization: &Organization4{
//   			Asn: []*string{
//   				jsii.String("asn"),
//   			},
//   			AsnOrg: []*string{
//   				jsii.String("asnOrg"),
//   			},
//   			Isp: []*string{
//   				jsii.String("isp"),
//   			},
//   			Org: []*string{
//   				jsii.String("org"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_PortProbeActionItem struct {
	// localIpDetails property.
	//
	// Specify an array of string values to match this event if the actual value of localIpDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LocalIpDetails *DetectorEvents_GuardDutyFinding_LocalIpDetails1 `field:"optional" json:"localIpDetails" yaml:"localIpDetails"`
	// localPortDetails property.
	//
	// Specify an array of string values to match this event if the actual value of localPortDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LocalPortDetails *DetectorEvents_GuardDutyFinding_LocalPortDetails1 `field:"optional" json:"localPortDetails" yaml:"localPortDetails"`
	// remoteIpDetails property.
	//
	// Specify an array of string values to match this event if the actual value of remoteIpDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RemoteIpDetails *DetectorEvents_GuardDutyFinding_RemoteIpDetails4 `field:"optional" json:"remoteIpDetails" yaml:"remoteIpDetails"`
}

