package mixinsawssagemaker


// Studio settings.
//
// If these settings are applied on a user level, they take priority over the settings applied on a domain level.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   studioWebPortalSettingsProperty := &StudioWebPortalSettingsProperty{
//   	HiddenAppTypes: []*string{
//   		jsii.String("hiddenAppTypes"),
//   	},
//   	HiddenInstanceTypes: []*string{
//   		jsii.String("hiddenInstanceTypes"),
//   	},
//   	HiddenMlTools: []*string{
//   		jsii.String("hiddenMlTools"),
//   	},
//   	HiddenSageMakerImageVersionAliases: []interface{}{
//   		&HiddenSageMakerImageProperty{
//   			SageMakerImageName: jsii.String("sageMakerImageName"),
//   			VersionAliases: []*string{
//   				jsii.String("versionAliases"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-studiowebportalsettings.html
//
type CfnDomainPropsMixin_StudioWebPortalSettingsProperty struct {
	// The [Applications supported in Studio](https://docs.aws.amazon.com/sagemaker/latest/dg/studio-updated-apps.html) that are hidden from the Studio left navigation pane.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-studiowebportalsettings.html#cfn-sagemaker-domain-studiowebportalsettings-hiddenapptypes
	//
	HiddenAppTypes *[]*string `field:"optional" json:"hiddenAppTypes" yaml:"hiddenAppTypes"`
	// The instance types you are hiding from the Studio user interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-studiowebportalsettings.html#cfn-sagemaker-domain-studiowebportalsettings-hiddeninstancetypes
	//
	HiddenInstanceTypes *[]*string `field:"optional" json:"hiddenInstanceTypes" yaml:"hiddenInstanceTypes"`
	// The machine learning tools that are hidden from the Studio left navigation pane.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-studiowebportalsettings.html#cfn-sagemaker-domain-studiowebportalsettings-hiddenmltools
	//
	HiddenMlTools *[]*string `field:"optional" json:"hiddenMlTools" yaml:"hiddenMlTools"`
	// The version aliases you are hiding from the Studio user interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-studiowebportalsettings.html#cfn-sagemaker-domain-studiowebportalsettings-hiddensagemakerimageversionaliases
	//
	HiddenSageMakerImageVersionAliases interface{} `field:"optional" json:"hiddenSageMakerImageVersionAliases" yaml:"hiddenSageMakerImageVersionAliases"`
}

