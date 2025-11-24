package mixinsawsomics


// Properties for CfnRunGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRunGroupMixinProps := &CfnRunGroupMixinProps{
//   	MaxCpus: jsii.Number(123),
//   	MaxDuration: jsii.Number(123),
//   	MaxGpus: jsii.Number(123),
//   	MaxRuns: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-rungroup.html
//
type CfnRunGroupMixinProps struct {
	// The group's maximum CPU count setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-rungroup.html#cfn-omics-rungroup-maxcpus
	//
	MaxCpus *float64 `field:"optional" json:"maxCpus" yaml:"maxCpus"`
	// The group's maximum duration setting in minutes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-rungroup.html#cfn-omics-rungroup-maxduration
	//
	MaxDuration *float64 `field:"optional" json:"maxDuration" yaml:"maxDuration"`
	// The maximum GPUs that can be used by a run group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-rungroup.html#cfn-omics-rungroup-maxgpus
	//
	MaxGpus *float64 `field:"optional" json:"maxGpus" yaml:"maxGpus"`
	// The group's maximum concurrent run setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-rungroup.html#cfn-omics-rungroup-maxruns
	//
	MaxRuns *float64 `field:"optional" json:"maxRuns" yaml:"maxRuns"`
	// The group's name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-rungroup.html#cfn-omics-rungroup-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Tags for the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-rungroup.html#cfn-omics-rungroup-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

