package previewawsgameliftmixins


// Configuration for player gateway.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   playerGatewayConfigurationProperty := &PlayerGatewayConfigurationProperty{
//   	GameServerIpProtocolSupported: jsii.String("gameServerIpProtocolSupported"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-playergatewayconfiguration.html
//
type CfnFleetPropsMixin_PlayerGatewayConfigurationProperty struct {
	// The IP protocol supported by the game server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-playergatewayconfiguration.html#cfn-gamelift-fleet-playergatewayconfiguration-gameserveripprotocolsupported
	//
	GameServerIpProtocolSupported *string `field:"optional" json:"gameServerIpProtocolSupported" yaml:"gameServerIpProtocolSupported"`
}

