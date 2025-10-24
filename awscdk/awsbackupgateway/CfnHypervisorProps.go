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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Username: jsii.String("username"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backupgateway-hypervisor.html
//
type CfnHypervisorProps struct {
	// The server host of the hypervisor.
	//
	// This can be either an IP address or a fully-qualified domain name (FQDN).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backupgateway-hypervisor.html#cfn-backupgateway-hypervisor-host
	//
	Host *string `field:"optional" json:"host" yaml:"host"`
	// The Amazon Resource Name (ARN) of the AWS Key Management Service used to encrypt the hypervisor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backupgateway-hypervisor.html#cfn-backupgateway-hypervisor-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The Amazon Resource Name (ARN) of the group of gateways within the requested log.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backupgateway-hypervisor.html#cfn-backupgateway-hypervisor-loggrouparn
	//
	LogGroupArn *string `field:"optional" json:"logGroupArn" yaml:"logGroupArn"`
	// The name of the hypervisor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backupgateway-hypervisor.html#cfn-backupgateway-hypervisor-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The password for the hypervisor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backupgateway-hypervisor.html#cfn-backupgateway-hypervisor-password
	//
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The tags of the hypervisor configuration to import.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backupgateway-hypervisor.html#cfn-backupgateway-hypervisor-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The username for the hypervisor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backupgateway-hypervisor.html#cfn-backupgateway-hypervisor-username
	//
	Username *string `field:"optional" json:"username" yaml:"username"`
}

