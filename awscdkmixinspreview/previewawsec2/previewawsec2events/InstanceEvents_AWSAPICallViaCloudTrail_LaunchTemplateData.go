package previewawsec2events


// Type definition for LaunchTemplateData.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   launchTemplateData := &LaunchTemplateData{
//   	ImageId: []*string{
//   		jsii.String("imageId"),
//   	},
//   	InstanceMarketOptions: &InstanceMarketOptions1{
//   		MarketType: []*string{
//   			jsii.String("marketType"),
//   		},
//   		SpotOptions: &SpotOptions2{
//   			MaxPrice: []*string{
//   				jsii.String("maxPrice"),
//   			},
//   			SpotInstanceType: []*string{
//   				jsii.String("spotInstanceType"),
//   			},
//   		},
//   	},
//   	InstanceType: []*string{
//   		jsii.String("instanceType"),
//   	},
//   	NetworkInterface: &NetworkInterface1{
//   		DeviceIndex: []*string{
//   			jsii.String("deviceIndex"),
//   		},
//   		SecurityGroupId: &SecurityGroupId{
//   			Content: []*string{
//   				jsii.String("content"),
//   			},
//   			Tag: []*string{
//   				jsii.String("tag"),
//   			},
//   		},
//   		SubnetId: []*string{
//   			jsii.String("subnetId"),
//   		},
//   		Tag: []*string{
//   			jsii.String("tag"),
//   		},
//   	},
//   	UserData: []*string{
//   		jsii.String("userData"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_LaunchTemplateData struct {
	// ImageId property.
	//
	// Specify an array of string values to match this event if the actual value of ImageId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ImageId *[]*string `field:"optional" json:"imageId" yaml:"imageId"`
	// InstanceMarketOptions property.
	//
	// Specify an array of string values to match this event if the actual value of InstanceMarketOptions is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceMarketOptions *InstanceEvents_AWSAPICallViaCloudTrail_InstanceMarketOptions1 `field:"optional" json:"instanceMarketOptions" yaml:"instanceMarketOptions"`
	// InstanceType property.
	//
	// Specify an array of string values to match this event if the actual value of InstanceType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceType *[]*string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// NetworkInterface property.
	//
	// Specify an array of string values to match this event if the actual value of NetworkInterface is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NetworkInterface *InstanceEvents_AWSAPICallViaCloudTrail_NetworkInterface1 `field:"optional" json:"networkInterface" yaml:"networkInterface"`
	// UserData property.
	//
	// Specify an array of string values to match this event if the actual value of UserData is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserData *[]*string `field:"optional" json:"userData" yaml:"userData"`
}

