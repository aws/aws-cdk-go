package awsdlm


// *[Event-based policies only]* Specifies a cross-Region copy action for event-based policies.
//
// > To specify a cross-Region copy rule for snapshot and AMI policies, use [CrossRegionCopyRule](https://docs.aws.amazon.com/dlm/latest/APIReference/API_CrossRegionCopyRule.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   crossRegionCopyActionProperty := &CrossRegionCopyActionProperty{
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		Encrypted: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		CmkArn: jsii.String("cmkArn"),
//   	},
//   	Target: jsii.String("target"),
//
//   	// the properties below are optional
//   	RetainRule: &CrossRegionCopyRetainRuleProperty{
//   		Interval: jsii.Number(123),
//   		IntervalUnit: jsii.String("intervalUnit"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-crossregioncopyaction.html
//
type CfnLifecyclePolicy_CrossRegionCopyActionProperty struct {
	// The encryption settings for the copied snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-crossregioncopyaction.html#cfn-dlm-lifecyclepolicy-crossregioncopyaction-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"required" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The target Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-crossregioncopyaction.html#cfn-dlm-lifecyclepolicy-crossregioncopyaction-target
	//
	Target *string `field:"required" json:"target" yaml:"target"`
	// Specifies a retention rule for cross-Region snapshot copies created by snapshot or event-based policies, or cross-Region AMI copies created by AMI policies.
	//
	// After the retention period expires, the cross-Region copy is deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-crossregioncopyaction.html#cfn-dlm-lifecyclepolicy-crossregioncopyaction-retainrule
	//
	RetainRule interface{} `field:"optional" json:"retainRule" yaml:"retainRule"`
}

