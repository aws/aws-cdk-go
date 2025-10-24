package awsemrserverless


// A configuration specification to be used when provisioning an application.
//
// A configuration consists of a classification, properties, and optional nested configurations. A classification refers to an application-specific configuration file. Properties are the settings you want to change in that file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationObjectProperty_ ConfigurationObjectProperty
//
//   configurationObjectProperty := &ConfigurationObjectProperty{
//   	Classification: jsii.String("classification"),
//
//   	// the properties below are optional
//   	Configurations: []interface{}{
//   		configurationObjectProperty_,
//   	},
//   	Properties: map[string]*string{
//   		"propertiesKey": jsii.String("properties"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-configurationobject.html
//
type CfnApplication_ConfigurationObjectProperty struct {
	// The classification within a configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-configurationobject.html#cfn-emrserverless-application-configurationobject-classification
	//
	Classification *string `field:"required" json:"classification" yaml:"classification"`
	// A list of additional configurations to apply within a configuration object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-configurationobject.html#cfn-emrserverless-application-configurationobject-configurations
	//
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
	// A set of properties specified within a configuration classification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-configurationobject.html#cfn-emrserverless-application-configurationobject-properties
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
}

