package awsbedrockagentcore


// The configuration for a component within a configuration bundle.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configuration interface{}
//
//   componentConfigurationProperty := &ComponentConfigurationProperty{
//   	Configuration: configuration,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-configurationbundle-componentconfiguration.html
//
type CfnConfigurationBundle_ComponentConfigurationProperty struct {
	// The configuration values as a flexible JSON document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-configurationbundle-componentconfiguration.html#cfn-bedrockagentcore-configurationbundle-componentconfiguration-configuration
	//
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
}

