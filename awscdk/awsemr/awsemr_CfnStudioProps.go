package awsemr

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnStudio`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStudioProps := &cfnStudioProps{
//   	authMode: jsii.String("authMode"),
//   	defaultS3Location: jsii.String("defaultS3Location"),
//   	engineSecurityGroupId: jsii.String("engineSecurityGroupId"),
//   	name: jsii.String("name"),
//   	serviceRole: jsii.String("serviceRole"),
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	vpcId: jsii.String("vpcId"),
//   	workspaceSecurityGroupId: jsii.String("workspaceSecurityGroupId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	idpAuthUrl: jsii.String("idpAuthUrl"),
//   	idpRelayStateParameterName: jsii.String("idpRelayStateParameterName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	userRole: jsii.String("userRole"),
//   }
//
type CfnStudioProps struct {
	// Specifies whether the Studio authenticates users using AWS SSO or IAM.
	AuthMode *string `field:"required" json:"authMode" yaml:"authMode"`
	// The Amazon S3 location to back up EMR Studio Workspaces and notebook files.
	DefaultS3Location *string `field:"required" json:"defaultS3Location" yaml:"defaultS3Location"`
	// The ID of the Amazon EMR Studio Engine security group.
	//
	// The Engine security group allows inbound network traffic from the Workspace security group, and it must be in the same VPC specified by `VpcId` .
	EngineSecurityGroupId *string `field:"required" json:"engineSecurityGroupId" yaml:"engineSecurityGroupId"`
	// A descriptive name for the Amazon EMR Studio.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the IAM role that will be assumed by the Amazon EMR Studio.
	//
	// The service role provides a way for Amazon EMR Studio to interoperate with other AWS services.
	ServiceRole *string `field:"required" json:"serviceRole" yaml:"serviceRole"`
	// A list of subnet IDs to associate with the Amazon EMR Studio.
	//
	// A Studio can have a maximum of 5 subnets. The subnets must belong to the VPC specified by `VpcId` . Studio users can create a Workspace in any of the specified subnets.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The ID of the Amazon Virtual Private Cloud (Amazon VPC) to associate with the Studio.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The ID of the Workspace security group associated with the Amazon EMR Studio.
	//
	// The Workspace security group allows outbound network traffic to resources in the Engine security group and to the internet.
	WorkspaceSecurityGroupId *string `field:"required" json:"workspaceSecurityGroupId" yaml:"workspaceSecurityGroupId"`
	// A detailed description of the Amazon EMR Studio.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Your identity provider's authentication endpoint.
	//
	// Amazon EMR Studio redirects federated users to this endpoint for authentication when logging in to a Studio with the Studio URL.
	IdpAuthUrl *string `field:"optional" json:"idpAuthUrl" yaml:"idpAuthUrl"`
	// The name of your identity provider's `RelayState` parameter.
	IdpRelayStateParameterName *string `field:"optional" json:"idpRelayStateParameterName" yaml:"idpRelayStateParameterName"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The Amazon Resource Name (ARN) of the IAM user role that will be assumed by users and groups logged in to a Studio.
	//
	// The permissions attached to this IAM role can be scoped down for each user or group using session policies. You only need to specify `UserRole` when you set `AuthMode` to `SSO` .
	UserRole *string `field:"optional" json:"userRole" yaml:"userRole"`
}

