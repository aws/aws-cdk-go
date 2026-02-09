package cxapi


// Models an AWS execution environment, for use within the CDK toolkit.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environment := &Environment{
//   	Account: jsii.String("account"),
//   	Name: jsii.String("name"),
//   	Region: jsii.String("region"),
//   }
//
// Deprecated: The official definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
type Environment struct {
	// The AWS account this environment deploys into.
	// Deprecated: The official definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	Account *string `field:"required" json:"account" yaml:"account"`
	// The arbitrary name of this environment (user-set, or at least user-meaningful).
	// Deprecated: The official definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The AWS region name where this environment deploys into.
	// Deprecated: The official definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	Region *string `field:"required" json:"region" yaml:"region"`
}

