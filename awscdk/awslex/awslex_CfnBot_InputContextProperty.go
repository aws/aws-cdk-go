package awslex


// The name of a context that must be active for an intent to be selected by Amazon Lex .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputContextProperty := &inputContextProperty{
//   	name: jsii.String("name"),
//   }
//
type CfnBot_InputContextProperty struct {
	// The name of the context.
	Name *string `field:"required" json:"name" yaml:"name"`
}

