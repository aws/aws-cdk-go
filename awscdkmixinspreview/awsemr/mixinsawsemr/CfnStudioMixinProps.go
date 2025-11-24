package mixinsawsemr

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnStudioPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStudioMixinProps := &CfnStudioMixinProps{
//   	AuthMode: jsii.String("authMode"),
//   	DefaultS3Location: jsii.String("defaultS3Location"),
//   	Description: jsii.String("description"),
//   	EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   	EngineSecurityGroupId: jsii.String("engineSecurityGroupId"),
//   	IdcInstanceArn: jsii.String("idcInstanceArn"),
//   	IdcUserAssignment: jsii.String("idcUserAssignment"),
//   	IdpAuthUrl: jsii.String("idpAuthUrl"),
//   	IdpRelayStateParameterName: jsii.String("idpRelayStateParameterName"),
//   	Name: jsii.String("name"),
//   	ServiceRole: jsii.String("serviceRole"),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrustedIdentityPropagationEnabled: jsii.Boolean(false),
//   	UserRole: jsii.String("userRole"),
//   	VpcId: jsii.String("vpcId"),
//   	WorkspaceSecurityGroupId: jsii.String("workspaceSecurityGroupId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-studio.html
//
type CfnStudioMixinProps struct {
	// Specifies whether the Studio authenticates users using IAM Identity Center or IAM.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-studio.html#cfn-emr-studio-authmode
	//
	AuthMode *string `field:"optional" json:"authMode" yaml:"authMode"`
	// The Amazon S3 location to back up EMR Studio Workspaces and notebook files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-studio.html#cfn-emr-studio-defaults3location
	//
	DefaultS3Location *string `field:"optional" json:"defaultS3Location" yaml:"defaultS3Location"`
	// A detailed description of the Amazon EMR Studio.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-studio.html#cfn-emr-studio-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The AWS  key identifier (ARN) used to encrypt Amazon EMR Studio workspace and notebook files when backed up to Amazon S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-studio.html#cfn-emr-studio-encryptionkeyarn
	//
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// The ID of the Amazon EMR Studio Engine security group.
	//
	// The Engine security group allows inbound network traffic from the Workspace security group, and it must be in the same VPC specified by `VpcId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-studio.html#cfn-emr-studio-enginesecuritygroupid
	//
	EngineSecurityGroupId *string `field:"optional" json:"engineSecurityGroupId" yaml:"engineSecurityGroupId"`
	// The ARN of the IAM Identity Center instance the Studio application belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-studio.html#cfn-emr-studio-idcinstancearn
	//
	IdcInstanceArn *string `field:"optional" json:"idcInstanceArn" yaml:"idcInstanceArn"`
	// Indicates whether the Studio has `REQUIRED` or `OPTIONAL` IAM Identity Center user assignment.
	//
	// If the value is set to `REQUIRED` , users must be explicitly assigned to the Studio application to access the Studio.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-studio.html#cfn-emr-studio-idcuserassignment
	//
	IdcUserAssignment *string `field:"optional" json:"idcUserAssignment" yaml:"idcUserAssignment"`
	// Your identity provider's authentication endpoint.
	//
	// Amazon EMR Studio redirects federated users to this endpoint for authentication when logging in to a Studio with the Studio URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-studio.html#cfn-emr-studio-idpauthurl
	//
	IdpAuthUrl *string `field:"optional" json:"idpAuthUrl" yaml:"idpAuthUrl"`
	// The name of your identity provider's `RelayState` parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-studio.html#cfn-emr-studio-idprelaystateparametername
	//
	IdpRelayStateParameterName *string `field:"optional" json:"idpRelayStateParameterName" yaml:"idpRelayStateParameterName"`
	// A descriptive name for the Amazon EMR Studio.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-studio.html#cfn-emr-studio-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the IAM role that will be assumed by the Amazon EMR Studio.
	//
	// The service role provides a way for Amazon EMR Studio to interoperate with other AWS services.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-studio.html#cfn-emr-studio-servicerole
	//
	ServiceRole *string `field:"optional" json:"serviceRole" yaml:"serviceRole"`
	// A list of subnet IDs to associate with the Amazon EMR Studio.
	//
	// A Studio can have a maximum of 5 subnets. The subnets must belong to the VPC specified by `VpcId` . Studio users can create a Workspace in any of the specified subnets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-studio.html#cfn-emr-studio-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-studio.html#cfn-emr-studio-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Indicates whether the Studio has Trusted identity propagation enabled.
	//
	// The default value is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-studio.html#cfn-emr-studio-trustedidentitypropagationenabled
	//
	TrustedIdentityPropagationEnabled interface{} `field:"optional" json:"trustedIdentityPropagationEnabled" yaml:"trustedIdentityPropagationEnabled"`
	// The Amazon Resource Name (ARN) of the IAM user role that will be assumed by users and groups logged in to a Studio.
	//
	// The permissions attached to this IAM role can be scoped down for each user or group using session policies. You only need to specify `UserRole` when you set `AuthMode` to `SSO` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-studio.html#cfn-emr-studio-userrole
	//
	UserRole *string `field:"optional" json:"userRole" yaml:"userRole"`
	// The ID of the Amazon Virtual Private Cloud (Amazon VPC) to associate with the Studio.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-studio.html#cfn-emr-studio-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
	// The ID of the Workspace security group associated with the Amazon EMR Studio.
	//
	// The Workspace security group allows outbound network traffic to resources in the Engine security group and to the internet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-studio.html#cfn-emr-studio-workspacesecuritygroupid
	//
	WorkspaceSecurityGroupId *string `field:"optional" json:"workspaceSecurityGroupId" yaml:"workspaceSecurityGroupId"`
}

