package awsdynamodb


// Configures a scalable target and an autoscaling policy for a table or global secondary index's read or write capacity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityAutoScalingSettingsProperty := &capacityAutoScalingSettingsProperty{
//   	maxCapacity: jsii.Number(123),
//   	minCapacity: jsii.Number(123),
//   	targetTrackingScalingPolicyConfiguration: &targetTrackingScalingPolicyConfigurationProperty{
//   		targetValue: jsii.Number(123),
//
//   		// the properties below are optional
//   		disableScaleIn: jsii.Boolean(false),
//   		scaleInCooldown: jsii.Number(123),
//   		scaleOutCooldown: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	seedCapacity: jsii.Number(123),
//   }
//
type CfnGlobalTable_CapacityAutoScalingSettingsProperty struct {
	// The maximum provisioned capacity units for the global table.
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
	// The minimum provisioned capacity units for the global table.
	MinCapacity *float64 `field:"required" json:"minCapacity" yaml:"minCapacity"`
	// Defines a target tracking scaling policy.
	TargetTrackingScalingPolicyConfiguration interface{} `field:"required" json:"targetTrackingScalingPolicyConfiguration" yaml:"targetTrackingScalingPolicyConfiguration"`
	// When switching billing mode from `PAY_PER_REQUEST` to `PROVISIONED` , DynamoDB requires you to specify read and write capacity unit values for the table and for each global secondary index.
	//
	// These values will be applied to all replicas. The table will use these provisioned values until CloudFormation creates the autoscaling policies you configured in your template. CloudFormation cannot determine what capacity the table and its global secondary indexes will require in this time period, since they are application-dependent.
	//
	// If you want to switch a table's billing mode from `PAY_PER_REQUEST` to `PROVISIONED` , you must specify a value for this property for each autoscaled resource. If you specify different values for the same resource in different regions, CloudFormation will use the highest value found in either the `SeedCapacity` or `ReadCapacityUnits` properties. For example, if your global secondary index `myGSI` has a `SeedCapacity` of 10 in us-east-1 and a fixed `ReadCapacityUnits` of 20 in eu-west-1, CloudFormation will initially set the read capacity for `myGSI` to 20. Note that if you disable `ScaleIn` for `myGSI` in us-east-1, its read capacity units might not be set back to 10.
	//
	// You must also specify a value for `SeedCapacity` when you plan to switch a table's billing mode from `PROVISIONED` to `PAY_PER_REQUEST` , because CloudFormation might need to roll back the operation (reverting the billing mode to `PROVISIONED` ) and this cannot succeed without specifying a value for `SeedCapacity` .
	SeedCapacity *float64 `field:"optional" json:"seedCapacity" yaml:"seedCapacity"`
}

