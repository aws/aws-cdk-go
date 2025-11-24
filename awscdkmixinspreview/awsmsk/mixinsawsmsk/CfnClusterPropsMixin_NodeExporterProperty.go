package mixinsawsmsk


// Indicates whether you want to enable or disable the Node Exporter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   nodeExporterProperty := &NodeExporterProperty{
//   	EnabledInBroker: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-nodeexporter.html
//
type CfnClusterPropsMixin_NodeExporterProperty struct {
	// Indicates whether you want to enable or disable the Node Exporter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-nodeexporter.html#cfn-msk-cluster-nodeexporter-enabledinbroker
	//
	EnabledInBroker interface{} `field:"optional" json:"enabledInBroker" yaml:"enabledInBroker"`
}

