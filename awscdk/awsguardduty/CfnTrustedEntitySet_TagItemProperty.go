package awsguardduty


// Describes a tag.
//
// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
//
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
	// The tag key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-trustedentityset-tagitem.html#cfn-guardduty-trustedentityset-tagitem-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The tag value.
	//
	// This is optional.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-trustedentityset-tagitem.html#cfn-guardduty-trustedentityset-tagitem-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

