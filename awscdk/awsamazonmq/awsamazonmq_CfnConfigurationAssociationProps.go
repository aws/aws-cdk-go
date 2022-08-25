package awsamazonmq


// Properties for defining a `CfnConfigurationAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationAssociationProps := &cfnConfigurationAssociationProps{
//   	broker: jsii.String("broker"),
//   	configuration: &configurationIdProperty{
//   		id: jsii.String("id"),
//   		revision: jsii.Number(123),
//   	},
//   }
//
type CfnConfigurationAssociationProps struct {
	// The broker to associate with a configuration.
	Broker *string `field:"required" json:"broker" yaml:"broker"`
	// The configuration to associate with a broker.
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
}

