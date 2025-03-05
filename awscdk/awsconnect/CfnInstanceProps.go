package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnInstance`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInstanceProps := &CfnInstanceProps{
//   	Attributes: &AttributesProperty{
//   		InboundCalls: jsii.Boolean(false),
//   		OutboundCalls: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		AutoResolveBestVoices: jsii.Boolean(false),
//   		ContactflowLogs: jsii.Boolean(false),
//   		ContactLens: jsii.Boolean(false),
//   		EarlyMedia: jsii.Boolean(false),
//   		UseCustomTtsVoices: jsii.Boolean(false),
//   	},
//   	IdentityManagementType: jsii.String("identityManagementType"),
//
//   	// the properties below are optional
//   	DirectoryId: jsii.String("directoryId"),
//   	InstanceAlias: jsii.String("instanceAlias"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instance.html
//
type CfnInstanceProps struct {
	// A toggle for an individual feature at the instance level.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instance.html#cfn-connect-instance-attributes
	//
	Attributes interface{} `field:"required" json:"attributes" yaml:"attributes"`
	// The identity management type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instance.html#cfn-connect-instance-identitymanagementtype
	//
	IdentityManagementType *string `field:"required" json:"identityManagementType" yaml:"identityManagementType"`
	// The identifier for the directory.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instance.html#cfn-connect-instance-directoryid
	//
	DirectoryId *string `field:"optional" json:"directoryId" yaml:"directoryId"`
	// The alias of instance.
	//
	// `InstanceAlias` is only required when `IdentityManagementType` is `CONNECT_MANAGED` or `SAML` . `InstanceAlias` is not required when `IdentityManagementType` is `EXISTING_DIRECTORY` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instance.html#cfn-connect-instance-instancealias
	//
	InstanceAlias *string `field:"optional" json:"instanceAlias" yaml:"instanceAlias"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instance.html#cfn-connect-instance-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

