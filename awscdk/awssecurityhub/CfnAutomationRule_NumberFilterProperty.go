package awssecurityhub


// A number filter for querying findings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   numberFilterProperty := &NumberFilterProperty{
//   	Eq: jsii.Number(123),
//   	Gte: jsii.Number(123),
//   	Lte: jsii.Number(123),
//   }
//
type CfnAutomationRule_NumberFilterProperty struct {
	// The equal-to condition to be applied to a single field when querying for findings.
	Eq *float64 `field:"optional" json:"eq" yaml:"eq"`
	// The greater-than-equal condition to be applied to a single field when querying for findings.
	Gte *float64 `field:"optional" json:"gte" yaml:"gte"`
	// The less-than-equal condition to be applied to a single field when querying for findings.
	Lte *float64 `field:"optional" json:"lte" yaml:"lte"`
}

