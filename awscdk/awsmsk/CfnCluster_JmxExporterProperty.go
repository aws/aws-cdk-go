package awsmsk


// Indicates whether you want to enable or disable the JMX Exporter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jmxExporterProperty := &JmxExporterProperty{
//   	EnabledInBroker: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-jmxexporter.html
//
type CfnCluster_JmxExporterProperty struct {
	// Indicates whether you want to enable or disable the JMX Exporter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-jmxexporter.html#cfn-msk-cluster-jmxexporter-enabledinbroker
	//
	EnabledInBroker interface{} `field:"required" json:"enabledInBroker" yaml:"enabledInBroker"`
}

