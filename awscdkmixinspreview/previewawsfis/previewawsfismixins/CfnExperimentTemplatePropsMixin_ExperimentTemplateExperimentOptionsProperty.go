package previewawsfismixins


// Describes the experiment options for an experiment template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   experimentTemplateExperimentOptionsProperty := &ExperimentTemplateExperimentOptionsProperty{
//   	AccountTargeting: jsii.String("accountTargeting"),
//   	EmptyTargetResolutionMode: jsii.String("emptyTargetResolutionMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateexperimentoptions.html
//
type CfnExperimentTemplatePropsMixin_ExperimentTemplateExperimentOptionsProperty struct {
	// The account targeting setting for an experiment template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateexperimentoptions.html#cfn-fis-experimenttemplate-experimenttemplateexperimentoptions-accounttargeting
	//
	AccountTargeting *string `field:"optional" json:"accountTargeting" yaml:"accountTargeting"`
	// The empty target resolution mode for an experiment template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateexperimentoptions.html#cfn-fis-experimenttemplate-experimenttemplateexperimentoptions-emptytargetresolutionmode
	//
	EmptyTargetResolutionMode *string `field:"optional" json:"emptyTargetResolutionMode" yaml:"emptyTargetResolutionMode"`
}

