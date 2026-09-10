package awsmediaconnect


// The fabric configuration settings for the router output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   fabricConfigurationProperty := &FabricConfigurationProperty{
//   	RecoveryLatencyMode: jsii.String("recoveryLatencyMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-fabricconfiguration.html
//
type CfnRouterOutputPropsMixin_FabricConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-fabricconfiguration.html#cfn-mediaconnect-routeroutput-fabricconfiguration-recoverylatencymode
	//
	RecoveryLatencyMode *string `field:"optional" json:"recoveryLatencyMode" yaml:"recoveryLatencyMode"`
}

