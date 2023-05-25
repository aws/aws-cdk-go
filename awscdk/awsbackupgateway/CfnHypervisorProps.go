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
	// The server host of the hypervisor.
	//
	// This can be either an IP address or a fully-qualified domain name (FQDN).
	Host *string `field:"optional" json:"host" yaml:"host"`
	// The Amazon Resource Name (ARN) of the AWS Key Management Service used to encrypt the hypervisor.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The Amazon Resource Name (ARN) of the group of gateways within the requested log.
	LogGroupArn *string `field:"optional" json:"logGroupArn" yaml:"logGroupArn"`
	// The name of the hypervisor.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The password for the hypervisor.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The tags of the hypervisor configuration to import.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The username for the hypervisor.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

