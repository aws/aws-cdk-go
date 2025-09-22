package awscognito


// A reference to a IdentityPool resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identityPoolReference := &IdentityPoolReference{
//   	IdentityPoolId: jsii.String("identityPoolId"),
//   }
//
type IdentityPoolReference struct {
	// The Id of the IdentityPool resource.
	IdentityPoolId *string `field:"required" json:"identityPoolId" yaml:"identityPoolId"`
}

