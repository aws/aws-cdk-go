package awsevents


// Contains the API key authorization parameters for the connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiKeyAuthParametersProperty := &apiKeyAuthParametersProperty{
//   	apiKeyName: jsii.String("apiKeyName"),
//   	apiKeyValue: jsii.String("apiKeyValue"),
//   }
//
type CfnConnection_ApiKeyAuthParametersProperty struct {
	// The name of the API key to use for authorization.
	ApiKeyName *string `field:"required" json:"apiKeyName" yaml:"apiKeyName"`
	// The value for the API key to use for authorization.
	ApiKeyValue *string `field:"required" json:"apiKeyValue" yaml:"apiKeyValue"`
}

