package awswaf


// A reference to a XssMatchSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   xssMatchSetReference := &XssMatchSetReference{
//   	XssMatchSetId: jsii.String("xssMatchSetId"),
//   }
//
type XssMatchSetReference struct {
	// The Id of the XssMatchSet resource.
	XssMatchSetId *string `field:"required" json:"xssMatchSetId" yaml:"xssMatchSetId"`
}

