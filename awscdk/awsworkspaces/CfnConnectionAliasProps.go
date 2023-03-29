package awsworkspaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnConnectionAlias`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectionAliasProps := &CfnConnectionAliasProps{
//   	ConnectionString: jsii.String("connectionString"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnConnectionAliasProps struct {
	// The connection string specified for the connection alias.
	//
	// The connection string must be in the form of a fully qualified domain name (FQDN), such as `www.example.com` .
	ConnectionString *string `field:"required" json:"connectionString" yaml:"connectionString"`
	// The tags to associate with the connection alias.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

