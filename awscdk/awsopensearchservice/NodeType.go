package awsopensearchservice


// NodeType is a string enum of the node types in OpenSearch domain.
//
// Example:
//   import opensearch "github.com/aws/aws-cdk-go/awscdk"
//
//
//   domain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
//   	Version: awscdk.EngineVersion_OPENSEARCH_1_3(),
//   	Capacity: &CapacityConfig{
//   		NodeOptions: []nodeOptions{
//   			&nodeOptions{
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
type NodeType string

const (
	// Coordinator node type.
	NodeType_COORDINATOR NodeType = "COORDINATOR"
)

