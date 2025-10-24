package awsdynamodb


// Properties for enabling DynamoDB capacity scaling.
//
// Example:
//   import dynamodb "github.com/aws/aws-cdk-go/awscdk"
//
//   var table Table
//
//
//   readCapacity := table.AutoScaleReadCapacity(&EnableScalingProps{
//   	MinCapacity: jsii.Number(10),
//   	MaxCapacity: jsii.Number(1000),
//   })
//   readCapacity.ScaleOnUtilization(&UtilizationScalingProps{
//   	TargetUtilizationPercent: jsii.Number(60),
//   })
//
type EnableScalingProps struct {
	// Maximum capacity to scale to.
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
	// Minimum capacity to scale to.
	MinCapacity *float64 `field:"required" json:"minCapacity" yaml:"minCapacity"`
}

