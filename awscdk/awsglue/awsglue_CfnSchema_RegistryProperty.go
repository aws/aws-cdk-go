package awsglue


// Specifies a registry in the AWS Glue Schema Registry.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   registryProperty := &registryProperty{
//   	arn: jsii.String("arn"),
//   	name: jsii.String("name"),
//   }
//
type CfnSchema_RegistryProperty struct {
	// The Amazon Resource Name (ARN) of the registry.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The name of the registry.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

