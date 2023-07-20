package awsdlm


// *[Event-based policies only]* Specifies an action for an event-based policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &ActionProperty{
//   	CrossRegionCopy: []interface{}{
//   		&CrossRegionCopyActionProperty{
//   			EncryptionConfiguration: &EncryptionConfigurationProperty{
//   				Encrypted: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				CmkArn: jsii.String("cmkArn"),
//   			},
//   			Target: jsii.String("target"),
//
//   			// the properties below are optional
//   			RetainRule: &CrossRegionCopyRetainRuleProperty{
//   				Interval: jsii.Number(123),
//   				IntervalUnit: jsii.String("intervalUnit"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-action.html
//
type CfnLifecyclePolicy_ActionProperty struct {
	// The rule for copying shared snapshots across Regions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-action.html#cfn-dlm-lifecyclepolicy-action-crossregioncopy
	//
	CrossRegionCopy interface{} `field:"required" json:"crossRegionCopy" yaml:"crossRegionCopy"`
	// A descriptive name for the action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-action.html#cfn-dlm-lifecyclepolicy-action-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
}

