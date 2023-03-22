package awsdynamodb


// Allows you to specify the read capacity settings for a replica table or a replica global secondary index when the `BillingMode` is set to `PROVISIONED` .
//
// You must specify a value for either `ReadCapacityUnits` or `ReadCapacityAutoScalingSettings` , but not both. You can switch between fixed capacity and auto scaling.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   readProvisionedThroughputSettingsProperty := &ReadProvisionedThroughputSettingsProperty{
//   	ReadCapacityAutoScalingSettings: &CapacityAutoScalingSettingsProperty{
//   		MaxCapacity: jsii.Number(123),
//   		MinCapacity: jsii.Number(123),
//   		TargetTrackingScalingPolicyConfiguration: &TargetTrackingScalingPolicyConfigurationProperty{
//   			TargetValue: jsii.Number(123),
//
//   			// the properties below are optional
//   			DisableScaleIn: jsii.Boolean(false),
//   			ScaleInCooldown: jsii.Number(123),
//   			ScaleOutCooldown: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		SeedCapacity: jsii.Number(123),
//   	},
//   	ReadCapacityUnits: jsii.Number(123),
//   }
//
type CfnGlobalTable_ReadProvisionedThroughputSettingsProperty struct {
	// Specifies auto scaling settings for the replica table or global secondary index.
	ReadCapacityAutoScalingSettings interface{} `field:"optional" json:"readCapacityAutoScalingSettings" yaml:"readCapacityAutoScalingSettings"`
	// Specifies a fixed read capacity for the replica table or global secondary index.
	ReadCapacityUnits *float64 `field:"optional" json:"readCapacityUnits" yaml:"readCapacityUnits"`
}

