package previewawscustomerprofilesmixins


// The threshold for the calculated attribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   thresholdProperty := &ThresholdProperty{
//   	Operator: jsii.String("operator"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-threshold.html
//
type CfnCalculatedAttributeDefinitionPropsMixin_ThresholdProperty struct {
	// The operator of the threshold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-threshold.html#cfn-customerprofiles-calculatedattributedefinition-threshold-operator
	//
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// The value of the threshold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-threshold.html#cfn-customerprofiles-calculatedattributedefinition-threshold-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

