package awsmediaconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var automatic interface{}
//
//   routerInputConfigurationProperty := &RouterInputConfigurationProperty{
//   	DestinationEndpoint: &DestinationEndpointRouterInputConfigurationProperty{
//   		EndpointLifecycleManagement: jsii.String("endpointLifecycleManagement"),
//   		SourceTransitDecryption: &EndpointEncryptionConfigurationProperty{
//   			Automatic: automatic,
//   			SecretsManager: &SecretsManagerEndpointEncryptionConfigurationProperty{
//   				EncryptionSecret: &EndpointEncryptionSecretConfigurationProperty{
//   					SecretArn: jsii.String("secretArn"),
//   				},
//   				EncryptionSecretType: jsii.String("encryptionSecretType"),
//   			},
//   		},
//   		SourceTransitDecryptionType: jsii.String("sourceTransitDecryptionType"),
//   	},
//   	Failover: &FailoverRouterInputConfigurationProperty{
//   		NetworkInterfaceArn: jsii.String("networkInterfaceArn"),
//   		PrimarySourceIndex: jsii.Number(123),
//   		ProtocolConfigurations: []interface{}{
//   			&FailoverRouterInputProtocolConfigurationProperty{
//   				Rist: &RistRouterInputConfigurationProperty{
//   					Port: jsii.Number(123),
//   					RecoveryLatencyMilliseconds: jsii.Number(123),
//   				},
//   				RtmpPush: &RtmpPushRouterInputConfigurationProperty{
//   					NamingConfiguration: &RtmpPushNamingConfigurationProperty{
//   						Automatic: automatic,
//   						Manual: &ManualRtmpPushNamingConfigurationProperty{
//   							ApplicationName: jsii.String("applicationName"),
//   							StreamName: jsii.String("streamName"),
//   						},
//   					},
//   					NamingType: jsii.String("namingType"),
//   					Port: jsii.Number(123),
//   				},
//   				Rtp: &RtpRouterInputConfigurationProperty{
//   					ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   					Port: jsii.Number(123),
//   				},
//   				SrtCaller: &SrtCallerRouterInputConfigurationProperty{
//   					DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   						EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   							RoleArn: jsii.String("roleArn"),
//   							SecretArn: jsii.String("secretArn"),
//   						},
//   					},
//   					MinimumLatencyMilliseconds: jsii.Number(123),
//   					SourceAddress: jsii.String("sourceAddress"),
//   					SourcePort: jsii.Number(123),
//   					StreamId: jsii.String("streamId"),
//   				},
//   				SrtListener: &SrtListenerRouterInputConfigurationProperty{
//   					DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   						EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   							RoleArn: jsii.String("roleArn"),
//   							SecretArn: jsii.String("secretArn"),
//   						},
//   					},
//   					MinimumLatencyMilliseconds: jsii.Number(123),
//   					Port: jsii.Number(123),
//   				},
//   			},
//   		},
//   		SourcePriorityMode: jsii.String("sourcePriorityMode"),
//   	},
//   	MediaConnectFlow: &MediaConnectFlowRouterInputConfigurationProperty{
//   		FlowArn: jsii.String("flowArn"),
//   		FlowOutputArn: jsii.String("flowOutputArn"),
//   		SourceTransitDecryption: &FlowTransitEncryptionProperty{
//   			EncryptionKeyConfiguration: &FlowTransitEncryptionKeyConfigurationProperty{
//   				Automatic: automatic,
//   				SecretsManager: &SecretsManagerEncryptionKeyConfigurationProperty{
//   					RoleArn: jsii.String("roleArn"),
//   					SecretArn: jsii.String("secretArn"),
//   				},
//   			},
//   			EncryptionKeyType: jsii.String("encryptionKeyType"),
//   		},
//   	},
//   	MediaLiveChannel: &MediaLiveChannelRouterInputConfigurationProperty{
//   		MediaLiveChannelArn: jsii.String("mediaLiveChannelArn"),
//   		MediaLiveChannelOutputName: jsii.String("mediaLiveChannelOutputName"),
//   		MediaLivePipelineId: jsii.String("mediaLivePipelineId"),
//   		SourceTransitDecryption: &MediaLiveTransitEncryptionProperty{
//   			EncryptionKeyConfiguration: &MediaLiveTransitEncryptionKeyConfigurationProperty{
//   				Automatic: automatic,
//   				SecretsManager: &SecretsManagerEncryptionKeyConfigurationProperty{
//   					RoleArn: jsii.String("roleArn"),
//   					SecretArn: jsii.String("secretArn"),
//   				},
//   			},
//   			EncryptionKeyType: jsii.String("encryptionKeyType"),
//   		},
//   	},
//   	Merge: &MergeRouterInputConfigurationProperty{
//   		MergeRecoveryWindowMilliseconds: jsii.Number(123),
//   		NetworkInterfaceArn: jsii.String("networkInterfaceArn"),
//   		ProtocolConfigurations: []interface{}{
//   			&MergeRouterInputProtocolConfigurationProperty{
//   				Rist: &RistRouterInputConfigurationProperty{
//   					Port: jsii.Number(123),
//   					RecoveryLatencyMilliseconds: jsii.Number(123),
//   				},
//   				Rtp: &RtpRouterInputConfigurationProperty{
//   					ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   					Port: jsii.Number(123),
//   				},
//   				SrtCaller: &SrtCallerRouterInputConfigurationProperty{
//   					DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   						EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   							RoleArn: jsii.String("roleArn"),
//   							SecretArn: jsii.String("secretArn"),
//   						},
//   					},
//   					MinimumLatencyMilliseconds: jsii.Number(123),
//   					SourceAddress: jsii.String("sourceAddress"),
//   					SourcePort: jsii.Number(123),
//   					StreamId: jsii.String("streamId"),
//   				},
//   				SrtListener: &SrtListenerRouterInputConfigurationProperty{
//   					DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   						EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   							RoleArn: jsii.String("roleArn"),
//   							SecretArn: jsii.String("secretArn"),
//   						},
//   					},
//   					MinimumLatencyMilliseconds: jsii.Number(123),
//   					Port: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	Standard: &StandardRouterInputConfigurationProperty{
//   		NetworkInterfaceArn: jsii.String("networkInterfaceArn"),
//   		Protocol: jsii.String("protocol"),
//   		ProtocolConfiguration: &RouterInputProtocolConfigurationProperty{
//   			Rist: &RistRouterInputConfigurationProperty{
//   				Port: jsii.Number(123),
//   				RecoveryLatencyMilliseconds: jsii.Number(123),
//   			},
//   			RtmpPush: &RtmpPushRouterInputConfigurationProperty{
//   				NamingConfiguration: &RtmpPushNamingConfigurationProperty{
//   					Automatic: automatic,
//   					Manual: &ManualRtmpPushNamingConfigurationProperty{
//   						ApplicationName: jsii.String("applicationName"),
//   						StreamName: jsii.String("streamName"),
//   					},
//   				},
//   				NamingType: jsii.String("namingType"),
//   				Port: jsii.Number(123),
//   			},
//   			Rtp: &RtpRouterInputConfigurationProperty{
//   				ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   				Port: jsii.Number(123),
//   			},
//   			SrtCaller: &SrtCallerRouterInputConfigurationProperty{
//   				DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   					EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   						RoleArn: jsii.String("roleArn"),
//   						SecretArn: jsii.String("secretArn"),
//   					},
//   				},
//   				MinimumLatencyMilliseconds: jsii.Number(123),
//   				SourceAddress: jsii.String("sourceAddress"),
//   				SourcePort: jsii.Number(123),
//   				StreamId: jsii.String("streamId"),
//   			},
//   			SrtListener: &SrtListenerRouterInputConfigurationProperty{
//   				DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputconfiguration.html
//
type CfnRouterInputPropsMixin_RouterInputConfigurationProperty struct {
	// Configuration settings for a router input that is represented by a destination endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputconfiguration.html#cfn-mediaconnect-routerinput-routerinputconfiguration-destinationendpoint
	//
	DestinationEndpoint interface{} `field:"optional" json:"destinationEndpoint" yaml:"destinationEndpoint"`
	// Configuration settings for a failover router input that allows switching between two input sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputconfiguration.html#cfn-mediaconnect-routerinput-routerinputconfiguration-failover
	//
	Failover interface{} `field:"optional" json:"failover" yaml:"failover"`
	// Configuration settings for connecting a router input to a flow output.
	//
	// To connect the router input to the flow output, provide both flowArn and flowOutputArn. To disconnect it, call the update operation and omit both fields while keeping this configuration object. You can't change sourceTransitDecryption while the router input is connected to a flow output. To change it, first disconnect the router input, then update the decryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputconfiguration.html#cfn-mediaconnect-routerinput-routerinputconfiguration-mediaconnectflow
	//
	MediaConnectFlow interface{} `field:"optional" json:"mediaConnectFlow" yaml:"mediaConnectFlow"`
	// Configuration settings for connecting a router input to a MediaLive channel output.
	//
	// To connect the router input to the MediaLive channel output, provide all three of mediaLiveChannelArn, mediaLivePipelineId, and mediaLiveChannelOutputName. To disconnect it, call the update operation and omit all three fields while keeping this configuration object. You can't change sourceTransitDecryption while the router input is connected to a MediaLive channel output. To change it, first disconnect the router input, then update the decryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputconfiguration.html#cfn-mediaconnect-routerinput-routerinputconfiguration-medialivechannel
	//
	MediaLiveChannel interface{} `field:"optional" json:"mediaLiveChannel" yaml:"mediaLiveChannel"`
	// Configuration settings for a merge router input that combines two input sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputconfiguration.html#cfn-mediaconnect-routerinput-routerinputconfiguration-merge
	//
	Merge interface{} `field:"optional" json:"merge" yaml:"merge"`
	// The configuration settings for a standard router input, including the protocol, protocol-specific configuration, network interface, and availability zone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputconfiguration.html#cfn-mediaconnect-routerinput-routerinputconfiguration-standard
	//
	Standard interface{} `field:"optional" json:"standard" yaml:"standard"`
}

