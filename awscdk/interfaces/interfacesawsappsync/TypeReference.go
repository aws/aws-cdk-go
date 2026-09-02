package interfacesawsappsync


// A reference to a Type resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   typeReference := &TypeReference{
//   	TypeArn: jsii.String("typeArn"),
//   }
//
type TypeReference struct {
	// The Arn of the Type resource.
	TypeArn *string `field:"required" json:"typeArn" yaml:"typeArn"`
}

