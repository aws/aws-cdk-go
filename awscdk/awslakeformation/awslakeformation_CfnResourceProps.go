package awslakeformation


// Properties for defining a `CfnResource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceProps := &CfnResourceProps{
//   	ResourceArn: jsii.String("resourceArn"),
//   	UseServiceLinkedRole: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	RoleArn: jsii.String("roleArn"),
//   	WithFederation: jsii.Boolean(false),
//   }
//
type CfnResourceProps struct {
	// The Amazon Resource Name (ARN) of the resource.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// Designates a trusted caller, an IAM principal, by registering this caller with the Data Catalog .
	UseServiceLinkedRole interface{} `field:"required" json:"useServiceLinkedRole" yaml:"useServiceLinkedRole"`
	// The IAM role that registered a resource.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// `AWS::LakeFormation::Resource.WithFederation`.
	WithFederation interface{} `field:"optional" json:"withFederation" yaml:"withFederation"`
}

