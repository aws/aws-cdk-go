package awsapprunner


// Identifies a version of code that AWS App Runner refers to within a source code repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceCodeVersionProperty := &SourceCodeVersionProperty{
//   	Type: jsii.String("type"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-sourcecodeversion.html
//
type CfnService_SourceCodeVersionProperty struct {
	// The type of version identifier.
	//
	// For a git-based repository, branches represent versions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-sourcecodeversion.html#cfn-apprunner-service-sourcecodeversion-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// A source code version.
	//
	// For a git-based repository, a branch name maps to a specific version. App Runner uses the most recent commit to the branch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-sourcecodeversion.html#cfn-apprunner-service-sourcecodeversion-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

