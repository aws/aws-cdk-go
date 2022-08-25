// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Represents the environment a given resource lives in.
//
// Used as the return value for the {@link IResource.env} property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceEnvironment := &resourceEnvironment{
//   	account: jsii.String("account"),
//   	region: jsii.String("region"),
//   }
//
type ResourceEnvironment struct {
	// The AWS account ID that this resource belongs to.
	//
	// Since this can be a Token
	// (for example, when the account is CloudFormation's AWS::AccountId intrinsic),
	// make sure to use Token.compareStrings()
	// instead of just comparing the values for equality.
	Account *string `field:"required" json:"account" yaml:"account"`
	// The AWS region that this resource belongs to.
	//
	// Since this can be a Token
	// (for example, when the region is CloudFormation's AWS::Region intrinsic),
	// make sure to use Token.compareStrings()
	// instead of just comparing the values for equality.
	Region *string `field:"required" json:"region" yaml:"region"`
}

