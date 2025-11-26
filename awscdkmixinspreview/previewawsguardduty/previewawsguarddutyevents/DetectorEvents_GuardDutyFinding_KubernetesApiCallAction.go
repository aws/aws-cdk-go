package previewawsguarddutyevents


// Type definition for KubernetesApiCallAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kubernetesApiCallAction := &KubernetesApiCallAction{
//   	Parameters: []*string{
//   		jsii.String("parameters"),
//   	},
//   	RemoteIpDetails: &RemoteIpDetails2{
//   		City: &City2{
//   			CityName: []*string{
//   				jsii.String("cityName"),
//   			},
//   		},
//   		Country: &Country2{
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
//   		Organization: &Organization2{
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
//   	RequestUri: []*string{
//   		jsii.String("requestUri"),
//   	},
//   	SourceIPs: []*string{
//   		jsii.String("sourceIPs"),
//   	},
//   	StatusCode: []*string{
//   		jsii.String("statusCode"),
//   	},
//   	UserAgent: []*string{
//   		jsii.String("userAgent"),
//   	},
//   	Verb: []*string{
//   		jsii.String("verb"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_KubernetesApiCallAction struct {
	// parameters property.
	//
	// Specify an array of string values to match this event if the actual value of parameters is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Parameters *[]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// remoteIpDetails property.
	//
	// Specify an array of string values to match this event if the actual value of remoteIpDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RemoteIpDetails *DetectorEvents_GuardDutyFinding_RemoteIpDetails2 `field:"optional" json:"remoteIpDetails" yaml:"remoteIpDetails"`
	// requestUri property.
	//
	// Specify an array of string values to match this event if the actual value of requestUri is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestUri *[]*string `field:"optional" json:"requestUri" yaml:"requestUri"`
	// sourceIPs property.
	//
	// Specify an array of string values to match this event if the actual value of sourceIPs is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceIPs *[]*string `field:"optional" json:"sourceIPs" yaml:"sourceIPs"`
	// statusCode property.
	//
	// Specify an array of string values to match this event if the actual value of statusCode is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StatusCode *[]*string `field:"optional" json:"statusCode" yaml:"statusCode"`
	// userAgent property.
	//
	// Specify an array of string values to match this event if the actual value of userAgent is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserAgent *[]*string `field:"optional" json:"userAgent" yaml:"userAgent"`
	// verb property.
	//
	// Specify an array of string values to match this event if the actual value of verb is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Verb *[]*string `field:"optional" json:"verb" yaml:"verb"`
}

