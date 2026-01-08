package awsgroundstation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDataflowEndpointGroupV2`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataflowEndpointGroupV2Props := &CfnDataflowEndpointGroupV2Props{
//   	ContactPostPassDurationSeconds: jsii.Number(123),
//   	ContactPrePassDurationSeconds: jsii.Number(123),
//   	Endpoints: []interface{}{
//   		&CreateEndpointDetailsProperty{
//   			DownlinkAwsGroundStationAgentEndpoint: &DownlinkAwsGroundStationAgentEndpointProperty{
//   				DataflowDetails: &DownlinkDataflowDetailsProperty{
//   					AgentConnectionDetails: &DownlinkConnectionDetailsProperty{
//   						AgentIpAndPortAddress: &RangedConnectionDetailsProperty{
//   							SocketAddress: &RangedSocketAddressProperty{
//   								Name: jsii.String("name"),
//   								PortRange: &IntegerRangeProperty{
//   									Maximum: jsii.Number(123),
//   									Minimum: jsii.Number(123),
//   								},
//   							},
//
//   							// the properties below are optional
//   							Mtu: jsii.Number(123),
//   						},
//   						EgressAddressAndPort: &ConnectionDetailsProperty{
//   							SocketAddress: &SocketAddressProperty{
//   								Name: jsii.String("name"),
//   								Port: jsii.Number(123),
//   							},
//
//   							// the properties below are optional
//   							Mtu: jsii.Number(123),
//   						},
//   					},
//   				},
//   				Name: jsii.String("name"),
//   			},
//   			UplinkAwsGroundStationAgentEndpoint: &UplinkAwsGroundStationAgentEndpointProperty{
//   				DataflowDetails: &UplinkDataflowDetailsProperty{
//   					AgentConnectionDetails: &UplinkConnectionDetailsProperty{
//   						AgentIpAndPortAddress: &RangedConnectionDetailsProperty{
//   							SocketAddress: &RangedSocketAddressProperty{
//   								Name: jsii.String("name"),
//   								PortRange: &IntegerRangeProperty{
//   									Maximum: jsii.Number(123),
//   									Minimum: jsii.Number(123),
//   								},
//   							},
//
//   							// the properties below are optional
//   							Mtu: jsii.Number(123),
//   						},
//   						IngressAddressAndPort: &ConnectionDetailsProperty{
//   							SocketAddress: &SocketAddressProperty{
//   								Name: jsii.String("name"),
//   								Port: jsii.Number(123),
//   							},
//
//   							// the properties below are optional
//   							Mtu: jsii.Number(123),
//   						},
//   					},
//   				},
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-dataflowendpointgroupv2.html
//
type CfnDataflowEndpointGroupV2Props struct {
	// Amount of time, in seconds, after a contact ends that the Ground Station Dataflow Endpoint Group will be in a POSTPASS state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-dataflowendpointgroupv2.html#cfn-groundstation-dataflowendpointgroupv2-contactpostpassdurationseconds
	//
	ContactPostPassDurationSeconds *float64 `field:"optional" json:"contactPostPassDurationSeconds" yaml:"contactPostPassDurationSeconds"`
	// Amount of time, in seconds, before a contact starts that the Ground Station Dataflow Endpoint Group will be in a PREPASS state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-dataflowendpointgroupv2.html#cfn-groundstation-dataflowendpointgroupv2-contactprepassdurationseconds
	//
	ContactPrePassDurationSeconds *float64 `field:"optional" json:"contactPrePassDurationSeconds" yaml:"contactPrePassDurationSeconds"`
	// List of endpoints for the dataflow endpoint group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-dataflowendpointgroupv2.html#cfn-groundstation-dataflowendpointgroupv2-endpoints
	//
	Endpoints interface{} `field:"optional" json:"endpoints" yaml:"endpoints"`
	// Tags assigned to a resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-dataflowendpointgroupv2.html#cfn-groundstation-dataflowendpointgroupv2-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

