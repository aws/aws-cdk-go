package awsproton

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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
	// `AWS::Proton::EnvironmentAccountConnection.CodebuildRoleArn`.
	CodebuildRoleArn *string `field:"optional" json:"codebuildRoleArn" yaml:"codebuildRoleArn"`
	// `AWS::Proton::EnvironmentAccountConnection.ComponentRoleArn`.
	ComponentRoleArn *string `field:"optional" json:"componentRoleArn" yaml:"componentRoleArn"`
	// `AWS::Proton::EnvironmentAccountConnection.EnvironmentAccountId`.
	EnvironmentAccountId *string `field:"optional" json:"environmentAccountId" yaml:"environmentAccountId"`
	// `AWS::Proton::EnvironmentAccountConnection.EnvironmentName`.
	EnvironmentName *string `field:"optional" json:"environmentName" yaml:"environmentName"`
	// `AWS::Proton::EnvironmentAccountConnection.ManagementAccountId`.
	ManagementAccountId *string `field:"optional" json:"managementAccountId" yaml:"managementAccountId"`
	// `AWS::Proton::EnvironmentAccountConnection.RoleArn`.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// `AWS::Proton::EnvironmentAccountConnection.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

