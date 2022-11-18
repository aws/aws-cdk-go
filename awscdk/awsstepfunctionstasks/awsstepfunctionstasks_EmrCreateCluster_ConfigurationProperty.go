package awsstepfunctionstasks


// An optional configuration specification to be used when provisioning cluster instances, which can include configurations for applications and software bundled with Amazon EMR.
//
// See the RunJobFlow API for complete documentation on input parameters.
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
//   	configurations: []*configurationProperty{
//   		&configurationProperty{
//   			classification: jsii.String("classification"),
//   			configurations: []*configurationProperty{
//   				configurationProperty_,
//   			},
//   			properties: map[string]*string{
//   				"propertiesKey": jsii.String("properties"),
//   			},
//   		},
//   	},
//   	properties: map[string]*string{
//   		"propertiesKey": jsii.String("properties"),
//   	},
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_Configuration.html
//
type EmrCreateCluster_ConfigurationProperty struct {
	// The classification within a configuration.
	Classification *string `field:"optional" json:"classification" yaml:"classification"`
	// A list of additional configurations to apply within a configuration object.
	Configurations *[]*EmrCreateCluster_ConfigurationProperty `field:"optional" json:"configurations" yaml:"configurations"`
	// A set of properties specified within a configuration classification.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}

