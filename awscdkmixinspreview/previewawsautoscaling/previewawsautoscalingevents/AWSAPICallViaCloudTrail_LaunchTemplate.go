package previewawsautoscalingevents


// Type definition for LaunchTemplate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   launchTemplate := &LaunchTemplate{
//   	LaunchTemplateName: []*string{
//   		jsii.String("launchTemplateName"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_LaunchTemplate struct {
	// launchTemplateName property.
	//
	// Specify an array of string values to match this event if the actual value of launchTemplateName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LaunchTemplateName *[]*string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
}

