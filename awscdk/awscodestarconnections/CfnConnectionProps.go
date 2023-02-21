package awscodestarconnections

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
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnConnectionProps struct {
	// The name of the connection.
	//
	// Connection names must be unique in an AWS user account.
	ConnectionName *string `field:"required" json:"connectionName" yaml:"connectionName"`
	// The Amazon Resource Name (ARN) of the host associated with the connection.
	HostArn *string `field:"optional" json:"hostArn" yaml:"hostArn"`
	// The name of the external provider where your third-party code repository is configured.
	ProviderType *string `field:"optional" json:"providerType" yaml:"providerType"`
	// Specifies the tags applied to the resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

