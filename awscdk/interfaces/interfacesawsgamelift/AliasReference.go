package interfacesawsgamelift


// A reference to a Alias resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aliasReference := &AliasReference{
//   	AliasArn: jsii.String("aliasArn"),
//   	AliasId: jsii.String("aliasId"),
//   }
//
type AliasReference struct {
	// The ARN of the Alias resource.
	AliasArn *string `field:"required" json:"aliasArn" yaml:"aliasArn"`
	// The AliasId of the Alias resource.
	AliasId *string `field:"required" json:"aliasId" yaml:"aliasId"`
}

