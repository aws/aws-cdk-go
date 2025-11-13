package interfacesawsconfig


// A reference to a ConfigurationAggregator resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationAggregatorReference := &ConfigurationAggregatorReference{
//   	ConfigurationAggregatorArn: jsii.String("configurationAggregatorArn"),
//   	ConfigurationAggregatorName: jsii.String("configurationAggregatorName"),
//   }
//
type ConfigurationAggregatorReference struct {
	// The ARN of the ConfigurationAggregator resource.
	ConfigurationAggregatorArn *string `field:"required" json:"configurationAggregatorArn" yaml:"configurationAggregatorArn"`
	// The ConfigurationAggregatorName of the ConfigurationAggregator resource.
	ConfigurationAggregatorName *string `field:"required" json:"configurationAggregatorName" yaml:"configurationAggregatorName"`
}

