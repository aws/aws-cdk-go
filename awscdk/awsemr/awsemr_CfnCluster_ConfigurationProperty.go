package awsemr


// > Used only with Amazon EMR release 4.0 and later.
//
// `Configuration` is a subproperty of `InstanceFleetConfig` or `InstanceGroupConfig` . `Configuration` specifies optional configurations for customizing open-source big data applications and environment parameters. A configuration consists of a classification, properties, and optional nested configurations. A classification refers to an application-specific configuration file. Properties are the settings you want to change in that file. For more information, see [Configuring Applications](https://docs.aws.amazon.com/emr/latest/ReleaseGuide/emr-configure-apps.html) in the *Amazon EMR Release Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ configurationProperty
//
//   configurationProperty := &configurationProperty{
//   	classification: jsii.String("classification"),
//   	configurationProperties: map[string]*string{
//   		"configurationPropertiesKey": jsii.String("configurationProperties"),
//   	},
//   	configurations: []interface{}{
//   		&configurationProperty{
//   			classification: jsii.String("classification"),
//   			configurationProperties: map[string]*string{
//   				"configurationPropertiesKey": jsii.String("configurationProperties"),
//   			},
//   			configurations: []interface{}{
//   				configurationProperty_,
//   			},
//   		},
//   	},
//   }
//
type CfnCluster_ConfigurationProperty struct {
	// The classification within a configuration.
	Classification *string `field:"optional" json:"classification" yaml:"classification"`
	// A list of additional configurations to apply within a configuration object.
	ConfigurationProperties interface{} `field:"optional" json:"configurationProperties" yaml:"configurationProperties"`
	// A list of additional configurations to apply within a configuration object.
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
}

