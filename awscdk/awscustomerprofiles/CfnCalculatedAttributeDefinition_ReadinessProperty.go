package awscustomerprofiles


// The readiness status of the calculated attribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   readinessProperty := &ReadinessProperty{
//   	Message: jsii.String("message"),
//   	ProgressPercentage: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-readiness.html
//
type CfnCalculatedAttributeDefinition_ReadinessProperty struct {
	// Any information pertaining to the status of the calculated attribute if required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-readiness.html#cfn-customerprofiles-calculatedattributedefinition-readiness-message
	//
	Message *string `field:"optional" json:"message" yaml:"message"`
	// The progress percentage for including historical data in your calculated attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-readiness.html#cfn-customerprofiles-calculatedattributedefinition-readiness-progresspercentage
	//
	ProgressPercentage *float64 `field:"optional" json:"progressPercentage" yaml:"progressPercentage"`
}

