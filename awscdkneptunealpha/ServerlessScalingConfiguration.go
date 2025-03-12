package awscdkneptunealpha


// Example:
//   cluster := neptune.NewDatabaseCluster(this, jsii.String("ServerlessDatabase"), &DatabaseClusterProps{
//   	Vpc: Vpc,
//   	InstanceType: neptune.InstanceType_SERVERLESS(),
//   	ServerlessScalingConfiguration: &ServerlessScalingConfiguration{
//   		MinCapacity: jsii.Number(1),
//   		MaxCapacity: jsii.Number(5),
//   	},
//   })
//
// Experimental.
type ServerlessScalingConfiguration struct {
	// Maximum NCU capacity (min value 2.5 - max value 128).
	// Experimental.
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
	// Minimum NCU capacity (min value 1).
	// Experimental.
	MinCapacity *float64 `field:"required" json:"minCapacity" yaml:"minCapacity"`
}

