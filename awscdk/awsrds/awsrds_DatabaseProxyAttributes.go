package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties that describe an existing DB Proxy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup securityGroup
//
//   databaseProxyAttributes := &databaseProxyAttributes{
//   	dbProxyArn: jsii.String("dbProxyArn"),
//   	dbProxyName: jsii.String("dbProxyName"),
//   	endpoint: jsii.String("endpoint"),
//   	securityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   }
//
type DatabaseProxyAttributes struct {
	// DB Proxy ARN.
	DbProxyArn *string `field:"required" json:"dbProxyArn" yaml:"dbProxyArn"`
	// DB Proxy Name.
	DbProxyName *string `field:"required" json:"dbProxyName" yaml:"dbProxyName"`
	// Endpoint.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// The security groups of the instance.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"required" json:"securityGroups" yaml:"securityGroups"`
}

