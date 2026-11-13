package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   eksPodSecurityContextProperty := &EksPodSecurityContextProperty{
//   	FsGroup: jsii.Number(123),
//   	RunAsGroup: jsii.Number(123),
//   	RunAsNonRoot: jsii.Boolean(false),
//   	RunAsUser: jsii.Number(123),
//   	SeccompProfile: &EksSeccompProfileProperty{
//   		LocalhostProfile: jsii.String("localhostProfile"),
//   		Type: jsii.String("type"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekspodsecuritycontext.html
//
type CfnJobDefinitionPropsMixin_EksPodSecurityContextProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekspodsecuritycontext.html#cfn-batch-jobdefinition-ekspodsecuritycontext-fsgroup
	//
	FsGroup *float64 `field:"optional" json:"fsGroup" yaml:"fsGroup"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekspodsecuritycontext.html#cfn-batch-jobdefinition-ekspodsecuritycontext-runasgroup
	//
	RunAsGroup *float64 `field:"optional" json:"runAsGroup" yaml:"runAsGroup"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekspodsecuritycontext.html#cfn-batch-jobdefinition-ekspodsecuritycontext-runasnonroot
	//
	RunAsNonRoot interface{} `field:"optional" json:"runAsNonRoot" yaml:"runAsNonRoot"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekspodsecuritycontext.html#cfn-batch-jobdefinition-ekspodsecuritycontext-runasuser
	//
	RunAsUser *float64 `field:"optional" json:"runAsUser" yaml:"runAsUser"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekspodsecuritycontext.html#cfn-batch-jobdefinition-ekspodsecuritycontext-seccompprofile
	//
	SeccompProfile interface{} `field:"optional" json:"seccompProfile" yaml:"seccompProfile"`
}

