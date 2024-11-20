package awscdk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   optionsProperty := &OptionsProperty{
//   	InputParams: &S3LocationProperty{
//   		Uri: jsii.String("uri"),
//
//   		// the properties below are optional
//   		VersionId: jsii.String("versionId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-options.html
//
type CfnGuardHook_OptionsProperty struct {
	// S3 Source Location for the Guard files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-options.html#cfn-cloudformation-guardhook-options-inputparams
	//
	InputParams interface{} `field:"optional" json:"inputParams" yaml:"inputParams"`
}

