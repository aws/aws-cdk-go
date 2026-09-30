package interfacesawsroute53


// A reference to a DelegationSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   delegationSetReference := &DelegationSetReference{
//   	DelegationSetArn: jsii.String("delegationSetArn"),
//   }
//
type DelegationSetReference struct {
	// The Arn of the DelegationSet resource.
	DelegationSetArn *string `field:"required" json:"delegationSetArn" yaml:"delegationSetArn"`
}

