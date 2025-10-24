package awsopensearchservice


// Configuration for a specific node type in OpenSearch domain.
//
// Example:
//   import opensearch "github.com/aws/aws-cdk-go/awscdk"
//
//
//   domain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
//   	Version: awscdk.EngineVersion_OPENSEARCH_1_3(),
//   	Capacity: &CapacityConfig{
//   		NodeOptions: []NodeOptions{
//   			&NodeOptions{
//   				NodeType: opensearch.NodeType_COORDINATOR,
//   				NodeConfig: &NodeConfig{
//   					Enabled: jsii.Boolean(true),
//   					Count: jsii.Number(2),
//   					Type: jsii.String("m5.large.search"),
//   				},
//   			},
//   		},
//   	},
//   })
//
type NodeConfig struct {
	// The number of nodes of this type.
	// Default: - 1.
	//
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Whether this node type is enabled.
	// Default: - false.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The instance type for the nodes.
	// Default: - m5.large.search
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

