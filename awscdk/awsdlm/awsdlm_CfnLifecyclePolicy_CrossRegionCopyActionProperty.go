package awsdlm


// Specifies a rule for copying shared snapshots across Regions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   crossRegionCopyActionProperty := &crossRegionCopyActionProperty{
//   	encryptionConfiguration: &encryptionConfigurationProperty{
//   		encrypted: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		cmkArn: jsii.String("cmkArn"),
//   	},
//   	target: jsii.String("target"),
//
//   	// the properties below are optional
//   	retainRule: &crossRegionCopyRetainRuleProperty{
//   		interval: jsii.Number(123),
//   		intervalUnit: jsii.String("intervalUnit"),
//   	},
//   }
//
type CfnLifecyclePolicy_CrossRegionCopyActionProperty struct {
	// The encryption settings for the copied snapshot.
	EncryptionConfiguration interface{} `field:"required" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The target Region.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Specifies the retention rule for cross-Region snapshot copies.
	RetainRule interface{} `field:"optional" json:"retainRule" yaml:"retainRule"`
}

