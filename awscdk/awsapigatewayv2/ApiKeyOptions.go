package awsapigatewayv2


// The options for creating an API Key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiKeyOptions := &ApiKeyOptions{
//   	ApiKeyName: jsii.String("apiKeyName"),
//   	Description: jsii.String("description"),
//   	Value: jsii.String("value"),
//   }
//
type ApiKeyOptions struct {
	// A name for the API key.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the API key name.
	// Default: automatically generated name.
	//
	ApiKeyName *string `field:"optional" json:"apiKeyName" yaml:"apiKeyName"`
	// A description of the purpose of the API key.
	// Default: none.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The value of the API key.
	//
	// Must be at least 20 characters long.
	// Default: none.
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

