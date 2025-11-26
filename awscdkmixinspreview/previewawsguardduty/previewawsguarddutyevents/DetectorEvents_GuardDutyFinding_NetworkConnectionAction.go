package previewawsguarddutyevents


// Type definition for NetworkConnectionAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkConnectionAction := &NetworkConnectionAction{
//   	Blocked: []*string{
//   		jsii.String("blocked"),
//   	},
//   	ConnectionDirection: []*string{
//   		jsii.String("connectionDirection"),
//   	},
//   	LocalIpDetails: &LocalIpDetails{
//   		IpAddressV4: []*string{
//   			jsii.String("ipAddressV4"),
//   		},
//   	},
//   	LocalPortDetails: &LocalPortDetails{
//   		Port: []*string{
//   			jsii.String("port"),
//   		},
//   		PortName: []*string{
//   			jsii.String("portName"),
//   		},
//   	},
//   	Protocol: []*string{
//   		jsii.String("protocol"),
//   	},
//   	RemoteIpDetails: &RemoteIpDetails3{
//   		City: &City3{
//   			CityName: []*string{
//   				jsii.String("cityName"),
//   			},
//   		},
//   		Country: &Country3{
//   			CountryName: []*string{
//   				jsii.String("countryName"),
//   			},
//   		},
//   		GeoLocation: &GeoLocation{
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
//   		Organization: &Organization3{
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
//   	RemotePortDetails: &RemotePortDetails{
//   		Port: []*string{
//   			jsii.String("port"),
//   		},
//   		PortName: []*string{
//   			jsii.String("portName"),
//   		},
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_NetworkConnectionAction struct {
	// blocked property.
	//
	// Specify an array of string values to match this event if the actual value of blocked is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Blocked *[]*string `field:"optional" json:"blocked" yaml:"blocked"`
	// connectionDirection property.
	//
	// Specify an array of string values to match this event if the actual value of connectionDirection is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ConnectionDirection *[]*string `field:"optional" json:"connectionDirection" yaml:"connectionDirection"`
	// localIpDetails property.
	//
	// Specify an array of string values to match this event if the actual value of localIpDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LocalIpDetails *DetectorEvents_GuardDutyFinding_LocalIpDetails `field:"optional" json:"localIpDetails" yaml:"localIpDetails"`
	// localPortDetails property.
	//
	// Specify an array of string values to match this event if the actual value of localPortDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LocalPortDetails *DetectorEvents_GuardDutyFinding_LocalPortDetails `field:"optional" json:"localPortDetails" yaml:"localPortDetails"`
	// protocol property.
	//
	// Specify an array of string values to match this event if the actual value of protocol is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Protocol *[]*string `field:"optional" json:"protocol" yaml:"protocol"`
	// remoteIpDetails property.
	//
	// Specify an array of string values to match this event if the actual value of remoteIpDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RemoteIpDetails *DetectorEvents_GuardDutyFinding_RemoteIpDetails3 `field:"optional" json:"remoteIpDetails" yaml:"remoteIpDetails"`
	// remotePortDetails property.
	//
	// Specify an array of string values to match this event if the actual value of remotePortDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RemotePortDetails *DetectorEvents_GuardDutyFinding_RemotePortDetails `field:"optional" json:"remotePortDetails" yaml:"remotePortDetails"`
}

