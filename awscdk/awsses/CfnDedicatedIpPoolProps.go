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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-dedicatedippool.html
//
type CfnDedicatedIpPoolProps struct {
	// The name of the dedicated IP pool that the IP address is associated with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-dedicatedippool.html#cfn-ses-dedicatedippool-poolname
	//
	PoolName *string `field:"optional" json:"poolName" yaml:"poolName"`
	// The type of scaling mode.
	//
	// The following options are available:
	//
	// - `STANDARD` - The customer controls which IPs are part of the dedicated IP pool.
	// - `MANAGED` - The reputation and number of IPs are automatically managed by Amazon SES .
	//
	// The `STANDARD` option is selected by default if no value is specified.
	//
	// > Updating *ScalingMode* doesn't require a replacement if you're updating its value from `STANDARD` to `MANAGED` . However, updating *ScalingMode* from `MANAGED` to `STANDARD` is not supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-dedicatedippool.html#cfn-ses-dedicatedippool-scalingmode
	//
	ScalingMode *string `field:"optional" json:"scalingMode" yaml:"scalingMode"`
}

