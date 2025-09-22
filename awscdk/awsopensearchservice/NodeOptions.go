package awsopensearchservice


// Configuration for node options in OpenSearch domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nodeOptions := &NodeOptions{
//   	NodeConfig: &NodeConfig{
//   		Count: jsii.Number(123),
//   		Enabled: jsii.Boolean(false),
//   		Type: jsii.String("type"),
//   	},
//   	NodeType: awscdk.Aws_opensearchservice.NodeType_COORDINATOR,
//   }
//
type NodeOptions struct {
	// Configuration for the node type.
	NodeConfig *NodeConfig `field:"required" json:"nodeConfig" yaml:"nodeConfig"`
	// The type of node.
	//
	// Currently only 'coordinator' is supported.
	NodeType NodeType `field:"required" json:"nodeType" yaml:"nodeType"`
}

