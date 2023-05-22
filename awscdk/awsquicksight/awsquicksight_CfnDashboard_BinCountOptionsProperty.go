package awsquicksight


// The options that determine the bin count of a histogram.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   binCountOptionsProperty := &BinCountOptionsProperty{
//   	Value: jsii.Number(123),
//   }
//
type CfnDashboard_BinCountOptionsProperty struct {
	// The options that determine the bin count value.
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

