package awsproton

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnEnvironmentAccountConnection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEnvironmentAccountConnectionProps := &CfnEnvironmentAccountConnectionProps{
//   	CodebuildRoleArn: jsii.String("codebuildRoleArn"),
//   	ComponentRoleArn: jsii.String("componentRoleArn"),
//   	EnvironmentAccountId: jsii.String("environmentAccountId"),
//   	EnvironmentName: jsii.String("environmentName"),
//   	ManagementAccountId: jsii.String("managementAccountId"),
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnEnvironmentAccountConnectionProps struct {
	// The Amazon Resource Name (ARN) of an IAM service role in the environment account.
	//
	// AWS Proton uses this role to provision infrastructure resources using CodeBuild-based provisioning in the associated environment account.
	CodebuildRoleArn *string `field:"optional" json:"codebuildRoleArn" yaml:"codebuildRoleArn"`
	// The Amazon Resource Name (ARN) of the IAM service role that AWS Proton uses when provisioning directly defined components in the associated environment account.
	//
	// It determines the scope of infrastructure that a component can provision in the account.
	//
	// The environment account connection must have a `componentRoleArn` to allow directly defined components to be associated with any environments running in the account.
	//
	// For more information about components, see [AWS Proton components](https://docs.aws.amazon.com/proton/latest/userguide/ag-components.html) in the *AWS Proton User Guide* .
	ComponentRoleArn *string `field:"optional" json:"componentRoleArn" yaml:"componentRoleArn"`
	// The environment account that's connected to the environment account connection.
	EnvironmentAccountId *string `field:"optional" json:"environmentAccountId" yaml:"environmentAccountId"`
	// The name of the environment that's associated with the environment account connection.
	EnvironmentName *string `field:"optional" json:"environmentName" yaml:"environmentName"`
	// The ID of the management account that's connected to the environment account connection.
	ManagementAccountId *string `field:"optional" json:"managementAccountId" yaml:"managementAccountId"`
	// The IAM service role that's associated with the environment account connection.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// An optional list of metadata items that you can associate with the AWS Proton environment account connection.
	//
	// A tag is a key-value pair.
	//
	// For more information, see [AWS Proton resources and tagging](https://docs.aws.amazon.com/proton/latest/userguide/resources.html) in the *AWS Proton User Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

