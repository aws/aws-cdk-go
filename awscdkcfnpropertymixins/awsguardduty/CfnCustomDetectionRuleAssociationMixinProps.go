package awsguardduty


// Properties for CfnCustomDetectionRuleAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnCustomDetectionRuleAssociationMixinProps := &CfnCustomDetectionRuleAssociationMixinProps{
//   	Mode: jsii.String("mode"),
//   	RuleId: jsii.String("ruleId"),
//   	Tags: []TagItemProperty{
//   		&TagItemProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-customdetectionruleassociation.html
//
type CfnCustomDetectionRuleAssociationMixinProps struct {
	// Whether the rule runs in LIVE mode (generates findings) or DRY_RUN mode (evaluates without generating findings).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-customdetectionruleassociation.html#cfn-guardduty-customdetectionruleassociation-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The catalog identifier of the custom detection rule to associate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-customdetectionruleassociation.html#cfn-guardduty-customdetectionruleassociation-ruleid
	//
	RuleId *string `field:"optional" json:"ruleId" yaml:"ruleId"`
	// The tags applied to the association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-customdetectionruleassociation.html#cfn-guardduty-customdetectionruleassociation-tags
	//
	Tags *[]*CfnCustomDetectionRuleAssociationPropsMixin_TagItemProperty `field:"optional" json:"tags" yaml:"tags"`
}

