package awsec2


// Base properties for an Amazon Linux SSM Parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var userData userData
//
//   amazonLinuxImageSsmParameterBaseProps := &AmazonLinuxImageSsmParameterBaseProps{
//   	ParameterName: jsii.String("parameterName"),
//
//   	// the properties below are optional
//   	CachedInContext: jsii.Boolean(false),
//   	UserData: userData,
//   }
//
type AmazonLinuxImageSsmParameterBaseProps struct {
	// Whether the AMI ID is cached to be stable between deployments.
	//
	// By default, the newest image is used on each deployment. This will cause
	// instances to be replaced whenever a new version is released, and may cause
	// downtime if there aren't enough running instances in the AutoScalingGroup
	// to reschedule the tasks on.
	//
	// If set to true, the AMI ID will be cached in `cdk.context.json` and the
	// same value will be used on future runs. Your instances will not be replaced
	// but your AMI version will grow old over time. To refresh the AMI lookup,
	// you will have to evict the value from the cache using the `cdk context`
	// command. See https://docs.aws.amazon.com/cdk/latest/guide/context.html for
	// more information.
	//
	// Can not be set to `true` in environment-agnostic stacks.
	// Default: false.
	//
	CachedInContext *bool `field:"optional" json:"cachedInContext" yaml:"cachedInContext"`
	// Initial user data.
	// Default: - Empty UserData for Linux machines.
	//
	UserData UserData `field:"optional" json:"userData" yaml:"userData"`
	// The name of the SSM parameter that contains the AMI value.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
}

