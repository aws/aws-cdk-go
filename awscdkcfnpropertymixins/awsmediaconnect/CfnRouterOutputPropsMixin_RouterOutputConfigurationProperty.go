package awsmediaconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var automatic interface{}
//   var public_ interface{}
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
//   			EncryptionKeyType: jsii.String("encryptionKeyType"),
//   		},
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
//   			EncryptionKeyType: jsii.String("encryptionKeyType"),
//   		},
//   		MediaLiveInputArn: jsii.String("mediaLiveInputArn"),
//   		MediaLivePipelineId: jsii.String("mediaLivePipelineId"),
//   	},
//   	SourceEndpoint: &SourceEndpointRouterOutputConfigurationProperty{
//   		DestinationTransitEncryption: &EndpointEncryptionConfigurationProperty{
//   			Automatic: automatic,
//   			SecretsManager: &SecretsManagerEndpointEncryptionConfigurationProperty{
//   				EncryptionSecret: &EndpointEncryptionSecretConfigurationProperty{
//   					SecretArn: jsii.String("secretArn"),
//   				},
//   				EncryptionSecretType: jsii.String("encryptionSecretType"),
//   			},
//   		},
//   		DestinationTransitEncryptionType: jsii.String("destinationTransitEncryptionType"),
//   		EndpointLifecycleManagement: jsii.String("endpointLifecycleManagement"),
//   	},
//   	Standard: &StandardRouterOutputConfigurationProperty{
//   		NetworkInterfaceArn: jsii.String("networkInterfaceArn"),
//   		Protocol: jsii.String("protocol"),
//   		ProtocolConfiguration: &RouterOutputProtocolConfigurationProperty{
//   			Rist: &RistRouterOutputConfigurationProperty{
//   				DestinationAddress: jsii.String("destinationAddress"),
//   				DestinationPort: jsii.Number(123),
//   			},
//   			RtmpPush: &RtmpPushRouterOutputConfigurationProperty{
//   				ApplicationName: jsii.String("applicationName"),
//   				DestinationAddress: jsii.String("destinationAddress"),
//   				DestinationPort: jsii.Number(123),
//   				StreamName: jsii.String("streamName"),
//   				TlsEncryption: &TlsEncryptionProperty{
//   					EncryptionConfiguration: &TlsEncryptionConfigurationProperty{
//   						Public: public_,
//   					},
//   					EncryptionType: jsii.String("encryptionType"),
//   				},
//   			},
//   			Rtp: &RtpRouterOutputConfigurationProperty{
//   				DestinationAddress: jsii.String("destinationAddress"),
//   				DestinationPort: jsii.Number(123),
//   				ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   			},
//   			SrtCaller: &SrtCallerRouterOutputConfigurationProperty{
//   				DestinationAddress: jsii.String("destinationAddress"),
//   				DestinationPort: jsii.Number(123),
//   				EncryptionConfiguration: &SrtEncryptionConfigurationProperty{
//   					EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   						RoleArn: jsii.String("roleArn"),
//   						SecretArn: jsii.String("secretArn"),
//   					},
//   				},
//   				MinimumLatencyMilliseconds: jsii.Number(123),
//   				StreamId: jsii.String("streamId"),
//   			},
//   			SrtListener: &SrtListenerRouterOutputConfigurationProperty{
//   				EncryptionConfiguration: &SrtEncryptionConfigurationProperty{
//   					EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   						RoleArn: jsii.String("roleArn"),
//   						SecretArn: jsii.String("secretArn"),
//   					},
//   				},
//   				MinimumLatencyMilliseconds: jsii.Number(123),
//   				Port: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-routeroutputconfiguration.html
//
type CfnRouterOutputPropsMixin_RouterOutputConfigurationProperty struct {
	// Configuration settings for connecting a router output to a MediaConnect flow source.
	//
	// To connect the router output to the flow source, provide both flowArn and flowSourceArn. To disconnect it, call the update operation and omit both fields while keeping this configuration object. You can't change destinationTransitEncryption while the router output is connected to a flow source. To change it, first disconnect the router output, then update the encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-routeroutputconfiguration.html#cfn-mediaconnect-routeroutput-routeroutputconfiguration-mediaconnectflow
	//
	MediaConnectFlow interface{} `field:"optional" json:"mediaConnectFlow" yaml:"mediaConnectFlow"`
	// Configuration settings for connecting a router output to a MediaLive input.
	//
	// To connect the router output to the MediaLive input, provide both mediaLiveInputArn and mediaLivePipelineId. To disconnect it, call the update operation and omit both fields while keeping this configuration object. You can't change destinationTransitEncryption while the router output is connected to a MediaLive input. To change it, first disconnect the router output, then update the encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-routeroutputconfiguration.html#cfn-mediaconnect-routeroutput-routeroutputconfiguration-medialiveinput
	//
	MediaLiveInput interface{} `field:"optional" json:"mediaLiveInput" yaml:"mediaLiveInput"`
	// Configuration settings for a router output that is represented by a source endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-routeroutputconfiguration.html#cfn-mediaconnect-routeroutput-routeroutputconfiguration-sourceendpoint
	//
	SourceEndpoint interface{} `field:"optional" json:"sourceEndpoint" yaml:"sourceEndpoint"`
	// The configuration settings for a standard router output, including the protocol, protocol-specific configuration, network interface, and availability zone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-routeroutputconfiguration.html#cfn-mediaconnect-routeroutput-routeroutputconfiguration-standard
	//
	Standard interface{} `field:"optional" json:"standard" yaml:"standard"`
}

