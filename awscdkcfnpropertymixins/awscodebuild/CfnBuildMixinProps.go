package awscodebuild


// Properties for CfnBuildPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnBuildMixinProps := &CfnBuildMixinProps{
//   	ProjectName: jsii.String("projectName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-build.html
//
type CfnBuildMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-build.html#cfn-codebuild-build-projectname
	//
	ProjectName *string `field:"optional" json:"projectName" yaml:"projectName"`
}

