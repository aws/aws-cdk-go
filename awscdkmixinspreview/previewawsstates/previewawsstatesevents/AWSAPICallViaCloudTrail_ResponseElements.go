package previewawsstatesevents


// Type definition for ResponseElements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   responseElements := &ResponseElements{
//   	ActivityArn: []*string{
//   		jsii.String("activityArn"),
//   	},
//   	CreationDate: []*string{
//   		jsii.String("creationDate"),
//   	},
//   	ExecutionArn: []*string{
//   		jsii.String("executionArn"),
//   	},
//   	StartDate: []*string{
//   		jsii.String("startDate"),
//   	},
//   	StateMachineArn: []*string{
//   		jsii.String("stateMachineArn"),
//   	},
//   	StopDate: []*string{
//   		jsii.String("stopDate"),
//   	},
//   	UpdateDate: []*string{
//   		jsii.String("updateDate"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_ResponseElements struct {
	// activityArn property.
	//
	// Specify an array of string values to match this event if the actual value of activityArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ActivityArn *[]*string `field:"optional" json:"activityArn" yaml:"activityArn"`
	// creationDate property.
	//
	// Specify an array of string values to match this event if the actual value of creationDate is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreationDate *[]*string `field:"optional" json:"creationDate" yaml:"creationDate"`
	// executionArn property.
	//
	// Specify an array of string values to match this event if the actual value of executionArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExecutionArn *[]*string `field:"optional" json:"executionArn" yaml:"executionArn"`
	// startDate property.
	//
	// Specify an array of string values to match this event if the actual value of startDate is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartDate *[]*string `field:"optional" json:"startDate" yaml:"startDate"`
	// stateMachineArn property.
	//
	// Specify an array of string values to match this event if the actual value of stateMachineArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StateMachineArn *[]*string `field:"optional" json:"stateMachineArn" yaml:"stateMachineArn"`
	// stopDate property.
	//
	// Specify an array of string values to match this event if the actual value of stopDate is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StopDate *[]*string `field:"optional" json:"stopDate" yaml:"stopDate"`
	// updateDate property.
	//
	// Specify an array of string values to match this event if the actual value of updateDate is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UpdateDate *[]*string `field:"optional" json:"updateDate" yaml:"updateDate"`
}

