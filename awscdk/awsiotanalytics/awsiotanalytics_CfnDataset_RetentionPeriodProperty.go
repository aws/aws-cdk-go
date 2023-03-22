package awsiotanalytics


// How long, in days, message data is kept.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   retentionPeriodProperty := &retentionPeriodProperty{
//   	numberOfDays: jsii.Number(123),
//   	unlimited: jsii.Boolean(false),
//   }
//
type CfnDataset_RetentionPeriodProperty struct {
	// The number of days that message data is kept.
	//
	// The `unlimited` parameter must be false.
	NumberOfDays *float64 `field:"optional" json:"numberOfDays" yaml:"numberOfDays"`
	// If true, message data is kept indefinitely.
	Unlimited interface{} `field:"optional" json:"unlimited" yaml:"unlimited"`
}

