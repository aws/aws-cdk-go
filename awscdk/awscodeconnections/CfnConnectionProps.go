package awscodeconnections

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
//   	ConnectionName: jsii.String("connectionName"),
//
//   	// the properties below are optional
//   	HostArn: jsii.String("hostArn"),
//   	ProviderType: jsii.String("providerType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeconnections-connection.html
//
type CfnConnectionProps struct {
	// The name of the connection.
	//
	// Connection names must be unique in an AWS account .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeconnections-connection.html#cfn-codeconnections-connection-connectionname
	//
	ConnectionName *string `field:"required" json:"connectionName" yaml:"connectionName"`
	// The Amazon Resource Name (ARN) of the host associated with the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeconnections-connection.html#cfn-codeconnections-connection-hostarn
	//
	HostArn *string `field:"optional" json:"hostArn" yaml:"hostArn"`
	// The name of the external provider where your third-party code repository is configured.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeconnections-connection.html#cfn-codeconnections-connection-providertype
	//
	ProviderType *string `field:"optional" json:"providerType" yaml:"providerType"`
	// Specifies the tags applied to a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeconnections-connection.html#cfn-codeconnections-connection-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

