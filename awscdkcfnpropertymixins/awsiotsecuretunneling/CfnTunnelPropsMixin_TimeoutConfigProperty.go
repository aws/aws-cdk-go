package awsiotsecuretunneling


// Tunnel timeout configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   timeoutConfigProperty := &TimeoutConfigProperty{
//   	MaxLifetimeTimeoutMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsecuretunneling-tunnel-timeoutconfig.html
//
type CfnTunnelPropsMixin_TimeoutConfigProperty struct {
	// The maximum amount of time (in minutes) a tunnel can remain open.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsecuretunneling-tunnel-timeoutconfig.html#cfn-iotsecuretunneling-tunnel-timeoutconfig-maxlifetimetimeoutminutes
	//
	MaxLifetimeTimeoutMinutes *float64 `field:"optional" json:"maxLifetimeTimeoutMinutes" yaml:"maxLifetimeTimeoutMinutes"`
}

