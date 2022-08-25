package awslakeformation


// Properties for defining a `CfnResource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceProps := &cfnResourceProps{
//   	resourceArn: jsii.String("resourceArn"),
//   	useServiceLinkedRole: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnResourceProps struct {
	// The Amazon Resource Name (ARN) of the resource.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// Designates a trusted caller, an IAM principal, by registering this caller with the Data Catalog.
	UseServiceLinkedRole interface{} `field:"required" json:"useServiceLinkedRole" yaml:"useServiceLinkedRole"`
	// The IAM role that registered a resource.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

