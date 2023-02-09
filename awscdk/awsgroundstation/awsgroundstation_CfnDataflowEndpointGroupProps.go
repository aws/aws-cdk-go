package awsgroundstation

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDataflowEndpointGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataflowEndpointGroupProps := &cfnDataflowEndpointGroupProps{
//   	endpointDetails: []interface{}{
//   		&endpointDetailsProperty{
//   			endpoint: &dataflowEndpointProperty{
//   				address: &socketAddressProperty{
//   					name: jsii.String("name"),
//   					port: jsii.Number(123),
//   				},
//   				mtu: jsii.Number(123),
//   				name: jsii.String("name"),
//   			},
//   			securityDetails: &securityDetailsProperty{
//   				roleArn: jsii.String("roleArn"),
//   				securityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				subnetIds: []*string{
//   					jsii.String("subnetIds"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	contactPostPassDurationSeconds: jsii.Number(123),
//   	contactPrePassDurationSeconds: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

