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
//   readProvisionedThroughputSettingsProperty := &readProvisionedThroughputSettingsProperty{
//   	readCapacityAutoScalingSettings: &capacityAutoScalingSettingsProperty{
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
//   	readCapacityUnits: jsii.Number(123),
//   }
//
type CfnGlobalTable_ReadProvisionedThroughputSettingsProperty struct {
	// Specifies auto scaling settings for the replica table or global secondary index.
	ReadCapacityAutoScalingSettings interface{} `field:"optional" json:"readCapacityAutoScalingSettings" yaml:"readCapacityAutoScalingSettings"`
	// Specifies a fixed read capacity for the replica table or global secondary index.
	ReadCapacityUnits *float64 `field:"optional" json:"readCapacityUnits" yaml:"readCapacityUnits"`
}

