package awsdatazone


// A configuration of the connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   connectionConfigurationProperty := &ConnectionConfigurationProperty{
//   	Classification: jsii.String("classification"),
//   	Properties: map[string]*string{
//   		"propertiesKey": jsii.String("properties"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-connectionconfiguration.html
//
type CfnConnectionPropsMixin_ConnectionConfigurationProperty struct {
	// The classification of the connection configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-connectionconfiguration.html#cfn-datazone-connection-connectionconfiguration-classification
	//
	Classification *string `field:"optional" json:"classification" yaml:"classification"`
	// Property Map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-connectionconfiguration.html#cfn-datazone-connection-connectionconfiguration-properties
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
}

