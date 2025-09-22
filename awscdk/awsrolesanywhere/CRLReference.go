package awsrolesanywhere


// A reference to a CRL resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cRLReference := &CRLReference{
//   	CrlId: jsii.String("crlId"),
//   }
//
type CRLReference struct {
	// The CrlId of the CRL resource.
	CrlId *string `field:"required" json:"crlId" yaml:"crlId"`
}

