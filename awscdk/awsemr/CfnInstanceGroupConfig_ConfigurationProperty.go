package awsemr


// `Configurations` is a property of the `AWS::EMR::Cluster` resource that specifies the configuration of applications on an Amazon EMR cluster.
//
// Configurations are optional. You can use them to have EMR customize applications and software bundled with Amazon EMR when a cluster is created. A configuration consists of a classification, properties, and optional nested configurations. A classification refers to an application-specific configuration file. Properties are the settings you want to change in that file. For more information, see [Configuring Applications](https://docs.aws.amazon.com/emr/latest/ReleaseGuide/emr-configure-apps.html) .
//
// > Applies only to Amazon EMR releases 4.0 and later.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ configurationProperty
//
//   configurationProperty := &configurationProperty{
//   	Classification: jsii.String("classification"),
//   	ConfigurationProperties: map[string]*string{
//   		"configurationPropertiesKey": jsii.String("configurationProperties"),
//   	},
//   	Configurations: []interface{}{
//   		configurationProperty_,
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancegroupconfig-configuration.html
//
type CfnInstanceGroupConfig_ConfigurationProperty struct {
	// The classification within a configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancegroupconfig-configuration.html#cfn-emr-instancegroupconfig-configuration-classification
	//
	Classification *string `field:"optional" json:"classification" yaml:"classification"`
	// Within a configuration classification, a set of properties that represent the settings that you want to change in the configuration file.
	//
	// Duplicates not allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancegroupconfig-configuration.html#cfn-emr-instancegroupconfig-configuration-configurationproperties
	//
	ConfigurationProperties interface{} `field:"optional" json:"configurationProperties" yaml:"configurationProperties"`
	// A list of additional configurations to apply within a configuration object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancegroupconfig-configuration.html#cfn-emr-instancegroupconfig-configuration-configurations
	//
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
}

