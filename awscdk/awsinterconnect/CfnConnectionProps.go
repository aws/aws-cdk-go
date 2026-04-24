package awsinterconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnConnection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectionProps := &CfnConnectionProps{
//   	AttachPoint: &AttachPointProperty{
//   		Arn: jsii.String("arn"),
//   		DirectConnectGateway: jsii.String("directConnectGateway"),
//   	},
//
//   	// the properties below are optional
//   	ActivationKey: jsii.String("activationKey"),
//   	Bandwidth: jsii.String("bandwidth"),
//   	Description: jsii.String("description"),
//   	EnvironmentId: jsii.String("environmentId"),
//   	RemoteAccount: &RemoteAccountProperty{
//   		Identifier: jsii.String("identifier"),
//   	},
//   	RemoteOwnerAccount: jsii.String("remoteOwnerAccount"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-interconnect-connection.html
//
type CfnConnectionProps struct {
	// The logical attachment point in your AWS network where the managed connection will be connected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-interconnect-connection.html#cfn-interconnect-connection-attachpoint
	//
	AttachPoint interface{} `field:"required" json:"attachPoint" yaml:"attachPoint"`
	// The activation key for accepting a connection proposal from a partner CSP.
	//
	// Mutually exclusive with EnvironmentId.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-interconnect-connection.html#cfn-interconnect-connection-activationkey
	//
	ActivationKey *string `field:"optional" json:"activationKey" yaml:"activationKey"`
	// The bandwidth of the connection (e.g., 50Mbps, 1Gbps). Required when creating a connection through AWS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-interconnect-connection.html#cfn-interconnect-connection-bandwidth
	//
	Bandwidth *string `field:"optional" json:"bandwidth" yaml:"bandwidth"`
	// A description of the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-interconnect-connection.html#cfn-interconnect-connection-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the environment for the connection.
	//
	// Required when creating a connection through AWS. Mutually exclusive with ActivationKey.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-interconnect-connection.html#cfn-interconnect-connection-environmentid
	//
	EnvironmentId *string `field:"optional" json:"environmentId" yaml:"environmentId"`
	// The remote account identifier for the connection.
	//
	// Required when creating a connection through AWS. Replaces RemoteOwnerAccount.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-interconnect-connection.html#cfn-interconnect-connection-remoteaccount
	//
	RemoteAccount interface{} `field:"optional" json:"remoteAccount" yaml:"remoteAccount"`
	// Deprecated.
	//
	// Use RemoteAccount instead. The account ID of the remote owner. Required when creating a connection through AWS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-interconnect-connection.html#cfn-interconnect-connection-remoteowneraccount
	//
	RemoteOwnerAccount *string `field:"optional" json:"remoteOwnerAccount" yaml:"remoteOwnerAccount"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-interconnect-connection.html#cfn-interconnect-connection-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

