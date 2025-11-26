package awsmediaconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRouterOutput`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var automatic interface{}
//   var default_ interface{}
//
//   cfnRouterOutputProps := &CfnRouterOutputProps{
//   	Configuration: &RouterOutputConfigurationProperty{
//   		MediaConnectFlow: &MediaConnectFlowRouterOutputConfigurationProperty{
//   			DestinationTransitEncryption: &FlowTransitEncryptionProperty{
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
//   			FlowSourceArn: jsii.String("flowSourceArn"),
//   		},
//   		MediaLiveInput: &MediaLiveInputRouterOutputConfigurationProperty{
//   			DestinationTransitEncryption: &MediaLiveTransitEncryptionProperty{
//   				EncryptionKeyConfiguration: &MediaLiveTransitEncryptionKeyConfigurationProperty{
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
//   			MediaLiveInputArn: jsii.String("mediaLiveInputArn"),
//   			MediaLivePipelineId: jsii.String("mediaLivePipelineId"),
//   		},
//   		Standard: &StandardRouterOutputConfigurationProperty{
//   			NetworkInterfaceArn: jsii.String("networkInterfaceArn"),
//   			ProtocolConfiguration: &RouterOutputProtocolConfigurationProperty{
//   				Rist: &RistRouterOutputConfigurationProperty{
//   					DestinationAddress: jsii.String("destinationAddress"),
//   					DestinationPort: jsii.Number(123),
//   				},
//   				Rtp: &RtpRouterOutputConfigurationProperty{
//   					DestinationAddress: jsii.String("destinationAddress"),
//   					DestinationPort: jsii.Number(123),
//
//   					// the properties below are optional
//   					ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   				},
//   				SrtCaller: &SrtCallerRouterOutputConfigurationProperty{
//   					DestinationAddress: jsii.String("destinationAddress"),
//   					DestinationPort: jsii.Number(123),
//   					MinimumLatencyMilliseconds: jsii.Number(123),
//
//   					// the properties below are optional
//   					EncryptionConfiguration: &SrtEncryptionConfigurationProperty{
//   						EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   							RoleArn: jsii.String("roleArn"),
//   							SecretArn: jsii.String("secretArn"),
//   						},
//   					},
//   					StreamId: jsii.String("streamId"),
//   				},
//   				SrtListener: &SrtListenerRouterOutputConfigurationProperty{
//   					MinimumLatencyMilliseconds: jsii.Number(123),
//   					Port: jsii.Number(123),
//
//   					// the properties below are optional
//   					EncryptionConfiguration: &SrtEncryptionConfigurationProperty{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routeroutput.html
//
type CfnRouterOutputProps struct {
	// The configuration settings for a router output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routeroutput.html#cfn-mediaconnect-routeroutput-configuration
	//
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
	// The maximum bitrate for the router output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routeroutput.html#cfn-mediaconnect-routeroutput-maximumbitrate
	//
	MaximumBitrate *float64 `field:"required" json:"maximumBitrate" yaml:"maximumBitrate"`
	// The name of the router output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routeroutput.html#cfn-mediaconnect-routeroutput-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Indicates whether the router output is configured for Regional or global routing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routeroutput.html#cfn-mediaconnect-routeroutput-routingscope
	//
	RoutingScope *string `field:"required" json:"routingScope" yaml:"routingScope"`
	// The tier level of the router output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routeroutput.html#cfn-mediaconnect-routeroutput-tier
	//
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// The Availability Zone of the router output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routeroutput.html#cfn-mediaconnect-routeroutput-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The maintenance configuration settings applied to this router output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routeroutput.html#cfn-mediaconnect-routeroutput-maintenanceconfiguration
	//
	MaintenanceConfiguration interface{} `field:"optional" json:"maintenanceConfiguration" yaml:"maintenanceConfiguration"`
	// The AWS Region where the router output is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routeroutput.html#cfn-mediaconnect-routeroutput-regionname
	//
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
	// Key-value pairs that can be used to tag and organize this router output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routeroutput.html#cfn-mediaconnect-routeroutput-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

