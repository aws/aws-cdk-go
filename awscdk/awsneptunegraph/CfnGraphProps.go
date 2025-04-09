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
	// The provisioned memory-optimized Neptune Capacity Units (m-NCUs) to use for the graph.
	//
	// Min = 16.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptunegraph-graph.html#cfn-neptunegraph-graph-provisionedmemory
	//
	ProvisionedMemory *float64 `field:"required" json:"provisionedMemory" yaml:"provisionedMemory"`
	// A value that indicates whether the graph has deletion protection enabled.
	//
	// The graph can't be deleted when deletion protection is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptunegraph-graph.html#cfn-neptunegraph-graph-deletionprotection
	//
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The graph name. For example: `my-graph-1` .
	//
	// The name must contain from 1 to 63 letters, numbers, or hyphens, and its first character must be a letter. It cannot end with a hyphen or contain two consecutive hyphens.
	//
	// If you don't specify a graph name, a unique graph name is generated for you using the prefix `graph-for` , followed by a combination of `Stack Name` and a `UUID` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptunegraph-graph.html#cfn-neptunegraph-graph-graphname
	//
	GraphName *string `field:"optional" json:"graphName" yaml:"graphName"`
	// Specifies whether or not the graph can be reachable over the internet. All access to graphs is IAM authenticated.
	//
	// When the graph is publicly available, its domain name system (DNS) endpoint resolves to the public IP address from the internet. When the graph isn't publicly available, you need to create a `PrivateGraphEndpoint` in a given VPC to ensure the DNS name resolves to a private IP address that is reachable from the VPC.
	//
	// Default: If not specified, the default value is false.
	//
	// > If enabling public connectivity for the first time, there will be a delay while it is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptunegraph-graph.html#cfn-neptunegraph-graph-publicconnectivity
	//
	PublicConnectivity interface{} `field:"optional" json:"publicConnectivity" yaml:"publicConnectivity"`
	// The number of replicas in other AZs.
	//
	// Default: If not specified, the default value is 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptunegraph-graph.html#cfn-neptunegraph-graph-replicacount
	//
	ReplicaCount *float64 `field:"optional" json:"replicaCount" yaml:"replicaCount"`
	// Adds metadata tags to the new graph.
	//
	// These tags can also be used with cost allocation reporting, or used in a Condition statement in an IAM policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptunegraph-graph.html#cfn-neptunegraph-graph-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies the number of dimensions for vector embeddings that will be loaded into the graph.
	//
	// The value is specified as `dimension=` value. Max = 65,535
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptunegraph-graph.html#cfn-neptunegraph-graph-vectorsearchconfiguration
	//
	VectorSearchConfiguration interface{} `field:"optional" json:"vectorSearchConfiguration" yaml:"vectorSearchConfiguration"`
}

