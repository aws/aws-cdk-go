package awslicensemanager


// A reference to a Grant resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grantReference := &GrantReference{
//   	GrantArn: jsii.String("grantArn"),
//   }
//
type GrantReference struct {
	// The GrantArn of the Grant resource.
	GrantArn *string `field:"required" json:"grantArn" yaml:"grantArn"`
}

