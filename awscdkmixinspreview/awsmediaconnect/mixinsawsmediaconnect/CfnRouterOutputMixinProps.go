package mixinsawsmediaconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnRouterOutputPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var automatic interface{}
//   var default_ interface{}
//
//   cfnRouterOutputMixinProps := &CfnRouterOutputMixinProps{
//   	AvailabilityZone: jsii.String("availabilityZone"),
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
//   				EncryptionKeyType: jsii.String("encryptionKeyType"),
//   			},
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
//   				EncryptionKeyType: jsii.String("encryptionKeyType"),
//   			},
//   			MediaLiveInputArn: jsii.String("mediaLiveInputArn"),
//   			MediaLivePipelineId: jsii.String("mediaLivePipelineId"),
//   		},
//   		Standard: &StandardRouterOutputConfigurationProperty{
//   			NetworkInterfaceArn: jsii.String("networkInterfaceArn"),
//   			Protocol: jsii.String("protocol"),
//   			ProtocolConfiguration: &RouterOutputProtocolConfigurationProperty{
//   				Rist: &RistRouterOutputConfigurationProperty{
//   					DestinationAddress: jsii.String("destinationAddress"),
//   					DestinationPort: jsii.Number(123),
//   				},
//   				Rtp: &RtpRouterOutputConfigurationProperty{
//   					DestinationAddress: jsii.String("destinationAddress"),
//   					DestinationPort: jsii.Number(123),
//   					ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   				},
//   				SrtCaller: &SrtCallerRouterOutputConfigurationProperty{
//   					DestinationAddress: jsii.String("destinationAddress"),
//   					DestinationPort: jsii.Number(123),
//   					EncryptionConfiguration: &SrtEncryptionConfigurationProperty{
//   						EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   							RoleArn: jsii.String("roleArn"),
//   							SecretArn: jsii.String("secretArn"),
//   						},
//   					},
//   					MinimumLatencyMilliseconds: jsii.Number(123),
//   					StreamId: jsii.String("streamId"),
//   				},
//   				SrtListener: &SrtListenerRouterOutputConfigurationProperty{
//   					EncryptionConfiguration: &SrtEncryptionConfigurationProperty{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routeroutput.html
//
type CfnRouterOutputMixinProps struct {
	// The Availability Zone where you want to create the router output.
	//
	// This must be a valid Availability Zone for the region specified by regionName, or the current region if no regionName is provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routeroutput.html#cfn-mediaconnect-routeroutput-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The configuration settings for a router output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routeroutput.html#cfn-mediaconnect-routeroutput-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The configuration settings for maintenance operations, including preferred maintenance windows and schedules.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routeroutput.html#cfn-mediaconnect-routeroutput-maintenanceconfiguration
	//
	MaintenanceConfiguration interface{} `field:"optional" json:"maintenanceConfiguration" yaml:"maintenanceConfiguration"`
	// The maximum bitrate for the router output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routeroutput.html#cfn-mediaconnect-routeroutput-maximumbitrate
	//
	MaximumBitrate *float64 `field:"optional" json:"maximumBitrate" yaml:"maximumBitrate"`
	// The name of the router output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routeroutput.html#cfn-mediaconnect-routeroutput-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The AWS Region for the router output.
	//
	// Defaults to the current region if not specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routeroutput.html#cfn-mediaconnect-routeroutput-regionname
	//
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routeroutput.html#cfn-mediaconnect-routeroutput-routingscope
	//
	RoutingScope *string `field:"optional" json:"routingScope" yaml:"routingScope"`
	// Key-value pairs that can be used to tag this router output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routeroutput.html#cfn-mediaconnect-routeroutput-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routeroutput.html#cfn-mediaconnect-routeroutput-tier
	//
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
}

