package awscodepipelineactions


// Describes whether AWS CloudFormation StackSets automatically deploys to AWS Organizations accounts that are added to a target organization or organizational unit (OU).
// Experimental.
type StackSetOrganizationsAutoDeployment string

const (
	// StackSets automatically deploys additional stack instances to AWS Organizations accounts that are added to a target organization or organizational unit (OU) in the specified Regions.
	//
	// If an account is removed from a target organization or OU, AWS CloudFormation StackSets
	// deletes stack instances from the account in the specified Regions.
	// Experimental.
	StackSetOrganizationsAutoDeployment_ENABLED StackSetOrganizationsAutoDeployment = "ENABLED"
	// StackSets does not automatically deploy additional stack instances to AWS Organizations accounts that are added to a target organization or organizational unit (OU) in the specified Regions.
	// Experimental.
	StackSetOrganizationsAutoDeployment_DISABLED StackSetOrganizationsAutoDeployment = "DISABLED"
	// Stack resources are retained when an account is removed from a target organization or OU.
	// Experimental.
	StackSetOrganizationsAutoDeployment_ENABLED_WITH_STACK_RETENTION StackSetOrganizationsAutoDeployment = "ENABLED_WITH_STACK_RETENTION"
)

