package previewawsec2events


// Type definition for LaunchTemplateConfigs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var overrides interface{}
//
//   launchTemplateConfigs := &LaunchTemplateConfigs{
//   	LaunchTemplateSpecification: &LaunchTemplateSpecification{
//   		LaunchTemplateId: []*string{
//   			jsii.String("launchTemplateId"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	Overrides: []interface{}{
//   		overrides,
//   	},
//   	Tag: []*string{
//   		jsii.String("tag"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_LaunchTemplateConfigs struct {
	// LaunchTemplateSpecification property.
	//
	// Specify an array of string values to match this event if the actual value of LaunchTemplateSpecification is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LaunchTemplateSpecification *InstanceEvents_AWSAPICallViaCloudTrail_LaunchTemplateSpecification `field:"optional" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// Overrides property.
	//
	// Specify an array of string values to match this event if the actual value of Overrides is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Overrides *[]interface{} `field:"optional" json:"overrides" yaml:"overrides"`
	// tag property.
	//
	// Specify an array of string values to match this event if the actual value of tag is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tag *[]*string `field:"optional" json:"tag" yaml:"tag"`
}

