package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties that describe an existing instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var instanceEngine iInstanceEngine
//   var securityGroup securityGroup
//
//   databaseInstanceAttributes := &DatabaseInstanceAttributes{
//   	InstanceEndpointAddress: jsii.String("instanceEndpointAddress"),
//   	InstanceIdentifier: jsii.String("instanceIdentifier"),
//   	Port: jsii.Number(123),
//   	SecurityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//
//   	// the properties below are optional
//   	Engine: instanceEngine,
//   	InstanceResourceId: jsii.String("instanceResourceId"),
//   }
//
type DatabaseInstanceAttributes struct {
	// The endpoint address.
	InstanceEndpointAddress *string `field:"required" json:"instanceEndpointAddress" yaml:"instanceEndpointAddress"`
	// The instance identifier.
	InstanceIdentifier *string `field:"required" json:"instanceIdentifier" yaml:"instanceIdentifier"`
	// The database port.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The security groups of the instance.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// The engine of the existing database Instance.
	// Default: - the imported Instance's engine is unknown.
	//
	Engine IInstanceEngine `field:"optional" json:"engine" yaml:"engine"`
	// The AWS Region-unique, immutable identifier for the DB instance.
	//
	// This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB instance is accessed.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#aws-resource-rds-dbinstance-return-values
	//
	InstanceResourceId *string `field:"optional" json:"instanceResourceId" yaml:"instanceResourceId"`
}

