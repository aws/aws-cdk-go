package awsemr


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var appConfigurationProperty_ AppConfigurationProperty
//
//   appConfigurationProperty := &AppConfigurationProperty{
//   	Classification: jsii.String("classification"),
//   	ConfigurationProperties: map[string]*string{
//   		"configurationPropertiesKey": jsii.String("configurationProperties"),
//   	},
//   	Configurations: []interface{}{
//   		appConfigurationProperty_,
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancegroupconfig-appconfiguration.html
//
type CfnInstanceGroupConfig_AppConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancegroupconfig-appconfiguration.html#cfn-emr-instancegroupconfig-appconfiguration-classification
	//
	Classification *string `field:"optional" json:"classification" yaml:"classification"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancegroupconfig-appconfiguration.html#cfn-emr-instancegroupconfig-appconfiguration-configurationproperties
	//
	ConfigurationProperties interface{} `field:"optional" json:"configurationProperties" yaml:"configurationProperties"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancegroupconfig-appconfiguration.html#cfn-emr-instancegroupconfig-appconfiguration-configurations
	//
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
}

