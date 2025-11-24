package mixinsawsmediaconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnRouterInputPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var automatic interface{}
//   var default_ interface{}
//
//   cfnRouterInputMixinProps := &CfnRouterInputMixinProps{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	Configuration: &RouterInputConfigurationProperty{
//   		Failover: &FailoverRouterInputConfigurationProperty{
//   			NetworkInterfaceArn: jsii.String("networkInterfaceArn"),
//   			PrimarySourceIndex: jsii.Number(123),
//   			ProtocolConfigurations: []interface{}{
//   				&FailoverRouterInputProtocolConfigurationProperty{
//   					Rist: &RistRouterInputConfigurationProperty{
//   						Port: jsii.Number(123),
//   						RecoveryLatencyMilliseconds: jsii.Number(123),
//   					},
//   					Rtp: &RtpRouterInputConfigurationProperty{
//   						ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   						Port: jsii.Number(123),
//   					},
//   					SrtCaller: &SrtCallerRouterInputConfigurationProperty{
//   						DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   							EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   								RoleArn: jsii.String("roleArn"),
//   								SecretArn: jsii.String("secretArn"),
//   							},
//   						},
//   						MinimumLatencyMilliseconds: jsii.Number(123),
//   						SourceAddress: jsii.String("sourceAddress"),
//   						SourcePort: jsii.Number(123),
//   						StreamId: jsii.String("streamId"),
//   					},
//   					SrtListener: &SrtListenerRouterInputConfigurationProperty{
//   						DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   							EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   								RoleArn: jsii.String("roleArn"),
//   								SecretArn: jsii.String("secretArn"),
//   							},
//   						},
//   						MinimumLatencyMilliseconds: jsii.Number(123),
//   						Port: jsii.Number(123),
//   					},
//   				},
//   			},
//   			SourcePriorityMode: jsii.String("sourcePriorityMode"),
//   		},
//   		MediaConnectFlow: &MediaConnectFlowRouterInputConfigurationProperty{
//   			FlowArn: jsii.String("flowArn"),
//   			FlowOutputArn: jsii.String("flowOutputArn"),
//   			SourceTransitDecryption: &FlowTransitEncryptionProperty{
//   				EncryptionKeyConfiguration: &FlowTransitEncryptionKeyConfigurationProperty{
//   					Automatic: automatic,
//   					SecretsManager: &SecretsManagerEncryptionKeyConfigurationProperty{
//   						RoleArn: jsii.String("roleArn"),
//   						SecretArn: jsii.String("secretArn"),
//   					},
//   				},
//   				EncryptionKeyType: jsii.String("encryptionKeyType"),
//   			},
//   		},
//   		Merge: &MergeRouterInputConfigurationProperty{
//   			MergeRecoveryWindowMilliseconds: jsii.Number(123),
//   			NetworkInterfaceArn: jsii.String("networkInterfaceArn"),
//   			ProtocolConfigurations: []interface{}{
//   				&MergeRouterInputProtocolConfigurationProperty{
//   					Rist: &RistRouterInputConfigurationProperty{
//   						Port: jsii.Number(123),
//   						RecoveryLatencyMilliseconds: jsii.Number(123),
//   					},
//   					Rtp: &RtpRouterInputConfigurationProperty{
//   						ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   						Port: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   		Standard: &StandardRouterInputConfigurationProperty{
//   			NetworkInterfaceArn: jsii.String("networkInterfaceArn"),
//   			Protocol: jsii.String("protocol"),
//   			ProtocolConfiguration: &RouterInputProtocolConfigurationProperty{
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
//   	MaintenanceConfiguration: &MaintenanceConfigurationProperty{
//   		Default: default_,
//   		PreferredDayTime: &PreferredDayTimeMaintenanceConfigurationProperty{
//   			Day: jsii.String("day"),
//   			Time: jsii.String("time"),
//   		},
//   	},
//   	MaximumBitrate: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	RegionName: jsii.String("regionName"),
//   	RoutingScope: jsii.String("routingScope"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Tier: jsii.String("tier"),
//   	TransitEncryption: &RouterInputTransitEncryptionProperty{
//   		EncryptionKeyConfiguration: &RouterInputTransitEncryptionKeyConfigurationProperty{
//   			Automatic: automatic,
//   			SecretsManager: &SecretsManagerEncryptionKeyConfigurationProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				SecretArn: jsii.String("secretArn"),
//   			},
//   		},
//   		EncryptionKeyType: jsii.String("encryptionKeyType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routerinput.html
//
type CfnRouterInputMixinProps struct {
	// The Availability Zone where you want to create the router input.
	//
	// This must be a valid Availability Zone for the region specified by regionName, or the current region if no regionName is provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routerinput.html#cfn-mediaconnect-routerinput-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The configuration settings for a router input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routerinput.html#cfn-mediaconnect-routerinput-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The configuration settings for maintenance operations, including preferred maintenance windows and schedules.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routerinput.html#cfn-mediaconnect-routerinput-maintenanceconfiguration
	//
	MaintenanceConfiguration interface{} `field:"optional" json:"maintenanceConfiguration" yaml:"maintenanceConfiguration"`
	// The maximum bitrate for the router input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routerinput.html#cfn-mediaconnect-routerinput-maximumbitrate
	//
	MaximumBitrate *float64 `field:"optional" json:"maximumBitrate" yaml:"maximumBitrate"`
	// The name of the router input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routerinput.html#cfn-mediaconnect-routerinput-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The AWS Region for the router input.
	//
	// Defaults to the current region if not specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routerinput.html#cfn-mediaconnect-routerinput-regionname
	//
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routerinput.html#cfn-mediaconnect-routerinput-routingscope
	//
	RoutingScope *string `field:"optional" json:"routingScope" yaml:"routingScope"`
	// Key-value pairs that can be used to tag and organize this router input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routerinput.html#cfn-mediaconnect-routerinput-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routerinput.html#cfn-mediaconnect-routerinput-tier
	//
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
	// Information about the encryption of the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routerinput.html#cfn-mediaconnect-routerinput-transitencryption
	//
	TransitEncryption interface{} `field:"optional" json:"transitEncryption" yaml:"transitEncryption"`
}

