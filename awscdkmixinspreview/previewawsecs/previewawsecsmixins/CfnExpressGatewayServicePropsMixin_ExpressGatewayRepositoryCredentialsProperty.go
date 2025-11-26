package previewawsecsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   expressGatewayRepositoryCredentialsProperty := &ExpressGatewayRepositoryCredentialsProperty{
//   	CredentialsParameter: jsii.String("credentialsParameter"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayrepositorycredentials.html
//
type CfnExpressGatewayServicePropsMixin_ExpressGatewayRepositoryCredentialsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayrepositorycredentials.html#cfn-ecs-expressgatewayservice-expressgatewayrepositorycredentials-credentialsparameter
	//
	CredentialsParameter *string `field:"optional" json:"credentialsParameter" yaml:"credentialsParameter"`
}

