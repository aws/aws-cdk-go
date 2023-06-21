package awscustomerprofiles


// The threshold for the calculated attribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   thresholdProperty := &ThresholdProperty{
//   	Operator: jsii.String("operator"),
//   	Value: jsii.String("value"),
//   }
//
type CfnCalculatedAttributeDefinition_ThresholdProperty struct {
	// The operator of the threshold.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// The value of the threshold.
	Value *string `field:"required" json:"value" yaml:"value"`
}

