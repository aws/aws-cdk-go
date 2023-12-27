package awsses


// Properties for a dedicated IP pool.
//
// Example:
//   ses.NewDedicatedIpPool(this, jsii.String("Pool"), &DedicatedIpPoolProps{
//   	DedicatedIpPoolName: jsii.String("mypool"),
//   	ScalingMode: ses.ScalingMode_STANDARD,
//   })
//
type DedicatedIpPoolProps struct {
	// A name for the dedicated IP pool.
	//
	// The name must adhere to specific constraints: it can only include
	// lowercase letters (a-z), numbers (0-9), underscores (_), and hyphens (-),
	// and must not exceed 64 characters in length.
	// Default: - a CloudFormation generated name.
	//
	DedicatedIpPoolName *string `field:"optional" json:"dedicatedIpPoolName" yaml:"dedicatedIpPoolName"`
	// The type of scailing mode to use for this IP pool.
	//
	// Updating ScalingMode doesn't require a replacement if you're updating its value from `STANDARD` to `MANAGED`.
	// However, updating ScalingMode from `MANAGED` to `STANDARD` is not supported.
	// Default: ScalingMode.STANDARD
	//
	ScalingMode ScalingMode `field:"optional" json:"scalingMode" yaml:"scalingMode"`
}

