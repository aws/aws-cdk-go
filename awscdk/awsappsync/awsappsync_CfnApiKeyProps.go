package awsappsync


// Properties for defining a `CfnApiKey`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApiKeyProps := &cfnApiKeyProps{
//   	apiId: jsii.String("apiId"),
//
//   	// the properties below are optional
//   	apiKeyId: jsii.String("apiKeyId"),
//   	description: jsii.String("description"),
//   	expires: jsii.Number(123),
//   }
//
type CfnApiKeyProps struct {
	// Unique AWS AppSync GraphQL API ID for this API key.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The API key ID.
	ApiKeyId *string `field:"optional" json:"apiKeyId" yaml:"apiKeyId"`
	// Unique description of your API key.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The time after which the API key expires.
	//
	// The date is represented as seconds since the epoch, rounded down to the nearest hour.
	Expires *float64 `field:"optional" json:"expires" yaml:"expires"`
}

