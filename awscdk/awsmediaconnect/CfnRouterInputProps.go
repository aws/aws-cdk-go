package awsmediaconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRouterInput`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var automatic interface{}
//   var default_ interface{}
//
//   cfnRouterInputProps := &CfnRouterInputProps{
//   	Configuration: &RouterInputConfigurationProperty{
//   		Failover: &FailoverRouterInputConfigurationProperty{
//   			NetworkInterfaceArn: jsii.String("networkInterfaceArn"),
//   			ProtocolConfigurations: []interface{}{
//   				&FailoverRouterInputProtocolConfigurationProperty{
//   					Rist: &RistRouterInputConfigurationProperty{
//   						Port: jsii.Number(123),
//   						RecoveryLatencyMilliseconds: jsii.Number(123),
//   					},
//   					Rtp: &RtpRouterInputConfigurationProperty{
//   						Port: jsii.Number(123),
//
//   						// the properties below are optional
//   						ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   					},
//   					SrtCaller: &SrtCallerRouterInputConfigurationProperty{
//   						MinimumLatencyMilliseconds: jsii.Number(123),
//   						SourceAddress: jsii.String("sourceAddress"),
//   						SourcePort: jsii.Number(123),
//
//   						// the properties below are optional
//   						DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   							EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   								RoleArn: jsii.String("roleArn"),
//   								SecretArn: jsii.String("secretArn"),
//   							},
//   						},
//   						StreamId: jsii.String("streamId"),
//   					},
//   					SrtListener: &SrtListenerRouterInputConfigurationProperty{
//   						MinimumLatencyMilliseconds: jsii.Number(123),
//   						Port: jsii.Number(123),
//
//   						// the properties below are optional
//   						DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   							EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   								RoleArn: jsii.String("roleArn"),
//   								SecretArn: jsii.String("secretArn"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			SourcePriorityMode: jsii.String("sourcePriorityMode"),
//
//   			// the properties below are optional
//   			PrimarySourceIndex: jsii.Number(123),
//   		},
//   		MediaConnectFlow: &MediaConnectFlowRouterInputConfigurationProperty{
//   			SourceTransitDecryption: &FlowTransitEncryptionProperty{
//   				EncryptionKeyConfiguration: &FlowTransitEncryptionKeyConfigurationProperty{
//   					Automatic: automatic,
//   					SecretsManager: &SecretsManagerEncryptionKeyConfigurationProperty{
//   						RoleArn: jsii.String("roleArn"),
//   						SecretArn: jsii.String("secretArn"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				EncryptionKeyType: jsii.String("encryptionKeyType"),
//   			},
//
//   			// the properties below are optional
//   			FlowArn: jsii.String("flowArn"),
//   			FlowOutputArn: jsii.String("flowOutputArn"),
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
//   						Port: jsii.Number(123),
//
//   						// the properties below are optional
//   						ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   					},
//   				},
//   			},
//   		},
//   		Standard: &StandardRouterInputConfigurationProperty{
//   			NetworkInterfaceArn: jsii.String("networkInterfaceArn"),
//   			ProtocolConfiguration: &RouterInputProtocolConfigurationProperty{
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
//
//   			// the properties below are optional
//   			Protocol: jsii.String("protocol"),
//   		},
//   	},
//   	MaximumBitrate: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	RoutingScope: jsii.String("routingScope"),
//   	Tier: jsii.String("tier"),
//
//   	// the properties below are optional
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	MaintenanceConfiguration: &MaintenanceConfigurationProperty{
//   		Default: default_,
//   		PreferredDayTime: &PreferredDayTimeMaintenanceConfigurationProperty{
//   			Day: jsii.String("day"),
//   			Time: jsii.String("time"),
//   		},
//   	},
//   	RegionName: jsii.String("regionName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TransitEncryption: &RouterInputTransitEncryptionProperty{
//   		EncryptionKeyConfiguration: &RouterInputTransitEncryptionKeyConfigurationProperty{
//   			Automatic: automatic,
//   			SecretsManager: &SecretsManagerEncryptionKeyConfigurationProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				SecretArn: jsii.String("secretArn"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		EncryptionKeyType: jsii.String("encryptionKeyType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routerinput.html
//
type CfnRouterInputProps struct {
	// The configuration settings for a router input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routerinput.html#cfn-mediaconnect-routerinput-configuration
	//
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
	// The maximum bitrate for the router input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routerinput.html#cfn-mediaconnect-routerinput-maximumbitrate
	//
	MaximumBitrate *float64 `field:"required" json:"maximumBitrate" yaml:"maximumBitrate"`
	// The name of the router input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routerinput.html#cfn-mediaconnect-routerinput-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routerinput.html#cfn-mediaconnect-routerinput-routingscope
	//
	RoutingScope *string `field:"required" json:"routingScope" yaml:"routingScope"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routerinput.html#cfn-mediaconnect-routerinput-tier
	//
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// The Availability Zone where you want to create the router input.
	//
	// This must be a valid Availability Zone for the region specified by regionName, or the current region if no regionName is provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routerinput.html#cfn-mediaconnect-routerinput-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The configuration settings for maintenance operations, including preferred maintenance windows and schedules.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routerinput.html#cfn-mediaconnect-routerinput-maintenanceconfiguration
	//
	MaintenanceConfiguration interface{} `field:"optional" json:"maintenanceConfiguration" yaml:"maintenanceConfiguration"`
	// The AWS Region for the router input.
	//
	// Defaults to the current region if not specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routerinput.html#cfn-mediaconnect-routerinput-regionname
	//
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
	// Key-value pairs that can be used to tag and organize this router input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routerinput.html#cfn-mediaconnect-routerinput-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Information about the encryption of the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routerinput.html#cfn-mediaconnect-routerinput-transitencryption
	//
	TransitEncryption interface{} `field:"optional" json:"transitEncryption" yaml:"transitEncryption"`
}

