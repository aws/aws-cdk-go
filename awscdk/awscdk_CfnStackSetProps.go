// Version 2 of the AWS Cloud Development Kit library
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
//   cfnStackSetProps := &cfnStackSetProps{
//   	permissionModel: jsii.String("permissionModel"),
//   	stackSetName: jsii.String("stackSetName"),
//
//   	// the properties below are optional
//   	administrationRoleArn: jsii.String("administrationRoleArn"),
//   	autoDeployment: &autoDeploymentProperty{
//   		enabled: jsii.Boolean(false),
//   		retainStacksOnAccountRemoval: jsii.Boolean(false),
//   	},
//   	callAs: jsii.String("callAs"),
//   	capabilities: []*string{
//   		jsii.String("capabilities"),
//   	},
//   	description: jsii.String("description"),
//   	executionRoleName: jsii.String("executionRoleName"),
//   	managedExecution: managedExecution,
//   	operationPreferences: &operationPreferencesProperty{
//   		failureToleranceCount: jsii.Number(123),
//   		failureTolerancePercentage: jsii.Number(123),
//   		maxConcurrentCount: jsii.Number(123),
//   		maxConcurrentPercentage: jsii.Number(123),
//   		regionConcurrencyType: jsii.String("regionConcurrencyType"),
//   		regionOrder: []*string{
//   			jsii.String("regionOrder"),
//   		},
//   	},
//   	parameters: []interface{}{
//   		&parameterProperty{
//   			parameterKey: jsii.String("parameterKey"),
//   			parameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	stackInstancesGroup: []interface{}{
//   		&stackInstancesProperty{
//   			deploymentTargets: &deploymentTargetsProperty{
//   				accountFilterType: jsii.String("accountFilterType"),
//   				accounts: []*string{
//   					jsii.String("accounts"),
//   				},
//   				organizationalUnitIds: []*string{
//   					jsii.String("organizationalUnitIds"),
//   				},
//   			},
//   			regions: []*string{
//   				jsii.String("regions"),
//   			},
//
//   			// the properties below are optional
//   			parameterOverrides: []interface{}{
//   				&parameterProperty{
//   					parameterKey: jsii.String("parameterKey"),
//   					parameterValue: jsii.String("parameterValue"),
//   				},
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	templateBody: jsii.String("templateBody"),
//   	templateUrl: jsii.String("templateUrl"),
//   }
//
type CfnStackSetProps struct {
	// Describes how the IAM roles required for stack set operations are created.
	//
	// - With `SELF_MANAGED` permissions, you must create the administrator and execution roles required to deploy to target accounts. For more information, see [Grant Self-Managed Stack Set Permissions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-self-managed.html) .
	// - With `SERVICE_MANAGED` permissions, StackSets automatically creates the IAM roles required to deploy to accounts managed by AWS Organizations . For more information, see [Grant Service-Managed Stack Set Permissions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-service-managed.html) .
	//
	// *Allowed Values* : `SERVICE_MANAGED` | `SELF_MANAGED`
	//
	// > The `PermissionModel` property is required.
	PermissionModel *string `field:"required" json:"permissionModel" yaml:"permissionModel"`
	// The name to associate with the stack set.
	//
	// The name must be unique in the Region where you create your stack set.
	//
	// *Maximum* : `128`
	//
	// *Pattern* : `^[a-zA-Z][a-zA-Z0-9-]{0,127}$`
	//
	// > The `StackSetName` property is required.
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
	AdministrationRoleArn *string `field:"optional" json:"administrationRoleArn" yaml:"administrationRoleArn"`
	// [ `Service-managed` permissions] Describes whether StackSets automatically deploys to AWS Organizations accounts that are added to a target organization or organizational unit (OU).
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
	CallAs *string `field:"optional" json:"callAs" yaml:"callAs"`
	// The capabilities that are allowed in the stack set.
	//
	// Some stack set templates might include resources that can affect permissions in your AWS account —for example, by creating new AWS Identity and Access Management ( IAM ) users. For more information, see [Acknowledging IAM Resources in AWS CloudFormation Templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#capabilities) .
	Capabilities *[]*string `field:"optional" json:"capabilities" yaml:"capabilities"`
	// A description of the stack set.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `1024`.
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
	ManagedExecution interface{} `field:"optional" json:"managedExecution" yaml:"managedExecution"`
	// The user-specified preferences for how AWS CloudFormation performs a stack set operation.
	OperationPreferences interface{} `field:"optional" json:"operationPreferences" yaml:"operationPreferences"`
	// The input parameters for the stack set template.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// A group of stack instances with parameters in some specific accounts and Regions.
	StackInstancesGroup interface{} `field:"optional" json:"stackInstancesGroup" yaml:"stackInstancesGroup"`
	// The key-value pairs to associate with this stack set and the stacks created from it.
	//
	// AWS CloudFormation also propagates these tags to supported resources that are created in the stacks. A maximum number of 50 tags can be specified.
	Tags *[]*CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The structure that contains the template body, with a minimum length of 1 byte and a maximum length of 51,200 bytes.
	//
	// You must include either `TemplateURL` or `TemplateBody` in a StackSet, but you can't use both. Dynamic references in the `TemplateBody` may not work correctly in all cases. It's recommended to pass templates containing dynamic references through `TemplateUrl` instead.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `51200`.
	TemplateBody *string `field:"optional" json:"templateBody" yaml:"templateBody"`
	// Location of file containing the template body.
	//
	// The URL must point to a template (max size: 460,800 bytes) that's located in an Amazon S3 bucket.
	//
	// You must include either `TemplateURL` or `TemplateBody` in a StackSet, but you can't use both.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `1024`.
	TemplateUrl *string `field:"optional" json:"templateUrl" yaml:"templateUrl"`
}

