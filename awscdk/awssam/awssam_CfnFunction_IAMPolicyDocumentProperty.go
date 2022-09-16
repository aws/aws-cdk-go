package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var statement interface{}
//
//   iAMPolicyDocumentProperty := map[string]interface{}{
//   	"statement": statement,
//
//   	// the properties below are optional
//   	"version": jsii.String("version"),
//   }
//
type CfnFunction_IAMPolicyDocumentProperty struct {
	// `CfnFunction.IAMPolicyDocumentProperty.Statement`.
	Statement interface{} `field:"required" json:"statement" yaml:"statement"`
	// `CfnFunction.IAMPolicyDocumentProperty.Version`.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

