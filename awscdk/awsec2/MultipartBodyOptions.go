package awsec2


// Options when creating `MultipartBody`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multipartBodyOptions := &MultipartBodyOptions{
//   	ContentType: jsii.String("contentType"),
//
//   	// the properties below are optional
//   	Body: jsii.String("body"),
//   	TransferEncoding: jsii.String("transferEncoding"),
//   }
//
type MultipartBodyOptions struct {
	// `Content-Type` header of this part.
	//
	// Some examples of content types:
	// * `text/x-shellscript; charset="utf-8"` (shell script)
	// * `text/cloud-boothook; charset="utf-8"` (shell script executed during boot phase)
	//
	// For Linux shell scripts use `text/x-shellscript`.
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// The body of message.
	// Default: undefined - body will not be added to part.
	//
	Body *string `field:"optional" json:"body" yaml:"body"`
	// `Content-Transfer-Encoding` header specifying part encoding.
	// Default: undefined - body is not encoded.
	//
	TransferEncoding *string `field:"optional" json:"transferEncoding" yaml:"transferEncoding"`
}

