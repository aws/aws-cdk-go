package awsstepfunctionstasks


// A configuration specification to be used when provisioning virtual clusters, which can include configurations for applications and software bundled with Amazon EMR on EKS.
//
// A configuration consists of a classification, properties, and optional nested configurations.
// A classification refers to an application-specific configuration file.
// Properties are the settings you want to change in that file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var applicationConfiguration_ applicationConfiguration
//   var classification classification
//
//   applicationConfiguration := &applicationConfiguration{
//   	classification: classification,
//
//   	// the properties below are optional
//   	nestedConfig: []*applicationConfiguration{
//   		&applicationConfiguration{
//   			classification: classification,
//
//   			// the properties below are optional
//   			nestedConfig: []*applicationConfiguration{
//   				applicationConfiguration_,
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
// See: https://docs.aws.amazon.com/emr/latest/ReleaseGuide/emr-configure-apps.html
//
type ApplicationConfiguration struct {
	// The classification within a configuration.
	//
	// Length Constraints: Minimum length of 1. Maximum length of 1024.
	Classification Classification `field:"required" json:"classification" yaml:"classification"`
	// A list of additional configurations to apply within a configuration object.
	//
	// Array Members: Maximum number of 100 items.
	NestedConfig *[]*ApplicationConfiguration `field:"optional" json:"nestedConfig" yaml:"nestedConfig"`
	// A set of properties specified within a configuration classification.
	//
	// Map Entries: Maximum number of 100 items.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}

