package awsiam


// A reference to a ManagedPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedPolicyReference := &ManagedPolicyReference{
//   	PolicyArn: jsii.String("policyArn"),
//   }
//
type ManagedPolicyReference struct {
	// The PolicyArn of the ManagedPolicy resource.
	PolicyArn *string `field:"required" json:"policyArn" yaml:"policyArn"`
}

