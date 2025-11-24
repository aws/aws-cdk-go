package mixinsawscustomerprofiles


// Information indicating if the Calculated Attribute is ready for use by confirming all historical data has been processed and reflected.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   readinessProperty := &ReadinessProperty{
//   	Message: jsii.String("message"),
//   	ProgressPercentage: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-readiness.html
//
type CfnCalculatedAttributeDefinitionPropsMixin_ReadinessProperty struct {
	// Any customer messaging.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-readiness.html#cfn-customerprofiles-calculatedattributedefinition-readiness-message
	//
	Message *string `field:"optional" json:"message" yaml:"message"`
	// Approximately how far the Calculated Attribute creation is from completion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-readiness.html#cfn-customerprofiles-calculatedattributedefinition-readiness-progresspercentage
	//
	ProgressPercentage *float64 `field:"optional" json:"progressPercentage" yaml:"progressPercentage"`
}

