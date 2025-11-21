package awsmediaconnect


// The configuration settings for a standard router output, including the protocol, protocol-specific configuration, network interface, and availability zone.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   standardRouterOutputConfigurationProperty := &StandardRouterOutputConfigurationProperty{
//   	NetworkInterfaceArn: jsii.String("networkInterfaceArn"),
//   	ProtocolConfiguration: &RouterOutputProtocolConfigurationProperty{
//   		Rist: &RistRouterOutputConfigurationProperty{
//   			DestinationAddress: jsii.String("destinationAddress"),
//   			DestinationPort: jsii.Number(123),
//   		},
//   		Rtp: &RtpRouterOutputConfigurationProperty{
//   			DestinationAddress: jsii.String("destinationAddress"),
//   			DestinationPort: jsii.Number(123),
//
//   			// the properties below are optional
//   			ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   		},
//   		SrtCaller: &SrtCallerRouterOutputConfigurationProperty{
//   			DestinationAddress: jsii.String("destinationAddress"),
//   			DestinationPort: jsii.Number(123),
//   			MinimumLatencyMilliseconds: jsii.Number(123),
//
//   			// the properties below are optional
//   			EncryptionConfiguration: &SrtEncryptionConfigurationProperty{
//   				EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   					RoleArn: jsii.String("roleArn"),
//   					SecretArn: jsii.String("secretArn"),
//   				},
//   			},
//   			StreamId: jsii.String("streamId"),
//   		},
//   		SrtListener: &SrtListenerRouterOutputConfigurationProperty{
//   			MinimumLatencyMilliseconds: jsii.Number(123),
//   			Port: jsii.Number(123),
//
//   			// the properties below are optional
//   			EncryptionConfiguration: &SrtEncryptionConfigurationProperty{
//   				EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   					RoleArn: jsii.String("roleArn"),
//   					SecretArn: jsii.String("secretArn"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	Protocol: jsii.String("protocol"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-standardrouteroutputconfiguration.html
//
type CfnRouterOutput_StandardRouterOutputConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the network interface associated with the standard router output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-standardrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-standardrouteroutputconfiguration-networkinterfacearn
	//
	NetworkInterfaceArn *string `field:"required" json:"networkInterfaceArn" yaml:"networkInterfaceArn"`
	// The protocol configuration settings for a router output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-standardrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-standardrouteroutputconfiguration-protocolconfiguration
	//
	ProtocolConfiguration interface{} `field:"required" json:"protocolConfiguration" yaml:"protocolConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-standardrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-standardrouteroutputconfiguration-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

