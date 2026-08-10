package awscodebuild


// Properties for CfnSandboxPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnSandboxMixinProps := &CfnSandboxMixinProps{
//   	ProjectName: jsii.String("projectName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-sandbox.html
//
type CfnSandboxMixinProps struct {
	// The CodeBuild project name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-sandbox.html#cfn-codebuild-sandbox-projectname
	//
	ProjectName *string `field:"optional" json:"projectName" yaml:"projectName"`
}

