package awsdlm


// Specifies an action for an event-based policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &actionProperty{
//   	crossRegionCopy: []interface{}{
//   		&crossRegionCopyActionProperty{
//   			encryptionConfiguration: &encryptionConfigurationProperty{
//   				encrypted: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				cmkArn: jsii.String("cmkArn"),
//   			},
//   			target: jsii.String("target"),
//
//   			// the properties below are optional
//   			retainRule: &crossRegionCopyRetainRuleProperty{
//   				interval: jsii.Number(123),
//   				intervalUnit: jsii.String("intervalUnit"),
//   			},
//   		},
//   	},
//   	name: jsii.String("name"),
//   }
//
type CfnLifecyclePolicy_ActionProperty struct {
	// The rule for copying shared snapshots across Regions.
	CrossRegionCopy interface{} `field:"required" json:"crossRegionCopy" yaml:"crossRegionCopy"`
	// A descriptive name for the action.
	Name *string `field:"required" json:"name" yaml:"name"`
}

