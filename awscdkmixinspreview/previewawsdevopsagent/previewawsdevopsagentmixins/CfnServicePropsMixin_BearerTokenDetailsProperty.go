package previewawsdevopsagentmixins


// Bearer token authentication details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   bearerTokenDetailsProperty := &BearerTokenDetailsProperty{
//   	AuthorizationHeader: jsii.String("authorizationHeader"),
//   	TokenName: jsii.String("tokenName"),
//   	TokenValue: jsii.String("tokenValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-bearertokendetails.html
//
type CfnServicePropsMixin_BearerTokenDetailsProperty struct {
	// HTTP header name to send the bearer token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-bearertokendetails.html#cfn-devopsagent-service-bearertokendetails-authorizationheader
	//
	// Default: - "Authorization".
	//
	AuthorizationHeader *string `field:"optional" json:"authorizationHeader" yaml:"authorizationHeader"`
	// User friendly bearer token name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-bearertokendetails.html#cfn-devopsagent-service-bearertokendetails-tokenname
	//
	TokenName *string `field:"optional" json:"tokenName" yaml:"tokenName"`
	// Bearer token value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-bearertokendetails.html#cfn-devopsagent-service-bearertokendetails-tokenvalue
	//
	TokenValue *string `field:"optional" json:"tokenValue" yaml:"tokenValue"`
}

