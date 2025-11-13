package interfacesawslambda


// A reference to a Alias resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aliasReference := &AliasReference{
//   	AliasArn: jsii.String("aliasArn"),
//   }
//
type AliasReference struct {
	// The AliasArn of the Alias resource.
	AliasArn *string `field:"required" json:"aliasArn" yaml:"aliasArn"`
}

