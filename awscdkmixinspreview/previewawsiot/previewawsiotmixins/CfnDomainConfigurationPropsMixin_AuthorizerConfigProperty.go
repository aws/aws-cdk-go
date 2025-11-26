package previewawsiotmixins


// An object that specifies the authorization service for a domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   authorizerConfigProperty := &AuthorizerConfigProperty{
//   	AllowAuthorizerOverride: jsii.Boolean(false),
//   	DefaultAuthorizerName: jsii.String("defaultAuthorizerName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-authorizerconfig.html
//
type CfnDomainConfigurationPropsMixin_AuthorizerConfigProperty struct {
	// A Boolean that specifies whether the domain configuration's authorization service can be overridden.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-authorizerconfig.html#cfn-iot-domainconfiguration-authorizerconfig-allowauthorizeroverride
	//
	AllowAuthorizerOverride interface{} `field:"optional" json:"allowAuthorizerOverride" yaml:"allowAuthorizerOverride"`
	// The name of the authorization service for a domain configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-authorizerconfig.html#cfn-iot-domainconfiguration-authorizerconfig-defaultauthorizername
	//
	DefaultAuthorizerName *string `field:"optional" json:"defaultAuthorizerName" yaml:"defaultAuthorizerName"`
}

