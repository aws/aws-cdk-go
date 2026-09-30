package awsemrcontainers


// A configuration specification to be used when provisioning virtual clusters, which can include configurations for applications and software bundled with Amazon EMR on EKS.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var configurationProperty_ ConfigurationProperty
//
//   configurationProperty := &ConfigurationProperty{
//   	Classification: jsii.String("classification"),
//   	Configurations: []interface{}{
//   		configurationProperty_,
//   	},
//   	Properties: map[string]*string{
//   		"propertiesKey": jsii.String("properties"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-configuration.html
//
type CfnJobTemplatePropsMixin_ConfigurationProperty struct {
	// The classification within a configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-configuration.html#cfn-emrcontainers-jobtemplate-configuration-classification
	//
	Classification *string `field:"optional" json:"classification" yaml:"classification"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-configuration.html#cfn-emrcontainers-jobtemplate-configuration-configurations
	//
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
	// A set of properties specified within a configuration classification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-configuration.html#cfn-emrcontainers-jobtemplate-configuration-properties
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
}

