package awsresiliencehub


// A reference to a ResiliencyPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resiliencyPolicyReference := &ResiliencyPolicyReference{
//   	PolicyArn: jsii.String("policyArn"),
//   }
//
type ResiliencyPolicyReference struct {
	// The PolicyArn of the ResiliencyPolicy resource.
	PolicyArn *string `field:"required" json:"policyArn" yaml:"policyArn"`
}

