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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-threshold.html
//
type CfnCalculatedAttributeDefinition_ThresholdProperty struct {
	// The operator of the threshold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-threshold.html#cfn-customerprofiles-calculatedattributedefinition-threshold-operator
	//
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// The value of the threshold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-threshold.html#cfn-customerprofiles-calculatedattributedefinition-threshold-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

