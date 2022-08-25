package cxapi


// Models an AWS execution environment, for use within the CDK toolkit.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environment := &environment{
//   	account: jsii.String("account"),
//   	name: jsii.String("name"),
//   	region: jsii.String("region"),
//   }
//
type Environment struct {
	// The AWS account this environment deploys into.
	Account *string `field:"required" json:"account" yaml:"account"`
	// The arbitrary name of this environment (user-set, or at least user-meaningful).
	Name *string `field:"required" json:"name" yaml:"name"`
	// The AWS region name where this environment deploys into.
	Region *string `field:"required" json:"region" yaml:"region"`
}

