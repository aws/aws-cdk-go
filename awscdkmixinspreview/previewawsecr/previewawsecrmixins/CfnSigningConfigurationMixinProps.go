package previewawsecrmixins


// Properties for CfnSigningConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSigningConfigurationMixinProps := &CfnSigningConfigurationMixinProps{
//   	Rules: []interface{}{
//   		&RuleProperty{
//   			RepositoryFilters: []interface{}{
//   				&RepositoryFilterProperty{
//   					Filter: jsii.String("filter"),
//   					FilterType: jsii.String("filterType"),
//   				},
//   			},
//   			SigningProfileArn: jsii.String("signingProfileArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-signingconfiguration.html
//
type CfnSigningConfigurationMixinProps struct {
	// Array of signing rules that define which repositories should be signed and with which signing profiles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-signingconfiguration.html#cfn-ecr-signingconfiguration-rules
	//
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
}

