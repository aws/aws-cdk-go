package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authorizationDataProperty := &AuthorizationDataProperty{
//   	Oauth2: &OAuth2AuthorizationDataProperty{
//   		AuthorizationUrl: jsii.String("authorizationUrl"),
//
//   		// the properties below are optional
//   		UserId: jsii.String("userId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-authorizationdata.html
//
type CfnGatewayTarget_AuthorizationDataProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-authorizationdata.html#cfn-bedrockagentcore-gatewaytarget-authorizationdata-oauth2
	//
	Oauth2 interface{} `field:"required" json:"oauth2" yaml:"oauth2"`
}

