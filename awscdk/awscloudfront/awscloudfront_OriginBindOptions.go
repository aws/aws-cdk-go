package awscloudfront


// Options passed to Origin.bind().
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   originBindOptions := &originBindOptions{
//   	originId: jsii.String("originId"),
//   }
//
type OriginBindOptions struct {
	// The identifier of this Origin, as assigned by the Distribution this Origin has been used added to.
	OriginId *string `field:"required" json:"originId" yaml:"originId"`
}

