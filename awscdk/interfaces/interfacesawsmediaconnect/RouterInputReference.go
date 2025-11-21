package interfacesawsmediaconnect


// A reference to a RouterInput resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routerInputReference := &RouterInputReference{
//   	RouterInputArn: jsii.String("routerInputArn"),
//   }
//
type RouterInputReference struct {
	// The Arn of the RouterInput resource.
	RouterInputArn *string `field:"required" json:"routerInputArn" yaml:"routerInputArn"`
}

