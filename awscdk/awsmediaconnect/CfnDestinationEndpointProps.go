package awsmediaconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDestinationEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDestinationEndpointProps := &CfnDestinationEndpointProps{
//   	Configuration: &DestinationEndpointConfigurationProperty{
//   		RouterInput: &RouterInputDestinationEndpointConfigurationProperty{
//   			RouterInputArn: jsii.String("routerInputArn"),
//   		},
//   		Standard: &StandardDestinationEndpointConfigurationProperty{
//   			DestinationNetworkingConfiguration: &StandardDestinationEndpointNetworkingConfigurationProperty{
//   				Public: &PublicDestinationEndpointNetworkingConfigurationProperty{
//   					DestinationAllowRules: []interface{}{
//   						&PublicDestinationEndpointNetworkingConfigurationAllowRuleProperty{
//   							Cidr: jsii.String("cidr"),
//   						},
//   					},
//   				},
//   				Vpc: &VpcDestinationEndpointNetworkingConfigurationProperty{
//   					SecurityGroupIds: []*string{
//   						jsii.String("securityGroupIds"),
//   					},
//   					SubnetId: jsii.String("subnetId"),
//   				},
//   			},
//   			Name: jsii.String("name"),
//   			ProtocolConfiguration: &StandardDestinationEndpointProtocolConfigurationProperty{
//   				Rist: &RistStandardDestinationEndpointConfigurationProperty{
//   					DestinationAddress: jsii.String("destinationAddress"),
//   					DestinationIngestPort: jsii.Number(123),
//   				},
//   				Rtp: &RtpStandardDestinationEndpointConfigurationProperty{
//   					DestinationAddress: jsii.String("destinationAddress"),
//   					DestinationIngestPort: jsii.Number(123),
//   					ForwardErrorCorrectionState: jsii.String("forwardErrorCorrectionState"),
//   				},
//   				SrtCaller: &SrtCallerStandardDestinationEndpointConfigurationProperty{
//   					MinimumLatencyMilliseconds: jsii.Number(123),
//   					SourceListenerPort: jsii.Number(123),
//
//   					// the properties below are optional
//   					EncryptionPassphrase: &EndpointEncryptionSecretConfigurationProperty{
//   						SecretArn: jsii.String("secretArn"),
//   					},
//   				},
//   				SrtListener: &SrtListenerStandardDestinationEndpointConfigurationProperty{
//   					DestinationAddress: jsii.String("destinationAddress"),
//   					DestinationIngestPort: jsii.Number(123),
//   					MinimumLatencyMilliseconds: jsii.Number(123),
//
//   					// the properties below are optional
//   					EncryptionPassphrase: &EndpointEncryptionSecretConfigurationProperty{
//   						SecretArn: jsii.String("secretArn"),
//   					},
//   					StreamId: jsii.String("streamId"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			DestinationNetworkingType: jsii.String("destinationNetworkingType"),
//   			Protocol: jsii.String("protocol"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	RegionName: jsii.String("regionName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-destinationendpoint.html
//
type CfnDestinationEndpointProps struct {
	// The configuration settings for a destination endpoint.
	//
	// This is a union type, so exactly one of its members is set, depending on the type of the destination endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-destinationendpoint.html#cfn-mediaconnect-destinationendpoint-configuration
	//
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
	// The description of the destination endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-destinationendpoint.html#cfn-mediaconnect-destinationendpoint-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the Amazon Web Services Region in which the destination endpoint is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-destinationendpoint.html#cfn-mediaconnect-destinationendpoint-regionname
	//
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
	// The tags associated with the destination endpoint.
	//
	// Each tag is a key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-destinationendpoint.html#cfn-mediaconnect-destinationendpoint-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

