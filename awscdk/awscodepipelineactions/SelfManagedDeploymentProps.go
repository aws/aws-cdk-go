package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for configuring self-managed permissions.
//
// Example:
//   existingAdminRole := iam.Role_FromRoleName(this, jsii.String("AdminRole"), jsii.String("AWSCloudFormationStackSetAdministrationRole"))
//
//   deploymentModel := codepipeline_actions.StackSetDeploymentModel_SelfManaged(&SelfManagedDeploymentProps{
//   	// Use an existing Role. Leave this out to create a new Role.
//   	AdministrationRole: existingAdminRole,
//   })
//
type SelfManagedDeploymentProps struct {
	// The IAM role in the administrator account used to assume execution roles in the target accounts.
	//
	// You must create this role before using the StackSet action.
	//
	// The role needs to be assumable by CloudFormation, and it needs to be able
	// to `sts:AssumeRole` each of the execution roles (whose names are specified
	// in the `executionRoleName` parameter) in each of the target accounts.
	//
	// If you do not specify the role, we assume you have created a role named
	// `AWSCloudFormationStackSetAdministrationRole`.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-self-managed.html
	//
	// Default: - Assume an existing role named `AWSCloudFormationStackSetAdministrationRole` in the same account as the pipeline.
	//
	AdministrationRole awsiam.IRole `field:"optional" json:"administrationRole" yaml:"administrationRole"`
	// The name of the IAM role in the target accounts used to perform stack set operations.
	//
	// You must create these roles in each of the target accounts before using the
	// StackSet action.
	//
	// The roles need to be assumable by by the `administrationRole`, and need to
	// have the permissions necessary to successfully create and modify the
	// resources that the subsequent CloudFormation deployments need.
	// Administrator permissions would be commonly granted to these, but if you can
	// scope the permissions down frome there you would be safer.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-self-managed.html
	//
	// Default: AWSCloudFormationStackSetExecutionRole.
	//
	ExecutionRoleName *string `field:"optional" json:"executionRoleName" yaml:"executionRoleName"`
}

