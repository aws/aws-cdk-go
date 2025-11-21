package awsmediaconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var automatic interface{}
//
//   routerOutputConfigurationProperty := &RouterOutputConfigurationProperty{
//   	MediaConnectFlow: &MediaConnectFlowRouterOutputConfigurationProperty{
//   		DestinationTransitEncryption: &FlowTransitEncryptionProperty{
//   			EncryptionKeyConfiguration: &FlowTransitEncryptionKeyConfigurationProperty{
//   				Automatic: automatic,
//   				SecretsManager: &SecretsManagerEncryptionKeyConfigurationProperty{
//   					RoleArn: jsii.String("roleArn"),
//   					SecretArn: jsii.String("secretArn"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			EncryptionKeyType: jsii.String("encryptionKeyType"),
//   		},
//
//   		// the properties below are optional
//   		FlowArn: jsii.String("flowArn"),
//   		FlowSourceArn: jsii.String("flowSourceArn"),
//   	},
//   	MediaLiveInput: &MediaLiveInputRouterOutputConfigurationProperty{
//   		DestinationTransitEncryption: &MediaLiveTransitEncryptionProperty{
//   			EncryptionKeyConfiguration: &MediaLiveTransitEncryptionKeyConfigurationProperty{
//   				Automatic: automatic,
//   				SecretsManager: &SecretsManagerEncryptionKeyConfigurationProperty{
//   					RoleArn: jsii.String("roleArn"),
//   					SecretArn: jsii.String("secretArn"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			EncryptionKeyType: jsii.String("encryptionKeyType"),
//   		},
//
//   		// the properties below are optional
//   		MediaLiveInputArn: jsii.String("mediaLiveInputArn"),
//   		MediaLivePipelineId: jsii.String("mediaLivePipelineId"),
//   	},
//   	Standard: &StandardRouterOutputConfigurationProperty{
//   		NetworkInterfaceArn: jsii.String("networkInterfaceArn"),
//   		ProtocolConfiguration: &RouterOutputProtocolConfigurationProperty{
//   			Rist: &RistRouterOutputConfigurationProperty{
//   				DestinationAddress: jsii.String("destinationAddress"),
//   				DestinationPort: jsii.Number(123),
//   			},
//   			Rtp: &RtpRouterOutputConfigurationProperty{
//   				DestinationAddress: jsii.String("destinationAddress"),
//   				DestinationPort: jsii.Number(123),
//
//   				// the properties below are optional
//   				ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   			},
//   			SrtCaller: &SrtCallerRouterOutputConfigurationProperty{
//   				DestinationAddress: jsii.String("destinationAddress"),
//   				DestinationPort: jsii.Number(123),
//   				MinimumLatencyMilliseconds: jsii.Number(123),
//
//   				// the properties below are optional
//   				EncryptionConfiguration: &SrtEncryptionConfigurationProperty{
//   					EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   						RoleArn: jsii.String("roleArn"),
//   						SecretArn: jsii.String("secretArn"),
//   					},
//   				},
//   				StreamId: jsii.String("streamId"),
//   			},
//   			SrtListener: &SrtListenerRouterOutputConfigurationProperty{
//   				MinimumLatencyMilliseconds: jsii.Number(123),
//   				Port: jsii.Number(123),
//
//   				// the properties below are optional
//   				EncryptionConfiguration: &SrtEncryptionConfigurationProperty{
//   					EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   						RoleArn: jsii.String("roleArn"),
//   						SecretArn: jsii.String("secretArn"),
//   					},
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		Protocol: jsii.String("protocol"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-routeroutputconfiguration.html
//
type CfnRouterOutput_RouterOutputConfigurationProperty struct {
	// Configuration settings for connecting a router output to a MediaConnect flow source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-routeroutputconfiguration.html#cfn-mediaconnect-routeroutput-routeroutputconfiguration-mediaconnectflow
	//
	MediaConnectFlow interface{} `field:"optional" json:"mediaConnectFlow" yaml:"mediaConnectFlow"`
	// Configuration settings for connecting a router output to a MediaLive input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-routeroutputconfiguration.html#cfn-mediaconnect-routeroutput-routeroutputconfiguration-medialiveinput
	//
	MediaLiveInput interface{} `field:"optional" json:"mediaLiveInput" yaml:"mediaLiveInput"`
	// The configuration settings for a standard router output, including the protocol, protocol-specific configuration, network interface, and availability zone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-routeroutputconfiguration.html#cfn-mediaconnect-routeroutput-routeroutputconfiguration-standard
	//
	Standard interface{} `field:"optional" json:"standard" yaml:"standard"`
}

