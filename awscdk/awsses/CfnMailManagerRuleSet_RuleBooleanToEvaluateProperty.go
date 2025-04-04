package awsses


// The union type representing the allowed types of operands for a boolean condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleBooleanToEvaluateProperty := &RuleBooleanToEvaluateProperty{
//   	Attribute: jsii.String("attribute"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulebooleantoevaluate.html
//
type CfnMailManagerRuleSet_RuleBooleanToEvaluateProperty struct {
	// The boolean type representing the allowed attribute types for an email.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulebooleantoevaluate.html#cfn-ses-mailmanagerruleset-rulebooleantoevaluate-attribute
	//
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
}

