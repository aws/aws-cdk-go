package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   authorizationDataProperty := &AuthorizationDataProperty{
//   	Oauth2: &OAuth2AuthorizationDataProperty{
//   		AuthorizationUrl: jsii.String("authorizationUrl"),
//   		UserId: jsii.String("userId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-authorizationdata.html
//
type CfnGatewayTargetPropsMixin_AuthorizationDataProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-authorizationdata.html#cfn-bedrockagentcore-gatewaytarget-authorizationdata-oauth2
	//
	Oauth2 interface{} `field:"optional" json:"oauth2" yaml:"oauth2"`
}

