package awsguardduty


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   tagItemProperty := &TagItemProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-customdetectionruleassociation-tagitem.html
//
type CfnCustomDetectionRuleAssociationPropsMixin_TagItemProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-customdetectionruleassociation-tagitem.html#cfn-guardduty-customdetectionruleassociation-tagitem-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-customdetectionruleassociation-tagitem.html#cfn-guardduty-customdetectionruleassociation-tagitem-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

