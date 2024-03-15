package awsssm


// Properties for defining a `CfnAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   cfnAssociationProps := &CfnAssociationProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ApplyOnlyAtCronInterval: jsii.Boolean(false),
//   	AssociationName: jsii.String("associationName"),
//   	AutomationTargetParameterName: jsii.String("automationTargetParameterName"),
//   	CalendarNames: []*string{
//   		jsii.String("calendarNames"),
//   	},
//   	ComplianceSeverity: jsii.String("complianceSeverity"),
//   	DocumentVersion: jsii.String("documentVersion"),
//   	InstanceId: jsii.String("instanceId"),
//   	MaxConcurrency: jsii.String("maxConcurrency"),
//   	MaxErrors: jsii.String("maxErrors"),
//   	OutputLocation: &InstanceAssociationOutputLocationProperty{
//   		S3Location: &S3OutputLocationProperty{
//   			OutputS3BucketName: jsii.String("outputS3BucketName"),
//   			OutputS3KeyPrefix: jsii.String("outputS3KeyPrefix"),
//   			OutputS3Region: jsii.String("outputS3Region"),
//   		},
//   	},
//   	Parameters: parameters,
//   	ScheduleExpression: jsii.String("scheduleExpression"),
//   	ScheduleOffset: jsii.Number(123),
//   	SyncCompliance: jsii.String("syncCompliance"),
//   	Targets: []interface{}{
//   		&TargetProperty{
//   			Key: jsii.String("key"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	WaitForSuccessTimeoutSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html
//
type CfnAssociationProps struct {
	// The name of the SSM document that contains the configuration information for the instance.
	//
	// You can specify `Command` or `Automation` documents. The documents can be AWS -predefined documents, documents you created, or a document that is shared with you from another account. For SSM documents that are shared with you from other AWS accounts , you must specify the complete SSM document ARN, in the following format:
	//
	// `arn:partition:ssm:region:account-id:document/document-name`
	//
	// For example: `arn:aws:ssm:us-east-2:12345678912:document/My-Shared-Document`
	//
	// For AWS -predefined documents and SSM documents you created in your account, you only need to specify the document name. For example, `AWS -ApplyPatchBaseline` or `My-Document` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// By default, when you create a new association, the system runs it immediately after it is created and then according to the schedule you specified.
	//
	// Specify this option if you don't want an association to run immediately after you create it. This parameter is not supported for rate expressions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-applyonlyatcroninterval
	//
	ApplyOnlyAtCronInterval interface{} `field:"optional" json:"applyOnlyAtCronInterval" yaml:"applyOnlyAtCronInterval"`
	// Specify a descriptive name for the association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-associationname
	//
	AssociationName *string `field:"optional" json:"associationName" yaml:"associationName"`
	// Choose the parameter that will define how your automation will branch out.
	//
	// This target is required for associations that use an Automation runbook and target resources by using rate controls. Automation is a capability of AWS Systems Manager .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-automationtargetparametername
	//
	AutomationTargetParameterName *string `field:"optional" json:"automationTargetParameterName" yaml:"automationTargetParameterName"`
	// The names or Amazon Resource Names (ARNs) of the Change Calendar type documents your associations are gated under.
	//
	// The associations only run when that Change Calendar is open. For more information, see [AWS Systems Manager Change Calendar](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-change-calendar) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-calendarnames
	//
	CalendarNames *[]*string `field:"optional" json:"calendarNames" yaml:"calendarNames"`
	// The severity level that is assigned to the association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-complianceseverity
	//
	ComplianceSeverity *string `field:"optional" json:"complianceSeverity" yaml:"complianceSeverity"`
	// The version of the SSM document to associate with the target.
	//
	// > Note the following important information.
	// >
	// > - State Manager doesn't support running associations that use a new version of a document if that document is shared from another account. State Manager always runs the `default` version of a document if shared from another account, even though the Systems Manager console shows that a new version was processed. If you want to run an association using a new version of a document shared form another account, you must set the document version to `default` .
	// > - `DocumentVersion` is not valid for documents owned by AWS , such as `AWS-RunPatchBaseline` or `AWS-UpdateSSMAgent` . If you specify `DocumentVersion` for an AWS document, the system returns the following error: "Error occurred during operation 'CreateAssociation'." (RequestToken: <token>, HandlerErrorCode: GeneralServiceException).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-documentversion
	//
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
	// The ID of the instance that the SSM document is associated with.
	//
	// You must specify the `InstanceId` or `Targets` property.
	//
	// > `InstanceId` has been deprecated. To specify an instance ID for an association, use the `Targets` parameter. If you use the parameter `InstanceId` , you cannot use the parameters `AssociationName` , `DocumentVersion` , `MaxErrors` , `MaxConcurrency` , `OutputLocation` , or `ScheduleExpression` . To use these parameters, you must use the `Targets` parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-instanceid
	//
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// The maximum number of targets allowed to run the association at the same time.
	//
	// You can specify a number, for example 10, or a percentage of the target set, for example 10%. The default value is 100%, which means all targets run the association at the same time.
	//
	// If a new managed node starts and attempts to run an association while Systems Manager is running `MaxConcurrency` associations, the association is allowed to run. During the next association interval, the new managed node will process its association within the limit specified for `MaxConcurrency` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-maxconcurrency
	//
	MaxConcurrency *string `field:"optional" json:"maxConcurrency" yaml:"maxConcurrency"`
	// The number of errors that are allowed before the system stops sending requests to run the association on additional targets.
	//
	// You can specify either an absolute number of errors, for example 10, or a percentage of the target set, for example 10%. If you specify 3, for example, the system stops sending requests when the fourth error is received. If you specify 0, then the system stops sending requests after the first error is returned. If you run an association on 50 managed nodes and set `MaxError` to 10%, then the system stops sending the request when the sixth error is received.
	//
	// Executions that are already running an association when `MaxErrors` is reached are allowed to complete, but some of these executions may fail as well. If you need to ensure that there won't be more than max-errors failed executions, set `MaxConcurrency` to 1 so that executions proceed one at a time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-maxerrors
	//
	MaxErrors *string `field:"optional" json:"maxErrors" yaml:"maxErrors"`
	// An Amazon Simple Storage Service (Amazon S3) bucket where you want to store the output details of the request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-outputlocation
	//
	OutputLocation interface{} `field:"optional" json:"outputLocation" yaml:"outputLocation"`
	// The parameters for the runtime configuration of the document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// A cron expression that specifies a schedule when the association runs.
	//
	// The schedule runs in Coordinated Universal Time (UTC).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-scheduleexpression
	//
	ScheduleExpression *string `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
	// Number of days to wait after the scheduled day to run an association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-scheduleoffset
	//
	ScheduleOffset *float64 `field:"optional" json:"scheduleOffset" yaml:"scheduleOffset"`
	// The mode for generating association compliance.
	//
	// You can specify `AUTO` or `MANUAL` . In `AUTO` mode, the system uses the status of the association execution to determine the compliance status. If the association execution runs successfully, then the association is `COMPLIANT` . If the association execution doesn't run successfully, the association is `NON-COMPLIANT` .
	//
	// In `MANUAL` mode, you must specify the `AssociationId` as a parameter for the `PutComplianceItems` API action. In this case, compliance data is not managed by State Manager. It is managed by your direct call to the `PutComplianceItems` API action.
	//
	// By default, all associations use `AUTO` mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-synccompliance
	//
	SyncCompliance *string `field:"optional" json:"syncCompliance" yaml:"syncCompliance"`
	// The targets for the association.
	//
	// You must specify the `InstanceId` or `Targets` property. You can target all instances in an AWS account by specifying t he `InstanceIds` key with a value of `*` .
	//
	// Supported formats include the following.
	//
	// - `Key=InstanceIds,Values=<instance-id-1>,<instance-id-2>,<instance-id-3>`
	// - `Key=tag-key,Values=<my-tag-key-1>,<my-tag-key-2>`
	//
	// To view a JSON and a YAML example that targets all instances, see "Create an association for all managed instances in an AWS account " on the Examples page.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-targets
	//
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
	// The number of seconds the service should wait for the association status to show "Success" before proceeding with the stack execution.
	//
	// If the association status doesn't show "Success" after the specified number of seconds, then stack creation fails.
	//
	// > When you specify a value for the `WaitForSuccessTimeoutSeconds` , [drift detection](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift.html) for your AWS CloudFormation stack’s configuration might yield inaccurate results. If drift detection is important in your scenario, we recommend that you don’t include `WaitForSuccessTimeoutSeconds` in your template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-waitforsuccesstimeoutseconds
	//
	WaitForSuccessTimeoutSeconds *float64 `field:"optional" json:"waitForSuccessTimeoutSeconds" yaml:"waitForSuccessTimeoutSeconds"`
}

