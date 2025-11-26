package previewawsecrmixins


// Properties for CfnRegistryScanningConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRegistryScanningConfigurationMixinProps := &CfnRegistryScanningConfigurationMixinProps{
//   	Rules: []interface{}{
//   		&ScanningRuleProperty{
//   			RepositoryFilters: []interface{}{
//   				&RepositoryFilterProperty{
//   					Filter: jsii.String("filter"),
//   					FilterType: jsii.String("filterType"),
//   				},
//   			},
//   			ScanFrequency: jsii.String("scanFrequency"),
//   		},
//   	},
//   	ScanType: jsii.String("scanType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-registryscanningconfiguration.html
//
type CfnRegistryScanningConfigurationMixinProps struct {
	// The scanning rules associated with the registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-registryscanningconfiguration.html#cfn-ecr-registryscanningconfiguration-rules
	//
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// The type of scanning configured for the registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-registryscanningconfiguration.html#cfn-ecr-registryscanningconfiguration-scantype
	//
	ScanType *string `field:"optional" json:"scanType" yaml:"scanType"`
}

