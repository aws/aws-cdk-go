package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-rolealias.html
//
type CfnRoleAliasProps struct {
	// The role ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-rolealias.html#cfn-iot-rolealias-rolearn
	//
	RoleArn interface{} `field:"required" json:"roleArn" yaml:"roleArn"`
	// The number of seconds for which the credential is valid.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-rolealias.html#cfn-iot-rolealias-credentialdurationseconds
	//
	// Default: - 3600.
	//
	CredentialDurationSeconds *float64 `field:"optional" json:"credentialDurationSeconds" yaml:"credentialDurationSeconds"`
	// The role alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-rolealias.html#cfn-iot-rolealias-rolealias
	//
	RoleAlias *string `field:"optional" json:"roleAlias" yaml:"roleAlias"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-rolealias.html#cfn-iot-rolealias-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

