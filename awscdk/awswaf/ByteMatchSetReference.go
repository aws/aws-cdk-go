package awswaf


// A reference to a ByteMatchSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   byteMatchSetReference := &ByteMatchSetReference{
//   	ByteMatchSetId: jsii.String("byteMatchSetId"),
//   }
//
type ByteMatchSetReference struct {
	// The Id of the ByteMatchSet resource.
	ByteMatchSetId *string `field:"required" json:"byteMatchSetId" yaml:"byteMatchSetId"`
}

