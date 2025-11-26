package previewawsec2events


// Type definition for LaunchTemplateSpecification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   launchTemplateSpecification := &LaunchTemplateSpecification{
//   	LaunchTemplateId: []*string{
//   		jsii.String("launchTemplateId"),
//   	},
//   	Version: []*string{
//   		jsii.String("version"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_LaunchTemplateSpecification struct {
	// LaunchTemplateId property.
	//
	// Specify an array of string values to match this event if the actual value of LaunchTemplateId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LaunchTemplateId *[]*string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// Version property.
	//
	// Specify an array of string values to match this event if the actual value of Version is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Version *[]*string `field:"optional" json:"version" yaml:"version"`
}

