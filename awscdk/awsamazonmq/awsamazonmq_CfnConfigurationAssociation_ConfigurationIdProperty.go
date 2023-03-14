package awsamazonmq


// The `ConfigurationId` property type specifies a configuration Id and the revision of a configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationIdProperty := &configurationIdProperty{
//   	id: jsii.String("id"),
//   	revision: jsii.Number(123),
//   }
//
type CfnConfigurationAssociation_ConfigurationIdProperty struct {
	// The unique ID that Amazon MQ generates for the configuration.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The revision number of the configuration.
	Revision *float64 `field:"required" json:"revision" yaml:"revision"`
}

