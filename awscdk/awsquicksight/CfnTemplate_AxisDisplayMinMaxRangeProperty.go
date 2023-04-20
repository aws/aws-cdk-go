package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   axisDisplayMinMaxRangeProperty := &AxisDisplayMinMaxRangeProperty{
//   	Maximum: jsii.Number(123),
//   	Minimum: jsii.Number(123),
//   }
//
type CfnTemplate_AxisDisplayMinMaxRangeProperty struct {
	// `CfnTemplate.AxisDisplayMinMaxRangeProperty.Maximum`.
	Maximum *float64 `field:"optional" json:"maximum" yaml:"maximum"`
	// `CfnTemplate.AxisDisplayMinMaxRangeProperty.Minimum`.
	Minimum *float64 `field:"optional" json:"minimum" yaml:"minimum"`
}

