package awscodepipelineactions


// Properties for configuring service-managed (Organizations) permissions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   organizationsDeploymentProps := &OrganizationsDeploymentProps{
//   	AutoDeployment: awscdk.Aws_codepipeline_actions.StackSetOrganizationsAutoDeployment_ENABLED,
//   }
//
type OrganizationsDeploymentProps struct {
	// Automatically deploy to new accounts added to Organizational Units.
	//
	// Whether AWS CloudFormation StackSets automatically deploys to AWS
	// Organizations accounts that are added to a target organization or
	// organizational unit (OU).
	// Default: Disabled.
	//
	AutoDeployment StackSetOrganizationsAutoDeployment `field:"optional" json:"autoDeployment" yaml:"autoDeployment"`
}

