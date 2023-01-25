package awsappstream

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnImageBuilder`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnImageBuilderProps := &cfnImageBuilderProps{
//   	instanceType: jsii.String("instanceType"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	accessEndpoints: []interface{}{
//   		&accessEndpointProperty{
//   			endpointType: jsii.String("endpointType"),
//   			vpceId: jsii.String("vpceId"),
//   		},
//   	},
//   	appstreamAgentVersion: jsii.String("appstreamAgentVersion"),
//   	description: jsii.String("description"),
//   	displayName: jsii.String("displayName"),
//   	domainJoinInfo: &domainJoinInfoProperty{
//   		directoryName: jsii.String("directoryName"),
//   		organizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   	},
//   	enableDefaultInternetAccess: jsii.Boolean(false),
//   	iamRoleArn: jsii.String("iamRoleArn"),
//   	imageArn: jsii.String("imageArn"),
//   	imageName: jsii.String("imageName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vpcConfig: &vpcConfigProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
type CfnImageBuilderProps struct {
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
	// - stream.graphics-desktop.2xlarge
	// - stream.graphics.g4dn.xlarge
	// - stream.graphics.g4dn.2xlarge
	// - stream.graphics.g4dn.4xlarge
	// - stream.graphics.g4dn.8xlarge
	// - stream.graphics.g4dn.12xlarge
	// - stream.graphics.g4dn.16xlarge
	// - stream.graphics-pro.4xlarge
	// - stream.graphics-pro.8xlarge
	// - stream.graphics-pro.16xlarge
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// A unique name for the image builder.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of virtual private cloud (VPC) interface endpoint objects.
	//
	// Administrators can connect to the image builder only through the specified endpoints.
	AccessEndpoints interface{} `field:"optional" json:"accessEndpoints" yaml:"accessEndpoints"`
	// The version of the AppStream 2.0 agent to use for this image builder. To use the latest version of the AppStream 2.0 agent, specify [LATEST].
	AppstreamAgentVersion *string `field:"optional" json:"appstreamAgentVersion" yaml:"appstreamAgentVersion"`
	// The description to display.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The image builder name to display.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The name of the directory and organizational unit (OU) to use to join the image builder to a Microsoft Active Directory domain.
	DomainJoinInfo interface{} `field:"optional" json:"domainJoinInfo" yaml:"domainJoinInfo"`
	// Enables or disables default internet access for the image builder.
	EnableDefaultInternetAccess interface{} `field:"optional" json:"enableDefaultInternetAccess" yaml:"enableDefaultInternetAccess"`
	// The ARN of the IAM role that is applied to the image builder.
	//
	// To assume a role, the image builder calls the AWS Security Token Service `AssumeRole` API operation and passes the ARN of the role to use. The operation creates a new session with temporary credentials. AppStream 2.0 retrieves the temporary credentials and creates the *appstream_machine_role* credential profile on the instance.
	//
	// For more information, see [Using an IAM Role to Grant Permissions to Applications and Scripts Running on AppStream 2.0 Streaming Instances](https://docs.aws.amazon.com/appstream2/latest/developerguide/using-iam-roles-to-grant-permissions-to-applications-scripts-streaming-instances.html) in the *Amazon AppStream 2.0 Administration Guide* .
	IamRoleArn *string `field:"optional" json:"iamRoleArn" yaml:"iamRoleArn"`
	// The ARN of the public, private, or shared image to use.
	ImageArn *string `field:"optional" json:"imageArn" yaml:"imageArn"`
	// The name of the image used to create the image builder.
	ImageName *string `field:"optional" json:"imageName" yaml:"imageName"`
	// An array of key-value pairs.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The VPC configuration for the image builder.
	//
	// You can specify only one subnet.
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

