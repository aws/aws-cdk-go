package previewawssyntheticsevents


// Type definition for TestCodeLayerVersionArn.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   testCodeLayerVersionArn := &TestCodeLayerVersionArn{
//   	CurrentValue: []*string{
//   		jsii.String("currentValue"),
//   	},
//   	PreviousValue: []*string{
//   		jsii.String("previousValue"),
//   	},
//   }
//
// Experimental.
type SyntheticsCanaryStatusChange_TestCodeLayerVersionArn struct {
	// current-value property.
	//
	// Specify an array of string values to match this event if the actual value of current-value is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CurrentValue *[]*string `field:"optional" json:"currentValue" yaml:"currentValue"`
	// previous-value property.
	//
	// Specify an array of string values to match this event if the actual value of previous-value is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PreviousValue *[]*string `field:"optional" json:"previousValue" yaml:"previousValue"`
}

