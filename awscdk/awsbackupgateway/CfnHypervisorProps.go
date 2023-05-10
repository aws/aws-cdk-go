package awsbackupgateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnHypervisor`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnHypervisorProps := &CfnHypervisorProps{
//   	Host: jsii.String("host"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	LogGroupArn: jsii.String("logGroupArn"),
//   	Name: jsii.String("name"),
//   	Password: jsii.String("password"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Username: jsii.String("username"),
//   }
//
type CfnHypervisorProps struct {
	// `AWS::BackupGateway::Hypervisor.Host`.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// `AWS::BackupGateway::Hypervisor.KmsKeyArn`.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// `AWS::BackupGateway::Hypervisor.LogGroupArn`.
	LogGroupArn *string `field:"optional" json:"logGroupArn" yaml:"logGroupArn"`
	// `AWS::BackupGateway::Hypervisor.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::BackupGateway::Hypervisor.Password`.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// `AWS::BackupGateway::Hypervisor.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::BackupGateway::Hypervisor.Username`.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

