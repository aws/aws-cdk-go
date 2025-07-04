package awsamazonmq


// Properties for defining a `CfnConfigurationAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationAssociationProps := &CfnConfigurationAssociationProps{
//   	Broker: jsii.String("broker"),
//   	Configuration: &ConfigurationIdProperty{
//   		Id: jsii.String("id"),
//   		Revision: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configurationassociation.html
//
type CfnConfigurationAssociationProps struct {
	// The broker to associate with a configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configurationassociation.html#cfn-amazonmq-configurationassociation-broker
	//
	Broker *string `field:"required" json:"broker" yaml:"broker"`
	// The configuration to associate with a broker.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configurationassociation.html#cfn-amazonmq-configurationassociation-configuration
	//
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
}

