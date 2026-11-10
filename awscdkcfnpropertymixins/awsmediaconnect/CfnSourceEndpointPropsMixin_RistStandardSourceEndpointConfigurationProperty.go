package awsmediaconnect


// The configuration settings for a source endpoint that uses the RIST protocol.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   ristStandardSourceEndpointConfigurationProperty := &RistStandardSourceEndpointConfigurationProperty{
//   	DestinationIngestPort: jsii.Number(123),
//   	RecoveryLatencyMilliseconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-riststandardsourceendpointconfiguration.html
//
type CfnSourceEndpointPropsMixin_RistStandardSourceEndpointConfigurationProperty struct {
	// The port on the destination that receives the content.
	//
	// The valid range is 3000 to 30000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-riststandardsourceendpointconfiguration.html#cfn-mediaconnect-sourceendpoint-riststandardsourceendpointconfiguration-destinationingestport
	//
	DestinationIngestPort *float64 `field:"optional" json:"destinationIngestPort" yaml:"destinationIngestPort"`
	// The amount of time, in milliseconds, that the source endpoint allows for recovering lost packets.
	//
	// The valid range is 10 to 10000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-riststandardsourceendpointconfiguration.html#cfn-mediaconnect-sourceendpoint-riststandardsourceendpointconfiguration-recoverylatencymilliseconds
	//
	RecoveryLatencyMilliseconds *float64 `field:"optional" json:"recoveryLatencyMilliseconds" yaml:"recoveryLatencyMilliseconds"`
}

