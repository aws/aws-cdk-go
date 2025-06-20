package awsappsync


// Properties for defining a `CfnApiKey`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApiKeyProps := &CfnApiKeyProps{
//   	ApiId: jsii.String("apiId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Expires: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apikey.html
//
type CfnApiKeyProps struct {
	// Unique AWS AppSync GraphQL API ID for this API key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apikey.html#cfn-appsync-apikey-apiid
	//
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// Unique description of your API key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apikey.html#cfn-appsync-apikey-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The time after which the API key expires.
	//
	// The date is represented as seconds since the epoch, rounded down to the nearest hour.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apikey.html#cfn-appsync-apikey-expires
	//
	Expires *float64 `field:"optional" json:"expires" yaml:"expires"`
}

