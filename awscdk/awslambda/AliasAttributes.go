package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var version version
//
//   aliasAttributes := &AliasAttributes{
//   	AliasName: jsii.String("aliasName"),
//   	AliasVersion: version,
//   }
//
type AliasAttributes struct {
	AliasName *string `field:"required" json:"aliasName" yaml:"aliasName"`
	AliasVersion IVersion `field:"required" json:"aliasVersion" yaml:"aliasVersion"`
}

