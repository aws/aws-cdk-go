package awscodebuild


// Properties for defining a `CfnBuild`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBuildProps := &CfnBuildProps{
//   	ProjectName: jsii.String("projectName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-build.html
//
type CfnBuildProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-build.html#cfn-codebuild-build-projectname
	//
	ProjectName *string `field:"optional" json:"projectName" yaml:"projectName"`
}

