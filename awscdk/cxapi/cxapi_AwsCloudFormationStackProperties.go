package cxapi


// Artifact properties for CloudFormation stacks.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   awsCloudFormationStackProperties := &awsCloudFormationStackProperties{
//   	templateFile: jsii.String("templateFile"),
//
//   	// the properties below are optional
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	stackName: jsii.String("stackName"),
//   	terminationProtection: jsii.Boolean(false),
//   }
//
type AwsCloudFormationStackProperties struct {
	// A file relative to the assembly root which contains the CloudFormation template for this stack.
	TemplateFile *string `field:"required" json:"templateFile" yaml:"templateFile"`
	// Values for CloudFormation stack parameters that should be passed when the stack is deployed.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// The name to use for the CloudFormation stack.
	StackName *string `field:"optional" json:"stackName" yaml:"stackName"`
	// Whether to enable termination protection for this stack.
	TerminationProtection *bool `field:"optional" json:"terminationProtection" yaml:"terminationProtection"`
}

