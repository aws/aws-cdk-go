package awsmediaconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSourceEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSourceEndpointProps := &CfnSourceEndpointProps{
//   	Configuration: &SourceEndpointConfigurationProperty{
//   		RouterOutput: &RouterOutputSourceEndpointConfigurationProperty{
//   			RouterOutputArn: jsii.String("routerOutputArn"),
//   		},
//   		Standard: &StandardSourceEndpointConfigurationProperty{
//   			Name: jsii.String("name"),
//   			ProtocolConfiguration: &StandardSourceEndpointProtocolConfigurationProperty{
//   				Rist: &RistStandardSourceEndpointConfigurationProperty{
//   					DestinationIngestPort: jsii.Number(123),
//   					RecoveryLatencyMilliseconds: jsii.Number(123),
//   				},
//   				Rtp: &RtpStandardSourceEndpointConfigurationProperty{
//   					DestinationIngestPort: jsii.Number(123),
//   					ForwardErrorCorrectionState: jsii.String("forwardErrorCorrectionState"),
//   				},
//   				SrtCaller: &SrtCallerStandardSourceEndpointConfigurationProperty{
//   					DestinationIngestPort: jsii.Number(123),
//   					MinimumLatencyMilliseconds: jsii.Number(123),
//
//   					// the properties below are optional
//   					EncryptionPassphrase: &EndpointEncryptionSecretConfigurationProperty{
//   						SecretArn: jsii.String("secretArn"),
//   					},
//   				},
//   				SrtListener: &SrtListenerStandardSourceEndpointConfigurationProperty{
//   					MinimumLatencyMilliseconds: jsii.Number(123),
//   					SourceAddress: jsii.String("sourceAddress"),
//   					SourceListenerPort: jsii.Number(123),
//
//   					// the properties below are optional
//   					EncryptionPassphrase: &EndpointEncryptionSecretConfigurationProperty{
//   						SecretArn: jsii.String("secretArn"),
//   					},
//   					StreamId: jsii.String("streamId"),
//   				},
//   			},
//   			SourceNetworkingConfiguration: &StandardSourceEndpointNetworkingConfigurationProperty{
//   				Public: &PublicSourceEndpointNetworkingConfigurationProperty{
//   					SourceAllowRules: []interface{}{
//   						&PublicSourceEndpointNetworkingConfigurationAllowRuleProperty{
//   							Cidr: jsii.String("cidr"),
//   						},
//   					},
//   				},
//   				Vpc: &VpcSourceEndpointNetworkingConfigurationProperty{
//   					SecurityGroupIds: []*string{
//   						jsii.String("securityGroupIds"),
//   					},
//   					SubnetId: jsii.String("subnetId"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			Protocol: jsii.String("protocol"),
//   			SourceNetworkingType: jsii.String("sourceNetworkingType"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-sourceendpoint.html
//
type CfnSourceEndpointProps struct {
	// The configuration settings for a source endpoint.
	//
	// This is a union type, so exactly one of its members is set, depending on the type of the source endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-sourceendpoint.html#cfn-mediaconnect-sourceendpoint-configuration
	//
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
	// The description of the source endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-sourceendpoint.html#cfn-mediaconnect-sourceendpoint-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the Amazon Web Services Region in which the source endpoint is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-sourceendpoint.html#cfn-mediaconnect-sourceendpoint-regionname
	//
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
	// The tags associated with the source endpoint.
	//
	// Each tag is a key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-sourceendpoint.html#cfn-mediaconnect-sourceendpoint-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

