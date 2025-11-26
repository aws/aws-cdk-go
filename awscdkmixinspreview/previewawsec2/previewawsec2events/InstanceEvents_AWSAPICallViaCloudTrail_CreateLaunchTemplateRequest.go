package previewawsec2events


// Type definition for CreateLaunchTemplateRequest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   createLaunchTemplateRequest := &CreateLaunchTemplateRequest{
//   	LaunchTemplateData: &LaunchTemplateData{
//   		ImageId: []*string{
//   			jsii.String("imageId"),
//   		},
//   		InstanceMarketOptions: &InstanceMarketOptions1{
//   			MarketType: []*string{
//   				jsii.String("marketType"),
//   			},
//   			SpotOptions: &SpotOptions2{
//   				MaxPrice: []*string{
//   					jsii.String("maxPrice"),
//   				},
//   				SpotInstanceType: []*string{
//   					jsii.String("spotInstanceType"),
//   				},
//   			},
//   		},
//   		InstanceType: []*string{
//   			jsii.String("instanceType"),
//   		},
//   		NetworkInterface: &NetworkInterface1{
//   			DeviceIndex: []*string{
//   				jsii.String("deviceIndex"),
//   			},
//   			SecurityGroupId: &SecurityGroupId{
//   				Content: []*string{
//   					jsii.String("content"),
//   				},
//   				Tag: []*string{
//   					jsii.String("tag"),
//   				},
//   			},
//   			SubnetId: []*string{
//   				jsii.String("subnetId"),
//   			},
//   			Tag: []*string{
//   				jsii.String("tag"),
//   			},
//   		},
//   		UserData: []*string{
//   			jsii.String("userData"),
//   		},
//   	},
//   	LaunchTemplateName: []*string{
//   		jsii.String("launchTemplateName"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_CreateLaunchTemplateRequest struct {
	// LaunchTemplateData property.
	//
	// Specify an array of string values to match this event if the actual value of LaunchTemplateData is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LaunchTemplateData *InstanceEvents_AWSAPICallViaCloudTrail_LaunchTemplateData `field:"optional" json:"launchTemplateData" yaml:"launchTemplateData"`
	// LaunchTemplateName property.
	//
	// Specify an array of string values to match this event if the actual value of LaunchTemplateName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LaunchTemplateName *[]*string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
}

