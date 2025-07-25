package awscdk


// Properties for defining a `CfnStackSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var managedExecution interface{}
//
//   cfnStackSetProps := &CfnStackSetProps{
//   	PermissionModel: jsii.String("permissionModel"),
//   	StackSetName: jsii.String("stackSetName"),
//
//   	// the properties below are optional
//   	AdministrationRoleArn: jsii.String("administrationRoleArn"),
//   	AutoDeployment: &AutoDeploymentProperty{
//   		Enabled: jsii.Boolean(false),
//   		RetainStacksOnAccountRemoval: jsii.Boolean(false),
//   	},
//   	CallAs: jsii.String("callAs"),
//   	Capabilities: []*string{
//   		jsii.String("capabilities"),
//   	},
//   	Description: jsii.String("description"),
//   	ExecutionRoleName: jsii.String("executionRoleName"),
//   	ManagedExecution: managedExecution,
//   	OperationPreferences: &OperationPreferencesProperty{
//   		ConcurrencyMode: jsii.String("concurrencyMode"),
//   		FailureToleranceCount: jsii.Number(123),
//   		FailureTolerancePercentage: jsii.Number(123),
//   		MaxConcurrentCount: jsii.Number(123),
//   		MaxConcurrentPercentage: jsii.Number(123),
//   		RegionConcurrencyType: jsii.String("regionConcurrencyType"),
//   		RegionOrder: []*string{
//   			jsii.String("regionOrder"),
//   		},
//   	},
//   	Parameters: []interface{}{
//   		&ParameterProperty{
//   			ParameterKey: jsii.String("parameterKey"),
//   			ParameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	StackInstancesGroup: []interface{}{
//   		&StackInstancesProperty{
//   			DeploymentTargets: &DeploymentTargetsProperty{
//   				AccountFilterType: jsii.String("accountFilterType"),
//   				Accounts: []*string{
//   					jsii.String("accounts"),
//   				},
//   				AccountsUrl: jsii.String("accountsUrl"),
//   				OrganizationalUnitIds: []*string{
//   					jsii.String("organizationalUnitIds"),
//   				},
//   			},
//   			Regions: []*string{
//   				jsii.String("regions"),
//   			},
//
//   			// the properties below are optional
//   			ParameterOverrides: []interface{}{
//   				&ParameterProperty{
//   					ParameterKey: jsii.String("parameterKey"),
//   					ParameterValue: jsii.String("parameterValue"),
//   				},
//   			},
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TemplateBody: jsii.String("templateBody"),
//   	TemplateUrl: jsii.String("templateUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html
//
type CfnStackSetProps struct {
	// Describes how the IAM roles required for StackSet operations are created.
	//
	// - With `SELF_MANAGED` permissions, you must create the administrator and execution roles required to deploy to target accounts. For more information, see [Grant self-managed permissions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-self-managed.html) in the *AWS CloudFormation User Guide* .
	// - With `SERVICE_MANAGED` permissions, StackSets automatically creates the IAM roles required to deploy to accounts managed by AWS Organizations . For more information, see [Activate trusted access for StackSets with AWS Organizations](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-activate-trusted-access.html) in the *AWS CloudFormation User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-permissionmodel
	//
	PermissionModel *string `field:"required" json:"permissionModel" yaml:"permissionModel"`
	// The name to associate with the StackSet.
	//
	// The name must be unique in the Region where you create your StackSet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-stacksetname
	//
	StackSetName *string `field:"required" json:"stackSetName" yaml:"stackSetName"`
	// The Amazon Resource Number (ARN) of the IAM role to use to create this StackSet.
	//
	// Specify an IAM role only if you are using customized administrator roles to control which users or groups can manage specific StackSets within the same administrator account.
	//
	// Use customized administrator roles to control which users or groups can manage specific StackSets within the same administrator account. For more information, see [Grant self-managed permissions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-self-managed.html) in the *AWS CloudFormation User Guide* .
	//
	// Valid only if the permissions model is `SELF_MANAGED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-administrationrolearn
	//
	AdministrationRoleArn *string `field:"optional" json:"administrationRoleArn" yaml:"administrationRoleArn"`
	// Describes whether StackSets automatically deploys to AWS Organizations accounts that are added to a target organization or organizational unit (OU).
	//
	// For more information, see [Enable or disable automatic deployments for StackSets in AWS Organizations](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-manage-auto-deployment.html) in the *AWS CloudFormation User Guide* .
	//
	// Required if the permissions model is `SERVICE_MANAGED` . (Not used with self-managed permissions.)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-autodeployment
	//
	AutoDeployment interface{} `field:"optional" json:"autoDeployment" yaml:"autoDeployment"`
	// Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account.
	//
	// By default, `SELF` is specified. Use `SELF` for StackSets with self-managed permissions.
	//
	// - To create a StackSet with service-managed permissions while signed in to the management account, specify `SELF` .
	// - To create a StackSet with service-managed permissions while signed in to a delegated administrator account, specify `DELEGATED_ADMIN` .
	//
	// Your AWS account must be registered as a delegated admin in the management account. For more information, see [Register a delegated administrator](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-delegated-admin.html) in the *AWS CloudFormation User Guide* .
	//
	// StackSets with service-managed permissions are created in the management account, including StackSets that are created by delegated administrators.
	//
	// Valid only if the permissions model is `SERVICE_MANAGED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-callas
	//
	CallAs *string `field:"optional" json:"callAs" yaml:"callAs"`
	// The capabilities that are allowed in the StackSet.
	//
	// Some StackSet templates might include resources that can affect permissions in your AWS account —for example, by creating new IAM users. For more information, see [Acknowledging IAM resources in CloudFormation templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/control-access-with-iam.html#using-iam-capabilities) in the *AWS CloudFormation User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-capabilities
	//
	Capabilities *[]*string `field:"optional" json:"capabilities" yaml:"capabilities"`
	// A description of the StackSet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the IAM execution role to use to create the StackSet.
	//
	// If you don't specify an execution role, CloudFormation uses the `AWSCloudFormationStackSetExecutionRole` role for the StackSet operation.
	//
	// Valid only if the permissions model is `SELF_MANAGED` .
	//
	// *Pattern* : `[a-zA-Z_0-9+=,.@-]+`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-executionrolename
	//
	ExecutionRoleName *string `field:"optional" json:"executionRoleName" yaml:"executionRoleName"`
	// Describes whether StackSets performs non-conflicting operations concurrently and queues conflicting operations.
	//
	// When active, StackSets performs non-conflicting operations concurrently and queues conflicting operations. After conflicting operations finish, StackSets starts queued operations in request order.
	//
	// > If there are already running or queued operations, StackSets queues all incoming operations even if they are non-conflicting.
	// >
	// > You can't modify your StackSet's execution configuration while there are running or queued operations for that StackSet.
	//
	// When inactive (default), StackSets performs one operation at a time in request order.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-managedexecution
	//
	ManagedExecution interface{} `field:"optional" json:"managedExecution" yaml:"managedExecution"`
	// The user-specified preferences for how CloudFormation performs a StackSet operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-operationpreferences
	//
	OperationPreferences interface{} `field:"optional" json:"operationPreferences" yaml:"operationPreferences"`
	// The input parameters for the StackSet template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// A group of stack instances with parameters in some specific accounts and Regions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-stackinstancesgroup
	//
	StackInstancesGroup interface{} `field:"optional" json:"stackInstancesGroup" yaml:"stackInstancesGroup"`
	// Key-value pairs to associate with this stack.
	//
	// CloudFormation also propagates these tags to supported resources in the stack. You can specify a maximum number of 50 tags.
	//
	// If you don't specify this parameter, CloudFormation doesn't modify the stack's tags. If you specify an empty value, CloudFormation removes all associated tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-tags
	//
	Tags *[]*CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The structure that contains the template body, with a minimum length of 1 byte and a maximum length of 51,200 bytes.
	//
	// You must include either `TemplateURL` or `TemplateBody` in a StackSet, but you can't use both. Dynamic references in the `TemplateBody` may not work correctly in all cases. It's recommended to pass templates that contain dynamic references through `TemplateUrl` instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-templatebody
	//
	TemplateBody *string `field:"optional" json:"templateBody" yaml:"templateBody"`
	// The URL of a file that contains the template body.
	//
	// The URL must point to a template (max size: 1 MB) that's located in an Amazon S3 bucket or a Systems Manager document. The location for an Amazon S3 bucket must start with `https://` .
	//
	// Conditional: You must specify only one of the following parameters: `TemplateBody` , `TemplateURL` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-templateurl
	//
	TemplateUrl *string `field:"optional" json:"templateUrl" yaml:"templateUrl"`
}

