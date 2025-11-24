package mixinsawsdlm


// *[Event-based policies only]* Specifies an action for an event-based policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   actionProperty := &ActionProperty{
//   	CrossRegionCopy: []interface{}{
//   		&CrossRegionCopyActionProperty{
//   			EncryptionConfiguration: &EncryptionConfigurationProperty{
//   				CmkArn: jsii.String("cmkArn"),
//   				Encrypted: jsii.Boolean(false),
//   			},
//   			RetainRule: &CrossRegionCopyRetainRuleProperty{
//   				Interval: jsii.Number(123),
//   				IntervalUnit: jsii.String("intervalUnit"),
//   			},
//   			Target: jsii.String("target"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-action.html
//
type CfnLifecyclePolicyPropsMixin_ActionProperty struct {
	// The rule for copying shared snapshots across Regions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-action.html#cfn-dlm-lifecyclepolicy-action-crossregioncopy
	//
	CrossRegionCopy interface{} `field:"optional" json:"crossRegionCopy" yaml:"crossRegionCopy"`
	// A descriptive name for the action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-action.html#cfn-dlm-lifecyclepolicy-action-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

