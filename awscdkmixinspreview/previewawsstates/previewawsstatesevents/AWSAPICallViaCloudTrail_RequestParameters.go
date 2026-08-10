package previewawsstatesevents


// Type definition for RequestParameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var tags interface{}
//
//   requestParameters := &RequestParameters{
//   	ActivityArn: []*string{
//   		jsii.String("activityArn"),
//   	},
//   	Definition: []*string{
//   		jsii.String("definition"),
//   	},
//   	ExecutionArn: []*string{
//   		jsii.String("executionArn"),
//   	},
//   	Input: []*string{
//   		jsii.String("input"),
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   	RoleArn: []*string{
//   		jsii.String("roleArn"),
//   	},
//   	StateMachineArn: []*string{
//   		jsii.String("stateMachineArn"),
//   	},
//   	Tags: []interface{}{
//   		tags,
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_RequestParameters struct {
	// activityArn property.
	//
	// Specify an array of string values to match this event if the actual value of activityArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ActivityArn *[]*string `field:"optional" json:"activityArn" yaml:"activityArn"`
	// definition property.
	//
	// Specify an array of string values to match this event if the actual value of definition is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Definition *[]*string `field:"optional" json:"definition" yaml:"definition"`
	// executionArn property.
	//
	// Specify an array of string values to match this event if the actual value of executionArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExecutionArn *[]*string `field:"optional" json:"executionArn" yaml:"executionArn"`
	// input property.
	//
	// Specify an array of string values to match this event if the actual value of input is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Input *[]*string `field:"optional" json:"input" yaml:"input"`
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
	// roleArn property.
	//
	// Specify an array of string values to match this event if the actual value of roleArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RoleArn *[]*string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// stateMachineArn property.
	//
	// Specify an array of string values to match this event if the actual value of stateMachineArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StateMachineArn *[]*string `field:"optional" json:"stateMachineArn" yaml:"stateMachineArn"`
	// tags property.
	//
	// Specify an array of string values to match this event if the actual value of tags is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tags *[]interface{} `field:"optional" json:"tags" yaml:"tags"`
}

