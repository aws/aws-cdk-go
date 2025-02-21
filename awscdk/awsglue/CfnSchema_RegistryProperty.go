package awsglue


// Specifies a registry in the AWS Glue Schema Registry.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   registryProperty := &RegistryProperty{
//   	Arn: jsii.String("arn"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-schema-registry.html
//
type CfnSchema_RegistryProperty struct {
	// The Amazon Resource Name (ARN) of the registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-schema-registry.html#cfn-glue-schema-registry-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The name of the registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-schema-registry.html#cfn-glue-schema-registry-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

