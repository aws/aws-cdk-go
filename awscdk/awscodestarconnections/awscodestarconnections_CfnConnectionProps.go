package awscodestarconnections

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnConnection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectionProps := &cfnConnectionProps{
//   	connectionName: jsii.String("connectionName"),
//
//   	// the properties below are optional
//   	hostArn: jsii.String("hostArn"),
//   	providerType: jsii.String("providerType"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

