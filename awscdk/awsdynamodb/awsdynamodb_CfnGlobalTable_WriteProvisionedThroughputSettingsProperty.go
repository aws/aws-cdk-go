package awsdynamodb


// Specifies an auto scaling policy for write capacity.
//
// This policy will be applied to all replicas. This setting must be specified if `BillingMode` is set to `PROVISIONED` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   writeProvisionedThroughputSettingsProperty := &writeProvisionedThroughputSettingsProperty{
//   	writeCapacityAutoScalingSettings: &capacityAutoScalingSettingsProperty{
//   		maxCapacity: jsii.Number(123),
//   		minCapacity: jsii.Number(123),
//   		targetTrackingScalingPolicyConfiguration: &targetTrackingScalingPolicyConfigurationProperty{
//   			targetValue: jsii.Number(123),
//
//   			// the properties below are optional
//   			disableScaleIn: jsii.Boolean(false),
//   			scaleInCooldown: jsii.Number(123),
//   			scaleOutCooldown: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		seedCapacity: jsii.Number(123),
//   	},
//   }
//
type CfnGlobalTable_WriteProvisionedThroughputSettingsProperty struct {
	// Specifies auto scaling settings for the replica table or global secondary index.
	WriteCapacityAutoScalingSettings interface{} `field:"optional" json:"writeCapacityAutoScalingSettings" yaml:"writeCapacityAutoScalingSettings"`
}

