package previewawsdevopsagentmixins


// API key authentication details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   apiKeyDetailsProperty := &ApiKeyDetailsProperty{
//   	ApiKeyHeader: jsii.String("apiKeyHeader"),
//   	ApiKeyName: jsii.String("apiKeyName"),
//   	ApiKeyValue: jsii.String("apiKeyValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-apikeydetails.html
//
type CfnServicePropsMixin_ApiKeyDetailsProperty struct {
	// HTTP header name to send the API key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-apikeydetails.html#cfn-devopsagent-service-apikeydetails-apikeyheader
	//
	ApiKeyHeader *string `field:"optional" json:"apiKeyHeader" yaml:"apiKeyHeader"`
	// User friendly API key name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-apikeydetails.html#cfn-devopsagent-service-apikeydetails-apikeyname
	//
	ApiKeyName *string `field:"optional" json:"apiKeyName" yaml:"apiKeyName"`
	// API key value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-apikeydetails.html#cfn-devopsagent-service-apikeydetails-apikeyvalue
	//
	ApiKeyValue *string `field:"optional" json:"apiKeyValue" yaml:"apiKeyValue"`
}

