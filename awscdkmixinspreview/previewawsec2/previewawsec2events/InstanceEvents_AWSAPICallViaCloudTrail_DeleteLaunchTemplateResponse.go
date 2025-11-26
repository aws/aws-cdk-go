package previewawsec2events


// Type definition for DeleteLaunchTemplateResponse.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   deleteLaunchTemplateResponse := &DeleteLaunchTemplateResponse{
//   	LaunchTemplate: &LaunchTemplate1{
//   		CreatedBy: []*string{
//   			jsii.String("createdBy"),
//   		},
//   		CreateTime: []*string{
//   			jsii.String("createTime"),
//   		},
//   		DefaultVersionNumber: []*string{
//   			jsii.String("defaultVersionNumber"),
//   		},
//   		LatestVersionNumber: []*string{
//   			jsii.String("latestVersionNumber"),
//   		},
//   		LaunchTemplateId: []*string{
//   			jsii.String("launchTemplateId"),
//   		},
//   		LaunchTemplateName: []*string{
//   			jsii.String("launchTemplateName"),
//   		},
//   	},
//   	RequestId: []*string{
//   		jsii.String("requestId"),
//   	},
//   	Xmlns: []*string{
//   		jsii.String("xmlns"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_DeleteLaunchTemplateResponse struct {
	// launchTemplate property.
	//
	// Specify an array of string values to match this event if the actual value of launchTemplate is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LaunchTemplate *InstanceEvents_AWSAPICallViaCloudTrail_LaunchTemplate1 `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// requestId property.
	//
	// Specify an array of string values to match this event if the actual value of requestId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestId *[]*string `field:"optional" json:"requestId" yaml:"requestId"`
	// xmlns property.
	//
	// Specify an array of string values to match this event if the actual value of xmlns is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Xmlns *[]*string `field:"optional" json:"xmlns" yaml:"xmlns"`
}

