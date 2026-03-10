package previewawsautoscalingevents


// Type definition for LaunchTemplate_1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var overrides interface{}
//
//   launchTemplate1 := &LaunchTemplate1{
//   	LaunchTemplateSpecification: &LaunchTemplateSpecification{
//   		LaunchTemplateName: []*string{
//   			jsii.String("launchTemplateName"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	Overrides: []interface{}{
//   		overrides,
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_LaunchTemplate1 struct {
	// launchTemplateSpecification property.
	//
	// Specify an array of string values to match this event if the actual value of launchTemplateSpecification is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LaunchTemplateSpecification *AWSAPICallViaCloudTrail_LaunchTemplateSpecification `field:"optional" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// overrides property.
	//
	// Specify an array of string values to match this event if the actual value of overrides is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Overrides *[]interface{} `field:"optional" json:"overrides" yaml:"overrides"`
}

