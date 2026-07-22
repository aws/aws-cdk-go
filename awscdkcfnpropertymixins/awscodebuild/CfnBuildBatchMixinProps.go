package awscodebuild


// Properties for CfnBuildBatchPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnBuildBatchMixinProps := &CfnBuildBatchMixinProps{
//   	ProjectName: jsii.String("projectName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-buildbatch.html
//
type CfnBuildBatchMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-buildbatch.html#cfn-codebuild-buildbatch-projectname
	//
	ProjectName *string `field:"optional" json:"projectName" yaml:"projectName"`
}

