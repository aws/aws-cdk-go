package customresources


// Options for the auto-generation of policies based on the configured SDK calls.
//
// Example:
//   getParameter := cr.NewAwsCustomResource(this, jsii.String("GetParameter"), &AwsCustomResourceProps{
//   	OnUpdate: &AwsSdkCall{
//   		Service: jsii.String("SSM"),
//   		Action: jsii.String("GetParameter"),
//   		Parameters: map[string]interface{}{
//   			"Name": jsii.String("my-parameter"),
//   			"WithDecryption": jsii.Boolean(true),
//   		},
//   		PhysicalResourceId: cr.PhysicalResourceId_Of(date.now().toString()),
//   		Logging: cr.Logging_WithDataHidden(),
//   	},
//   	Policy: cr.AwsCustomResourcePolicy_FromSdkCalls(&SdkCallsPolicyOptions{
//   		Resources: cr.AwsCustomResourcePolicy_ANY_RESOURCE(),
//   	}),
//   })
//
type SdkCallsPolicyOptions struct {
	// The resources that the calls will have access to.
	//
	// It is best to use specific resource ARN's when possible. However, you can also use `AwsCustomResourcePolicy.ANY_RESOURCE`
	// to allow access to all resources. For example, when `onCreate` is used to create a resource which you don't
	// know the physical name of in advance.
	//
	// Note that will apply to ALL SDK calls.
	Resources *[]*string `field:"required" json:"resources" yaml:"resources"`
}

