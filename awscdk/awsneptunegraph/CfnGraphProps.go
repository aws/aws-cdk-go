package awsneptunegraph

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnGraph`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGraphProps := &CfnGraphProps{
//   	ProvisionedMemory: jsii.Number(123),
//
//   	// the properties below are optional
//   	DeletionProtection: jsii.Boolean(false),
//   	GraphName: jsii.String("graphName"),
//   	PublicConnectivity: jsii.Boolean(false),
//   	ReplicaCount: jsii.Number(123),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VectorSearchConfiguration: &VectorSearchConfigurationProperty{
//   		VectorSearchDimension: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptunegraph-graph.html
//
type CfnGraphProps struct {
	// Memory for the Graph.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptunegraph-graph.html#cfn-neptunegraph-graph-provisionedmemory
	//
	ProvisionedMemory *float64 `field:"required" json:"provisionedMemory" yaml:"provisionedMemory"`
	// Value that indicates whether the Graph has deletion protection enabled.
	//
	// The graph can't be deleted when deletion protection is enabled.
	//
	// _Default_: If not specified, the default value is true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptunegraph-graph.html#cfn-neptunegraph-graph-deletionprotection
	//
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Contains a user-supplied name for the Graph.
	//
	// If you don't specify a name, we generate a unique Graph Name using a combination of Stack Name and a UUID comprising of 4 characters.
	//
	// _Important_: If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptunegraph-graph.html#cfn-neptunegraph-graph-graphname
	//
	GraphName *string `field:"optional" json:"graphName" yaml:"graphName"`
	// Specifies whether the Graph can be reached over the internet. Access to all graphs requires IAM authentication.
	//
	// When the Graph is publicly reachable, its Domain Name System (DNS) endpoint resolves to the public IP address from the internet.
	//
	// When the Graph isn't publicly reachable, you need to create a PrivateGraphEndpoint in a given VPC to ensure the DNS name resolves to a private IP address that is reachable from the VPC.
	//
	// _Default_: If not specified, the default value is false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptunegraph-graph.html#cfn-neptunegraph-graph-publicconnectivity
	//
	PublicConnectivity interface{} `field:"optional" json:"publicConnectivity" yaml:"publicConnectivity"`
	// Specifies the number of replicas you want when finished. All replicas will be provisioned in different availability zones.
	//
	// Replica Count should always be less than or equal to 2.
	//
	// _Default_: If not specified, the default value is 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptunegraph-graph.html#cfn-neptunegraph-graph-replicacount
	//
	ReplicaCount *float64 `field:"optional" json:"replicaCount" yaml:"replicaCount"`
	// The tags associated with this graph.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptunegraph-graph.html#cfn-neptunegraph-graph-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The vector search configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptunegraph-graph.html#cfn-neptunegraph-graph-vectorsearchconfiguration
	//
	VectorSearchConfiguration interface{} `field:"optional" json:"vectorSearchConfiguration" yaml:"vectorSearchConfiguration"`
}

