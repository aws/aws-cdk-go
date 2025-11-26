package previewawsguarddutyevents


// Type definition for PortProbeAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   portProbeAction := &PortProbeAction{
//   	Blocked: []*string{
//   		jsii.String("blocked"),
//   	},
//   	PortProbeDetails: []PortProbeActionItem{
//   		&PortProbeActionItem{
//   			LocalIpDetails: &LocalIpDetails1{
//   				IpAddressV4: []*string{
//   					jsii.String("ipAddressV4"),
//   				},
//   			},
//   			LocalPortDetails: &LocalPortDetails1{
//   				Port: []*string{
//   					jsii.String("port"),
//   				},
//   				PortName: []*string{
//   					jsii.String("portName"),
//   				},
//   			},
//   			RemoteIpDetails: &RemoteIpDetails4{
//   				City: &City4{
//   					CityName: []*string{
//   						jsii.String("cityName"),
//   					},
//   				},
//   				Country: &Country4{
//   					CountryName: []*string{
//   						jsii.String("countryName"),
//   					},
//   				},
//   				GeoLocation: &GeoLocation1{
//   					Lat: []*string{
//   						jsii.String("lat"),
//   					},
//   					Lon: []*string{
//   						jsii.String("lon"),
//   					},
//   				},
//   				IpAddressV4: []*string{
//   					jsii.String("ipAddressV4"),
//   				},
//   				Organization: &Organization4{
//   					Asn: []*string{
//   						jsii.String("asn"),
//   					},
//   					AsnOrg: []*string{
//   						jsii.String("asnOrg"),
//   					},
//   					Isp: []*string{
//   						jsii.String("isp"),
//   					},
//   					Org: []*string{
//   						jsii.String("org"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_PortProbeAction struct {
	// blocked property.
	//
	// Specify an array of string values to match this event if the actual value of blocked is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Blocked *[]*string `field:"optional" json:"blocked" yaml:"blocked"`
	// portProbeDetails property.
	//
	// Specify an array of string values to match this event if the actual value of portProbeDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PortProbeDetails *[]*DetectorEvents_GuardDutyFinding_PortProbeActionItem `field:"optional" json:"portProbeDetails" yaml:"portProbeDetails"`
}

