package previewawssyntheticsevents


// Type definition for Changed-config.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   changedConfig := &ChangedConfig{
//   	ExecutionArn: &ExecutionArn{
//   		CurrentValue: []*string{
//   			jsii.String("currentValue"),
//   		},
//   		PreviousValue: []*string{
//   			jsii.String("previousValue"),
//   		},
//   	},
//   	TestCodeLayerVersionArn: &TestCodeLayerVersionArn{
//   		CurrentValue: []*string{
//   			jsii.String("currentValue"),
//   		},
//   		PreviousValue: []*string{
//   			jsii.String("previousValue"),
//   		},
//   	},
//   }
//
// Experimental.
type SyntheticsCanaryStatusChange_ChangedConfig struct {
	// executionArn property.
	//
	// Specify an array of string values to match this event if the actual value of executionArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExecutionArn *SyntheticsCanaryStatusChange_ExecutionArn `field:"optional" json:"executionArn" yaml:"executionArn"`
	// testCodeLayerVersionArn property.
	//
	// Specify an array of string values to match this event if the actual value of testCodeLayerVersionArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TestCodeLayerVersionArn *SyntheticsCanaryStatusChange_TestCodeLayerVersionArn `field:"optional" json:"testCodeLayerVersionArn" yaml:"testCodeLayerVersionArn"`
}

