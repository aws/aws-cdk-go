package awsgroundstation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDataflowEndpointGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataflowEndpointGroupProps := &CfnDataflowEndpointGroupProps{
//   	EndpointDetails: []interface{}{
//   		&EndpointDetailsProperty{
//   			Endpoint: &DataflowEndpointProperty{
//   				Address: &SocketAddressProperty{
//   					Name: jsii.String("name"),
//   					Port: jsii.Number(123),
//   				},
//   				Mtu: jsii.Number(123),
//   				Name: jsii.String("name"),
//   			},
//   			SecurityDetails: &SecurityDetailsProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				SecurityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				SubnetIds: []*string{
//   					jsii.String("subnetIds"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	ContactPostPassDurationSeconds: jsii.Number(123),
//   	ContactPrePassDurationSeconds: jsii.Number(123),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDataflowEndpointGroupProps struct {
	// List of Endpoint Details, containing address and port for each endpoint.
	EndpointDetails interface{} `field:"required" json:"endpointDetails" yaml:"endpointDetails"`
	// `AWS::GroundStation::DataflowEndpointGroup.ContactPostPassDurationSeconds`.
	ContactPostPassDurationSeconds *float64 `field:"optional" json:"contactPostPassDurationSeconds" yaml:"contactPostPassDurationSeconds"`
	// `AWS::GroundStation::DataflowEndpointGroup.ContactPrePassDurationSeconds`.
	ContactPrePassDurationSeconds *float64 `field:"optional" json:"contactPrePassDurationSeconds" yaml:"contactPrePassDurationSeconds"`
	// Tags assigned to a resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

