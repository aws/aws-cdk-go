package interfacesawsiot


// A reference to a Policy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyReference := &PolicyReference{
//   	PolicyArn: jsii.String("policyArn"),
//   	PolicyName: jsii.String("policyName"),
//   }
//
type PolicyReference struct {
	// The ARN of the Policy resource.
	PolicyArn *string `field:"required" json:"policyArn" yaml:"policyArn"`
	// The PolicyName of the Policy resource.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
}

