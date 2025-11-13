package interfacesawscassandra


// A reference to a Type resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   typeReference := &TypeReference{
//   	KeyspaceName: jsii.String("keyspaceName"),
//   	TypeName: jsii.String("typeName"),
//   }
//
type TypeReference struct {
	// The KeyspaceName of the Type resource.
	KeyspaceName *string `field:"required" json:"keyspaceName" yaml:"keyspaceName"`
	// The TypeName of the Type resource.
	TypeName *string `field:"required" json:"typeName" yaml:"typeName"`
}

