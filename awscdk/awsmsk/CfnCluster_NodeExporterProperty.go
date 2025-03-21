package awsmsk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nodeExporterProperty := &NodeExporterProperty{
//   	EnabledInBroker: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-nodeexporter.html
//
type CfnCluster_NodeExporterProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-nodeexporter.html#cfn-msk-cluster-nodeexporter-enabledinbroker
	//
	EnabledInBroker interface{} `field:"required" json:"enabledInBroker" yaml:"enabledInBroker"`
}

