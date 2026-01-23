package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskms"
)

// Properties for a LogGroup.
//
// Example:
//   var myRole Role
//
//   cr.NewAwsCustomResource(this, jsii.String("Customized"), &AwsCustomResourceProps{
//   	Role: myRole,
//   	 // must be assumable by the `lambda.amazonaws.com` service principal
//   	Timeout: awscdk.Duration_Minutes(jsii.Number(10)),
//   	 // defaults to 2 minutes
//   	MemorySize: jsii.Number(1025),
//   	 // defaults to 512 if installLatestAwsSdk is true
//   	LogGroup: logs.NewLogGroup(this, jsii.String("AwsCustomResourceLogs"), &LogGroupProps{
//   		Retention: logs.RetentionDays_ONE_DAY,
//   	}),
//   	FunctionName: jsii.String("my-custom-name"),
//   	 // defaults to a CloudFormation generated name
//   	RemovalPolicy: awscdk.RemovalPolicy_RETAIN,
//   	 // defaults to `RemovalPolicy.DESTROY`
//   	Policy: cr.AwsCustomResourcePolicy_FromSdkCalls(&SdkCallsPolicyOptions{
//   		Resources: cr.AwsCustomResourcePolicy_ANY_RESOURCE(),
//   	}),
//   })
//
type LogGroupProps struct {
	// Data Protection Policy for this log group.
	// Default: - no data protection policy.
	//
	DataProtectionPolicy DataProtectionPolicy `field:"optional" json:"dataProtectionPolicy" yaml:"dataProtectionPolicy"`
	// Indicates whether deletion protection is enabled for this log group.
	//
	// When enabled,
	// deletion protection blocks all deletion operations until it is explicitly disabled.
	// Default: false.
	//
	DeletionProtectionEnabled *bool `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// The KMS customer managed key to encrypt the log group with.
	// Default: Server-side encryption managed by the CloudWatch Logs service.
	//
	EncryptionKey interfacesawskms.IKeyRef `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Field Index Policies for this log group.
	// Default: - no field index policies for this log group.
	//
	FieldIndexPolicies *[]FieldIndexPolicy `field:"optional" json:"fieldIndexPolicies" yaml:"fieldIndexPolicies"`
	// The class of the log group. Possible values are: STANDARD and INFREQUENT_ACCESS.
	//
	// INFREQUENT_ACCESS class provides customers a cost-effective way to consolidate
	// logs which supports querying using Logs Insights. The logGroupClass property cannot
	// be changed once the log group is created.
	// Default: LogGroupClass.STANDARD
	//
	LogGroupClass LogGroupClass `field:"optional" json:"logGroupClass" yaml:"logGroupClass"`
	// Name of the log group.
	// Default: Automatically generated.
	//
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// Determine the removal policy of this log group.
	//
	// Normally you want to retain the log group so you can diagnose issues
	// from logs even after a deployment that no longer includes the log group.
	// In that case, use the normal date-based retention policy to age out your
	// logs.
	// Default: RemovalPolicy.Retain
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// How long, in days, the log contents will be retained.
	//
	// To retain all logs, set this value to RetentionDays.INFINITE.
	// Default: RetentionDays.TWO_YEARS
	//
	Retention RetentionDays `field:"optional" json:"retention" yaml:"retention"`
}

