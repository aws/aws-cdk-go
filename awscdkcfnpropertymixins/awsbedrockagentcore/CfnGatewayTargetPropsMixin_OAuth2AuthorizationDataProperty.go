package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   oAuth2AuthorizationDataProperty := &OAuth2AuthorizationDataProperty{
//   	AuthorizationUrl: jsii.String("authorizationUrl"),
//   	UserId: jsii.String("userId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-oauth2authorizationdata.html
//
type CfnGatewayTargetPropsMixin_OAuth2AuthorizationDataProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-oauth2authorizationdata.html#cfn-bedrockagentcore-gatewaytarget-oauth2authorizationdata-authorizationurl
	//
	AuthorizationUrl *string `field:"optional" json:"authorizationUrl" yaml:"authorizationUrl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-oauth2authorizationdata.html#cfn-bedrockagentcore-gatewaytarget-oauth2authorizationdata-userid
	//
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
}

