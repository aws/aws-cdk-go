// The CDK Construct Library for AWS::Amplify
package awscdkamplifyalpha


// Sub domain settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import amplify_alpha "github.com/aws/aws-cdk-go/awscdkamplifyalpha"
//
//   var branch branch
//
//   subDomain := &subDomain{
//   	branch: branch,
//
//   	// the properties below are optional
//   	prefix: jsii.String("prefix"),
//   }
//
// Experimental.
type SubDomain struct {
	// The branch.
	// Experimental.
	Branch IBranch `field:"required" json:"branch" yaml:"branch"`
	// The prefix.
	//
	// Use '' to map to the root of the domain.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

