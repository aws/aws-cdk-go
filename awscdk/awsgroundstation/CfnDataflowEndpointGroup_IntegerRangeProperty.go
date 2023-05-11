package awsgroundstation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integerRangeProperty := &IntegerRangeProperty{
//   	Maximum: jsii.Number(123),
//   	Minimum: jsii.Number(123),
//   }
//
type CfnDataflowEndpointGroup_IntegerRangeProperty struct {
	// `CfnDataflowEndpointGroup.IntegerRangeProperty.Maximum`.
	Maximum *float64 `field:"optional" json:"maximum" yaml:"maximum"`
	// `CfnDataflowEndpointGroup.IntegerRangeProperty.Minimum`.
	Minimum *float64 `field:"optional" json:"minimum" yaml:"minimum"`
}

