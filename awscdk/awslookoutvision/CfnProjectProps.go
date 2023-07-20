package awslookoutvision


// Properties for defining a `CfnProject`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProjectProps := &CfnProjectProps{
//   	ProjectName: jsii.String("projectName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutvision-project.html
//
type CfnProjectProps struct {
	// The name of the project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutvision-project.html#cfn-lookoutvision-project-projectname
	//
	ProjectName *string `field:"required" json:"projectName" yaml:"projectName"`
}

