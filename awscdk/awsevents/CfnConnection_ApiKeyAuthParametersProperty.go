package awsevents


// Contains the API key authorization parameters for the connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiKeyAuthParametersProperty := &ApiKeyAuthParametersProperty{
//   	ApiKeyName: jsii.String("apiKeyName"),
//   	ApiKeyValue: jsii.String("apiKeyValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-apikeyauthparameters.html
//
type CfnConnection_ApiKeyAuthParametersProperty struct {
	// The name of the API key to use for authorization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-apikeyauthparameters.html#cfn-events-connection-apikeyauthparameters-apikeyname
	//
	ApiKeyName *string `field:"required" json:"apiKeyName" yaml:"apiKeyName"`
	// The value for the API key to use for authorization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-apikeyauthparameters.html#cfn-events-connection-apikeyauthparameters-apikeyvalue
	//
	ApiKeyValue *string `field:"required" json:"apiKeyValue" yaml:"apiKeyValue"`
}

