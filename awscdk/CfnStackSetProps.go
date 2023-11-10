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
	// Describes how the IAM roles required for stack set operations are created.
	//
	// - With `SELF_MANAGED` permissions, you must create the administrator and execution roles required to deploy to target accounts. For more information, see [Grant Self-Managed Stack Set Permissions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-self-managed.html) .
	// - With `SERVICE_MANAGED` permissions, StackSets automatically creates the IAM roles required to deploy to accounts managed by AWS Organizations .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-permissionmodel
	//
	PermissionModel *string `field:"required" json:"permissionModel" yaml:"permissionModel"`
	// The name to associate with the stack set.
	//
	// The name must be unique in the Region where you create your stack set.
	//
	// > The `StackSetName` property is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-stacksetname
	//
	StackSetName *string `field:"required" json:"stackSetName" yaml:"stackSetName"`
	// The Amazon Resource Number (ARN) of the IAM role to use to create this stack set.
	//
	// Specify an IAM role only if you are using customized administrator roles to control which users or groups can manage specific stack sets within the same administrator account.
	//
	// Use customized administrator roles to control which users or groups can manage specific stack sets within the same administrator account. For more information, see [Prerequisites: Granting Permissions for Stack Set Operations](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs.html) in the *AWS CloudFormation User Guide* .
	//
	// *Minimum* : `20`
	//
	// *Maximum* : `2048`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-administrationrolearn
	//
	AdministrationRoleArn *string `field:"optional" json:"administrationRoleArn" yaml:"administrationRoleArn"`
	// [ `Service-managed` permissions] Describes whether StackSets automatically deploys to AWS Organizations accounts that are added to a target organization or organizational unit (OU).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-autodeployment
	//
	AutoDeployment interface{} `field:"optional" json:"autoDeployment" yaml:"autoDeployment"`
	// [Service-managed permissions] Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account.
	//
	// By default, `SELF` is specified. Use `SELF` for stack sets with self-managed permissions.
	//
	// - To create a stack set with service-managed permissions while signed in to the management account, specify `SELF` .
	// - To create a stack set with service-managed permissions while signed in to a delegated administrator account, specify `DELEGATED_ADMIN` .
	//
	// Your AWS account must be registered as a delegated admin in the management account. For more information, see [Register a delegated administrator](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-delegated-admin.html) in the *AWS CloudFormation User Guide* .
	//
	// Stack sets with service-managed permissions are created in the management account, including stack sets that are created by delegated administrators.
	//
	// *Valid Values* : `SELF` | `DELEGATED_ADMIN`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-callas
	//
	CallAs *string `field:"optional" json:"callAs" yaml:"callAs"`
	// The capabilities that are allowed in the stack set.
	//
	// Some stack set templates might include resources that can affect permissions in your AWS account —for example, by creating new AWS Identity and Access Management ( IAM ) users. For more information, see [Acknowledging IAM Resources in AWS CloudFormation Templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#capabilities) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-capabilities
	//
	Capabilities *[]*string `field:"optional" json:"capabilities" yaml:"capabilities"`
	// A description of the stack set.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `1024`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the IAM execution role to use to create the stack set.
	//
	// If you don't specify an execution role, AWS CloudFormation uses the `AWSCloudFormationStackSetExecutionRole` role for the stack set operation.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `64`
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
	// > You can't modify your stack set's execution configuration while there are running or queued operations for that stack set.
	//
	// When inactive (default), StackSets performs one operation at a time in request order.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-managedexecution
	//
	ManagedExecution interface{} `field:"optional" json:"managedExecution" yaml:"managedExecution"`
	// The user-specified preferences for how AWS CloudFormation performs a stack set operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-operationpreferences
	//
	OperationPreferences interface{} `field:"optional" json:"operationPreferences" yaml:"operationPreferences"`
	// The input parameters for the stack set template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// A group of stack instances with parameters in some specific accounts and Regions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-stackinstancesgroup
	//
	StackInstancesGroup interface{} `field:"optional" json:"stackInstancesGroup" yaml:"stackInstancesGroup"`
	// Key-value pairs to associate with this stack.
	//
	// AWS CloudFormation also propagates these tags to supported resources in the stack. You can specify a maximum number of 50 tags.
	//
	// If you don't specify this parameter, AWS CloudFormation doesn't modify the stack's tags. If you specify an empty value, AWS CloudFormation removes all associated tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-tags
	//
	Tags *[]*CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The structure that contains the template body, with a minimum length of 1 byte and a maximum length of 51,200 bytes.
	//
	// You must include either `TemplateURL` or `TemplateBody` in a StackSet, but you can't use both. Dynamic references in the `TemplateBody` may not work correctly in all cases. It's recommended to pass templates containing dynamic references through `TemplateUrl` instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-templatebody
	//
	TemplateBody *string `field:"optional" json:"templateBody" yaml:"templateBody"`
	// Location of file containing the template body.
	//
	// The URL must point to a template that's located in an Amazon S3 bucket or a Systems Manager document. For more information, go to [Template Anatomy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html) in the AWS CloudFormation User Guide.
	//
	// Conditional: You must specify only one of the following parameters: `TemplateBody` , `TemplateURL` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-templateurl
	//
	TemplateUrl *string `field:"optional" json:"templateUrl" yaml:"templateUrl"`
}

