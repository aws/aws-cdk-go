package interfacesawslambda


// A reference to a RuntimeAlias resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runtimeAliasReference := &RuntimeAliasReference{
//   	RuntimeAliasArn: jsii.String("runtimeAliasArn"),
//   }
//
type RuntimeAliasReference struct {
	// The Arn of the RuntimeAlias resource.
	RuntimeAliasArn *string `field:"required" json:"runtimeAliasArn" yaml:"runtimeAliasArn"`
}

