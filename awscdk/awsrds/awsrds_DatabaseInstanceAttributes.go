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
//   databaseInstanceAttributes := &databaseInstanceAttributes{
//   	instanceEndpointAddress: jsii.String("instanceEndpointAddress"),
//   	instanceIdentifier: jsii.String("instanceIdentifier"),
//   	port: jsii.Number(123),
//   	securityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//
//   	// the properties below are optional
//   	engine: instanceEngine,
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
	Engine IInstanceEngine `field:"optional" json:"engine" yaml:"engine"`
}

