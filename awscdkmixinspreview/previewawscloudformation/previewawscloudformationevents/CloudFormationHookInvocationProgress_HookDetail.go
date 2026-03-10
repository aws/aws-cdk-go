package previewawscloudformationevents


// Type definition for Hook-detail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   hookDetail := &HookDetail{
//   	Description: []*string{
//   		jsii.String("description"),
//   	},
//   	DocumentationUrl: []*string{
//   		jsii.String("documentationUrl"),
//   	},
//   	FailureMode: []*string{
//   		jsii.String("failureMode"),
//   	},
//   	HookTypeArn: []*string{
//   		jsii.String("hookTypeArn"),
//   	},
//   	HookTypeName: []*string{
//   		jsii.String("hookTypeName"),
//   	},
//   	HookVersion: []*string{
//   		jsii.String("hookVersion"),
//   	},
//   	SourceUrl: []*string{
//   		jsii.String("sourceUrl"),
//   	},
//   }
//
// Experimental.
type CloudFormationHookInvocationProgress_HookDetail struct {
	// description property.
	//
	// Specify an array of string values to match this event if the actual value of description is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Description *[]*string `field:"optional" json:"description" yaml:"description"`
	// documentation-url property.
	//
	// Specify an array of string values to match this event if the actual value of documentation-url is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DocumentationUrl *[]*string `field:"optional" json:"documentationUrl" yaml:"documentationUrl"`
	// failure-mode property.
	//
	// Specify an array of string values to match this event if the actual value of failure-mode is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FailureMode *[]*string `field:"optional" json:"failureMode" yaml:"failureMode"`
	// hook-type-arn property.
	//
	// Specify an array of string values to match this event if the actual value of hook-type-arn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	HookTypeArn *[]*string `field:"optional" json:"hookTypeArn" yaml:"hookTypeArn"`
	// hook-type-name property.
	//
	// Specify an array of string values to match this event if the actual value of hook-type-name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	HookTypeName *[]*string `field:"optional" json:"hookTypeName" yaml:"hookTypeName"`
	// hook-version property.
	//
	// Specify an array of string values to match this event if the actual value of hook-version is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	HookVersion *[]*string `field:"optional" json:"hookVersion" yaml:"hookVersion"`
	// source-url property.
	//
	// Specify an array of string values to match this event if the actual value of source-url is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceUrl *[]*string `field:"optional" json:"sourceUrl" yaml:"sourceUrl"`
}

