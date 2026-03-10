package previewawsdevopsguruevents


// Type definition for CloudFormation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cloudFormation := &CloudFormation{
//   	StackNames: []*string{
//   		jsii.String("stackNames"),
//   	},
//   }
//
// Experimental.
type DevOpsGuruInsightSeverityUpgraded_CloudFormation struct {
	// stackNames property.
	//
	// Specify an array of string values to match this event if the actual value of stackNames is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StackNames *[]*string `field:"optional" json:"stackNames" yaml:"stackNames"`
}

