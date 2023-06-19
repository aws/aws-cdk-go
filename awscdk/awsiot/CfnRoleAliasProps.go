package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnRoleAlias`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRoleAliasProps := &CfnRoleAliasProps{
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	CredentialDurationSeconds: jsii.Number(123),
//   	RoleAlias: jsii.String("roleAlias"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRoleAliasProps struct {
	// The role ARN.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The number of seconds for which the credential is valid.
	CredentialDurationSeconds *float64 `field:"optional" json:"credentialDurationSeconds" yaml:"credentialDurationSeconds"`
	// The role alias.
	RoleAlias *string `field:"optional" json:"roleAlias" yaml:"roleAlias"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

