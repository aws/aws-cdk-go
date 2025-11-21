package awsmediaconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var automatic interface{}
//
//   routerInputConfigurationProperty := &RouterInputConfigurationProperty{
//   	Failover: &FailoverRouterInputConfigurationProperty{
//   		NetworkInterfaceArn: jsii.String("networkInterfaceArn"),
//   		ProtocolConfigurations: []interface{}{
//   			&FailoverRouterInputProtocolConfigurationProperty{
//   				Rist: &RistRouterInputConfigurationProperty{
//   					Port: jsii.Number(123),
//   					RecoveryLatencyMilliseconds: jsii.Number(123),
//   				},
//   				Rtp: &RtpRouterInputConfigurationProperty{
//   					Port: jsii.Number(123),
//
//   					// the properties below are optional
//   					ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   				},
//   				SrtCaller: &SrtCallerRouterInputConfigurationProperty{
//   					MinimumLatencyMilliseconds: jsii.Number(123),
//   					SourceAddress: jsii.String("sourceAddress"),
//   					SourcePort: jsii.Number(123),
//
//   					// the properties below are optional
//   					DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   						EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   							RoleArn: jsii.String("roleArn"),
//   							SecretArn: jsii.String("secretArn"),
//   						},
//   					},
//   					StreamId: jsii.String("streamId"),
//   				},
//   				SrtListener: &SrtListenerRouterInputConfigurationProperty{
//   					MinimumLatencyMilliseconds: jsii.Number(123),
//   					Port: jsii.Number(123),
//
//   					// the properties below are optional
//   					DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   						EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   							RoleArn: jsii.String("roleArn"),
//   							SecretArn: jsii.String("secretArn"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		SourcePriorityMode: jsii.String("sourcePriorityMode"),
//
//   		// the properties below are optional
//   		PrimarySourceIndex: jsii.Number(123),
//   	},
//   	MediaConnectFlow: &MediaConnectFlowRouterInputConfigurationProperty{
//   		SourceTransitDecryption: &FlowTransitEncryptionProperty{
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
//   		FlowOutputArn: jsii.String("flowOutputArn"),
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
//   					Port: jsii.Number(123),
//
//   					// the properties below are optional
//   					ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   				},
//   			},
//   		},
//   	},
//   	Standard: &StandardRouterInputConfigurationProperty{
//   		NetworkInterfaceArn: jsii.String("networkInterfaceArn"),
//   		ProtocolConfiguration: &RouterInputProtocolConfigurationProperty{
//   			Rist: &RistRouterInputConfigurationProperty{
//   				Port: jsii.Number(123),
//   				RecoveryLatencyMilliseconds: jsii.Number(123),
//   			},
//   			Rtp: &RtpRouterInputConfigurationProperty{
//   				Port: jsii.Number(123),
//
//   				// the properties below are optional
//   				ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   			},
//   			SrtCaller: &SrtCallerRouterInputConfigurationProperty{
//   				MinimumLatencyMilliseconds: jsii.Number(123),
//   				SourceAddress: jsii.String("sourceAddress"),
//   				SourcePort: jsii.Number(123),
//
//   				// the properties below are optional
//   				DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   					EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   						RoleArn: jsii.String("roleArn"),
//   						SecretArn: jsii.String("secretArn"),
//   					},
//   				},
//   				StreamId: jsii.String("streamId"),
//   			},
//   			SrtListener: &SrtListenerRouterInputConfigurationProperty{
//   				MinimumLatencyMilliseconds: jsii.Number(123),
//   				Port: jsii.Number(123),
//
//   				// the properties below are optional
//   				DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputconfiguration.html
//
type CfnRouterInput_RouterInputConfigurationProperty struct {
	// Configuration settings for a failover router input that allows switching between two input sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputconfiguration.html#cfn-mediaconnect-routerinput-routerinputconfiguration-failover
	//
	Failover interface{} `field:"optional" json:"failover" yaml:"failover"`
	// Configuration settings for connecting a router input to a flow output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputconfiguration.html#cfn-mediaconnect-routerinput-routerinputconfiguration-mediaconnectflow
	//
	MediaConnectFlow interface{} `field:"optional" json:"mediaConnectFlow" yaml:"mediaConnectFlow"`
	// Configuration settings for a merge router input that combines two input sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputconfiguration.html#cfn-mediaconnect-routerinput-routerinputconfiguration-merge
	//
	Merge interface{} `field:"optional" json:"merge" yaml:"merge"`
	// The configuration settings for a standard router input, including the protocol, protocol-specific configuration, network interface, and availability zone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputconfiguration.html#cfn-mediaconnect-routerinput-routerinputconfiguration-standard
	//
	Standard interface{} `field:"optional" json:"standard" yaml:"standard"`
}

