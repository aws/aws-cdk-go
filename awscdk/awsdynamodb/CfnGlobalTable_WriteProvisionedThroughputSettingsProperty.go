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
//   writeProvisionedThroughputSettingsProperty := &WriteProvisionedThroughputSettingsProperty{
//   	WriteCapacityAutoScalingSettings: &CapacityAutoScalingSettingsProperty{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-writeprovisionedthroughputsettings.html
//
type CfnGlobalTable_WriteProvisionedThroughputSettingsProperty struct {
	// Specifies auto scaling settings for the replica table or global secondary index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-writeprovisionedthroughputsettings.html#cfn-dynamodb-globaltable-writeprovisionedthroughputsettings-writecapacityautoscalingsettings
	//
	WriteCapacityAutoScalingSettings interface{} `field:"optional" json:"writeCapacityAutoScalingSettings" yaml:"writeCapacityAutoScalingSettings"`
}

