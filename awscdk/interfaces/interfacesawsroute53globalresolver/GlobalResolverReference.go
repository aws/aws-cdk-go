package interfacesawsroute53globalresolver


// A reference to a GlobalResolver resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   globalResolverReference := &GlobalResolverReference{
//   	GlobalResolverArn: jsii.String("globalResolverArn"),
//   	GlobalResolverId: jsii.String("globalResolverId"),
//   }
//
type GlobalResolverReference struct {
	// The ARN of the GlobalResolver resource.
	GlobalResolverArn *string `field:"required" json:"globalResolverArn" yaml:"globalResolverArn"`
	// The GlobalResolverId of the GlobalResolver resource.
	GlobalResolverId *string `field:"required" json:"globalResolverId" yaml:"globalResolverId"`
}

