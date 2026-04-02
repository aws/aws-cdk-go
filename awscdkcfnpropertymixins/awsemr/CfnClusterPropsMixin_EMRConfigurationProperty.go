package awsemr


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var eMRConfigurationProperty_ EMRConfigurationProperty
//
//   eMRConfigurationProperty := &EMRConfigurationProperty{
//   	Classification: jsii.String("classification"),
//   	ConfigurationProperties: map[string]*string{
//   		"configurationPropertiesKey": jsii.String("configurationProperties"),
//   	},
//   	Configurations: []interface{}{
//   		eMRConfigurationProperty_,
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-emrconfiguration.html
//
type CfnClusterPropsMixin_EMRConfigurationProperty struct {
	// The classification within a configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-emrconfiguration.html#cfn-emr-cluster-emrconfiguration-classification
	//
	Classification *string `field:"optional" json:"classification" yaml:"classification"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-emrconfiguration.html#cfn-emr-cluster-emrconfiguration-configurationproperties
	//
	ConfigurationProperties interface{} `field:"optional" json:"configurationProperties" yaml:"configurationProperties"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-emrconfiguration.html#cfn-emr-cluster-emrconfiguration-configurations
	//
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
}

