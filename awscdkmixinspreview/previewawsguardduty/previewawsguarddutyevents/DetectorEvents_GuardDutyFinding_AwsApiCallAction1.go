package previewawsguarddutyevents


// Type definition for AwsApiCallAction_1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   awsApiCallAction1 := &AwsApiCallAction1{
//   	AffectedResources: &AffectedResources1{
//   		AwsCloudTrailTrail: []*string{
//   			jsii.String("awsCloudTrailTrail"),
//   		},
//   		AwsEc2Instance: []*string{
//   			jsii.String("awsEc2Instance"),
//   		},
//   		AwsS3Bucket: []*string{
//   			jsii.String("awsS3Bucket"),
//   		},
//   	},
//   	Api: []*string{
//   		jsii.String("api"),
//   	},
//   	CallerType: []*string{
//   		jsii.String("callerType"),
//   	},
//   	ErrorCode: []*string{
//   		jsii.String("errorCode"),
//   	},
//   	RemoteAccountDetails: &RemoteAccountDetails{
//   		AccountId: []*string{
//   			jsii.String("accountId"),
//   		},
//   		Affiliated: []*string{
//   			jsii.String("affiliated"),
//   		},
//   	},
//   	RemoteIpDetails: &RemoteIpDetails1{
//   		City: &City1{
//   			CityName: []*string{
//   				jsii.String("cityName"),
//   			},
//   		},
//   		Country: &Country1{
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
//   		Organization: &Organization1{
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
//   	ServiceName: []*string{
//   		jsii.String("serviceName"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_AwsApiCallAction1 struct {
	// affectedResources property.
	//
	// Specify an array of string values to match this event if the actual value of affectedResources is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AffectedResources *DetectorEvents_GuardDutyFinding_AffectedResources1 `field:"optional" json:"affectedResources" yaml:"affectedResources"`
	// api property.
	//
	// Specify an array of string values to match this event if the actual value of api is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Api *[]*string `field:"optional" json:"api" yaml:"api"`
	// callerType property.
	//
	// Specify an array of string values to match this event if the actual value of callerType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CallerType *[]*string `field:"optional" json:"callerType" yaml:"callerType"`
	// errorCode property.
	//
	// Specify an array of string values to match this event if the actual value of errorCode is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ErrorCode *[]*string `field:"optional" json:"errorCode" yaml:"errorCode"`
	// remoteAccountDetails property.
	//
	// Specify an array of string values to match this event if the actual value of remoteAccountDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RemoteAccountDetails *DetectorEvents_GuardDutyFinding_RemoteAccountDetails `field:"optional" json:"remoteAccountDetails" yaml:"remoteAccountDetails"`
	// remoteIpDetails property.
	//
	// Specify an array of string values to match this event if the actual value of remoteIpDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RemoteIpDetails *DetectorEvents_GuardDutyFinding_RemoteIpDetails1 `field:"optional" json:"remoteIpDetails" yaml:"remoteIpDetails"`
	// serviceName property.
	//
	// Specify an array of string values to match this event if the actual value of serviceName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ServiceName *[]*string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

