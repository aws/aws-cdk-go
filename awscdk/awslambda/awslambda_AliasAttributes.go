package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var version version
//
//   aliasAttributes := &aliasAttributes{
//   	aliasName: jsii.String("aliasName"),
//   	aliasVersion: version,
//   }
//
// Experimental.
type AliasAttributes struct {
	// Experimental.
	AliasName *string `field:"required" json:"aliasName" yaml:"aliasName"`
	// Experimental.
	AliasVersion IVersion `field:"required" json:"aliasVersion" yaml:"aliasVersion"`
}

