package awsguardduty


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagItemProperty := &TagItemProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-trustedentityset-tagitem.html
//
type CfnTrustedEntitySet_TagItemProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-trustedentityset-tagitem.html#cfn-guardduty-trustedentityset-tagitem-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-trustedentityset-tagitem.html#cfn-guardduty-trustedentityset-tagitem-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

