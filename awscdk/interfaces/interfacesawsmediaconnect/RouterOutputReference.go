package interfacesawsmediaconnect


// A reference to a RouterOutput resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routerOutputReference := &RouterOutputReference{
//   	RouterOutputArn: jsii.String("routerOutputArn"),
//   }
//
type RouterOutputReference struct {
	// The Arn of the RouterOutput resource.
	RouterOutputArn *string `field:"required" json:"routerOutputArn" yaml:"routerOutputArn"`
}

