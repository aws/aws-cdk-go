package previewawsappstreammixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnImageBuilderPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnImageBuilderMixinProps := &CfnImageBuilderMixinProps{
//   	AccessEndpoints: []interface{}{
//   		&AccessEndpointProperty{
//   			EndpointType: jsii.String("endpointType"),
//   			VpceId: jsii.String("vpceId"),
//   		},
//   	},
//   	AppstreamAgentVersion: jsii.String("appstreamAgentVersion"),
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	DomainJoinInfo: &DomainJoinInfoProperty{
//   		DirectoryName: jsii.String("directoryName"),
//   		OrganizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   	},
//   	EnableDefaultInternetAccess: jsii.Boolean(false),
//   	IamRoleArn: jsii.String("iamRoleArn"),
//   	ImageArn: jsii.String("imageArn"),
//   	ImageName: jsii.String("imageName"),
//   	InstanceType: jsii.String("instanceType"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcConfig: &VpcConfigProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html
//
type CfnImageBuilderMixinProps struct {
	// The list of virtual private cloud (VPC) interface endpoint objects.
	//
	// Administrators can connect to the image builder only through the specified endpoints.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-accessendpoints
	//
	AccessEndpoints interface{} `field:"optional" json:"accessEndpoints" yaml:"accessEndpoints"`
	// The version of the WorkSpaces Applications agent to use for this image builder.
	//
	// To use the latest version of the WorkSpaces Applications agent, specify [LATEST].
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-appstreamagentversion
	//
	AppstreamAgentVersion *string `field:"optional" json:"appstreamAgentVersion" yaml:"appstreamAgentVersion"`
	// The description to display.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The image builder name to display.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The name of the directory and organizational unit (OU) to use to join the image builder to a Microsoft Active Directory domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-domainjoininfo
	//
	DomainJoinInfo interface{} `field:"optional" json:"domainJoinInfo" yaml:"domainJoinInfo"`
	// Enables or disables default internet access for the image builder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-enabledefaultinternetaccess
	//
	EnableDefaultInternetAccess interface{} `field:"optional" json:"enableDefaultInternetAccess" yaml:"enableDefaultInternetAccess"`
	// The ARN of the IAM role that is applied to the image builder.
	//
	// To assume a role, the image builder calls the AWS Security Token Service `AssumeRole` API operation and passes the ARN of the role to use. The operation creates a new session with temporary credentials. AppStream 2.0 retrieves the temporary credentials and creates the *appstream_machine_role* credential profile on the instance.
	//
	// For more information, see [Using an IAM Role to Grant Permissions to Applications and Scripts Running on AppStream 2.0 Streaming Instances](https://docs.aws.amazon.com/appstream2/latest/developerguide/using-iam-roles-to-grant-permissions-to-applications-scripts-streaming-instances.html) in the *Amazon AppStream 2.0 Administration Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-iamrolearn
	//
	IamRoleArn *string `field:"optional" json:"iamRoleArn" yaml:"iamRoleArn"`
	// The ARN of the public, private, or shared image to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-imagearn
	//
	ImageArn *string `field:"optional" json:"imageArn" yaml:"imageArn"`
	// The name of the image used to create the image builder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-imagename
	//
	ImageName *string `field:"optional" json:"imageName" yaml:"imageName"`
	// The instance type to use when launching the image builder. The following instance types are available:.
	//
	// - stream.standard.small
	// - stream.standard.medium
	// - stream.standard.large
	// - stream.compute.large
	// - stream.compute.xlarge
	// - stream.compute.2xlarge
	// - stream.compute.4xlarge
	// - stream.compute.8xlarge
	// - stream.memory.large
	// - stream.memory.xlarge
	// - stream.memory.2xlarge
	// - stream.memory.4xlarge
	// - stream.memory.8xlarge
	// - stream.memory.z1d.large
	// - stream.memory.z1d.xlarge
	// - stream.memory.z1d.2xlarge
	// - stream.memory.z1d.3xlarge
	// - stream.memory.z1d.6xlarge
	// - stream.memory.z1d.12xlarge
	// - stream.graphics-design.large
	// - stream.graphics-design.xlarge
	// - stream.graphics-design.2xlarge
	// - stream.graphics-design.4xlarge
	// - stream.graphics.g4dn.xlarge
	// - stream.graphics.g4dn.2xlarge
	// - stream.graphics.g4dn.4xlarge
	// - stream.graphics.g4dn.8xlarge
	// - stream.graphics.g4dn.12xlarge
	// - stream.graphics.g4dn.16xlarge
	// - stream.graphics.g5.xlarge
	// - stream.graphics.g5.2xlarge
	// - stream.graphics.g5.4xlarge
	// - stream.graphics.g5.8xlarge
	// - stream.graphics.g5.16xlarge
	// - stream.graphics.g5.12xlarge
	// - stream.graphics.g5.24xlarge
	// - stream.graphics.g6.xlarge
	// - stream.graphics.g6.2xlarge
	// - stream.graphics.g6.4xlarge
	// - stream.graphics.g6.8xlarge
	// - stream.graphics.g6.16xlarge
	// - stream.graphics.g6.12xlarge
	// - stream.graphics.g6.24xlarge
	// - stream.graphics.gr6.4xlarge
	// - stream.graphics.gr6.8xlarge
	// - stream.graphics.g6f.large
	// - stream.graphics.g6f.xlarge
	// - stream.graphics.g6f.2xlarge
	// - stream.graphics.g6f.4xlarge
	// - stream.graphics.gr6f.4xlarge
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-instancetype
	//
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// A unique name for the image builder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The VPC configuration for the image builder.
	//
	// You can specify only one subnet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-vpcconfig
	//
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

