package mixinsawsappconfig


// A validator provides a syntactic or semantic check to ensure the configuration that you want to deploy functions as intended.
//
// To validate your application configuration data, you provide a schema or an AWS Lambda function that runs against the configuration. The configuration deployment or update can only proceed when the configuration data is valid. For more information, see [About validators](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-configuration-profile.html#appconfig-creating-configuration-and-profile-validators) in the *AWS AppConfig User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   validatorsProperty := &ValidatorsProperty{
//   	Content: jsii.String("content"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-configurationprofile-validators.html
//
type CfnConfigurationProfilePropsMixin_ValidatorsProperty struct {
	// Either the JSON Schema content or the Amazon Resource Name (ARN) of an Lambda function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-configurationprofile-validators.html#cfn-appconfig-configurationprofile-validators-content
	//
	Content *string `field:"optional" json:"content" yaml:"content"`
	// AWS AppConfig supports validators of type `JSON_SCHEMA` and `LAMBDA`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-configurationprofile-validators.html#cfn-appconfig-configurationprofile-validators-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

