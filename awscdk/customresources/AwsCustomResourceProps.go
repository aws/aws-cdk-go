package customresources

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Properties for AwsCustomResource.
//
// Note that at least onCreate, onUpdate or onDelete must be specified.
//
// Example:
//   getParameter := cr.NewAwsCustomResource(this, jsii.String("GetParameter"), &AwsCustomResourceProps{
//   	OnUpdate: &AwsSdkCall{
//   		 // will also be called for a CREATE event
//   		Service: jsii.String("SSM"),
//   		Action: jsii.String("GetParameter"),
//   		Parameters: map[string]interface{}{
//   			"Name": jsii.String("my-parameter"),
//   			"WithDecryption": jsii.Boolean(true),
//   		},
//   		PhysicalResourceId: cr.PhysicalResourceId_Of(date.now().toString()),
//   	},
//   	Policy: cr.AwsCustomResourcePolicy_FromSdkCalls(&SdkCallsPolicyOptions{
//   		Resources: cr.AwsCustomResourcePolicy_ANY_RESOURCE(),
//   	}),
//   })
//
//   // Use the value in another construct with
//   getParameter.GetResponseField(jsii.String("Parameter.Value"))
//
type AwsCustomResourceProps struct {
	// A name for the singleton Lambda function implementing this custom resource.
	//
	// The function name will remain the same after the first AwsCustomResource is created in a stack.
	// Default: - AWS CloudFormation generates a unique physical ID and uses that
	// ID for the function's name. For more information, see Name Type.
	//
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// Whether to install the latest AWS SDK v3.
	//
	// If not specified, this uses whatever JavaScript SDK version is the default in
	// AWS Lambda at the time of execution.
	//
	// Otherwise, installs the latest version from 'npmjs.com'. The installation takes
	// around 60 seconds and requires internet connectivity.
	//
	// The default can be controlled using the context key
	// `@aws-cdk/customresources:installLatestAwsSdkDefault` is.
	// Default: - The value of `@aws-cdk/customresources:installLatestAwsSdkDefault`, otherwise `true`.
	//
	InstallLatestAwsSdk *bool `field:"optional" json:"installLatestAwsSdk" yaml:"installLatestAwsSdk"`
	// The Log Group used for logging of events emitted by the custom resource's lambda function.
	//
	// Providing a user-controlled log group was rolled out to commercial regions on 2023-11-16.
	// If you are deploying to another type of region, please check regional availability first.
	// Default: - a default log group created by AWS Lambda.
	//
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// The number of days log events of the singleton Lambda function implementing this custom resource are kept in CloudWatch Logs.
	//
	// This is a legacy API and we strongly recommend you migrate to `logGroup` if you can.
	// `logGroup` allows you to create a fully customizable log group and instruct the Lambda function to send logs to it.
	// Default: logs.RetentionDays.INFINITE
	//
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// The memory size for the singleton Lambda function implementing this custom resource.
	// Default: 512 mega in case if installLatestAwsSdk is false.
	//
	MemorySize *float64 `field:"optional" json:"memorySize" yaml:"memorySize"`
	// The AWS SDK call to make when the resource is created.
	// Default: - the call when the resource is updated.
	//
	OnCreate *AwsSdkCall `field:"optional" json:"onCreate" yaml:"onCreate"`
	// The AWS SDK call to make when the resource is deleted.
	// Default: - no call.
	//
	OnDelete *AwsSdkCall `field:"optional" json:"onDelete" yaml:"onDelete"`
	// The AWS SDK call to make when the resource is updated.
	// Default: - no call.
	//
	OnUpdate *AwsSdkCall `field:"optional" json:"onUpdate" yaml:"onUpdate"`
	// The policy that will be added to the execution role of the Lambda function implementing this custom resource provider.
	//
	// The custom resource also implements `iam.IGrantable`, making it possible
	// to use the `grantXxx()` methods.
	//
	// As this custom resource uses a singleton Lambda function, it's important
	// to note the that function's role will eventually accumulate the
	// permissions/grants from all resources.
	//
	// Note that a policy must be specified if `role` is not provided, as
	// by default a new role is created which requires policy changes to access
	// resources.
	// See: Policy.fromSdkCalls
	//
	// Default: - no policy added.
	//
	Policy AwsCustomResourcePolicy `field:"optional" json:"policy" yaml:"policy"`
	// The policy to apply when this resource is removed from the application.
	// Default: cdk.RemovalPolicy.Destroy
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Cloudformation Resource type.
	// Default: - Custom::AWS.
	//
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// The execution role for the singleton Lambda function implementing this custom resource provider.
	//
	// This role will apply to all `AwsCustomResource`
	// instances in the stack. The role must be assumable by the
	// `lambda.amazonaws.com` service principal.
	// Default: - a new role is created.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The maximum time that can elapse before a custom resource operation times out.
	//
	// You should not need to set this property. It is intended to allow quick turnaround
	// even if the implementor of the custom resource forgets to include a `try/catch`.
	// We have included the `try/catch`, and AWS service calls usually do not take an hour
	// to complete.
	//
	// The value must be between 1 second and 3600 seconds.
	// Default: Duration.seconds(3600)
	//
	ServiceTimeout awscdk.Duration `field:"optional" json:"serviceTimeout" yaml:"serviceTimeout"`
	// The timeout for the singleton Lambda function implementing this custom resource.
	// Default: Duration.minutes(2)
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The vpc to provision the lambda function in.
	// Default: - the function is not provisioned inside a vpc.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Which subnets from the VPC to place the lambda function in.
	//
	// Only used if 'vpc' is supplied. Note: internet access for Lambdas
	// requires a NAT gateway, so picking Public subnets is not allowed.
	// Default: - the Vpc default strategy if not specified.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

