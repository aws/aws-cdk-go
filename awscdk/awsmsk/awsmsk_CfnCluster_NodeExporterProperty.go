package awsmsk


// Indicates whether you want to enable or disable the Node Exporter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nodeExporterProperty := &nodeExporterProperty{
//   	enabledInBroker: jsii.Boolean(false),
//   }
//
type CfnCluster_NodeExporterProperty struct {
	// Indicates whether you want to enable or disable the Node Exporter.
	EnabledInBroker interface{} `field:"required" json:"enabledInBroker" yaml:"enabledInBroker"`
}

