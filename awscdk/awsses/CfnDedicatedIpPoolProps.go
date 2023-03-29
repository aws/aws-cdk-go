package awsses


// Properties for defining a `CfnDedicatedIpPool`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDedicatedIpPoolProps := &CfnDedicatedIpPoolProps{
//   	PoolName: jsii.String("poolName"),
//   	ScalingMode: jsii.String("scalingMode"),
//   }
//
type CfnDedicatedIpPoolProps struct {
	// The name of the dedicated IP pool that the IP address is associated with.
	PoolName *string `field:"optional" json:"poolName" yaml:"poolName"`
	// The type of scaling mode.
	//
	// The following options are available:
	//
	// - `STANDARD` - The customer controls which IPs are part of the dedicated IP pool.
	// - `MANAGED` - The reputation and number of IPs is automatically managed by Amazon SES .
	//
	// The `STANDARD` option is selected by default if no value is specified.
	ScalingMode *string `field:"optional" json:"scalingMode" yaml:"scalingMode"`
}

