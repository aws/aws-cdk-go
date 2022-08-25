package customresources


// Options for the auto-generation of policies based on the configured SDK calls.
//
// Example:
//   getParameter := cr.NewAwsCustomResource(this, jsii.String("GetParameter"), &awsCustomResourceProps{
//   	onUpdate: &awsSdkCall{
//   		 // will also be called for a CREATE event
//   		service: jsii.String("SSM"),
//   		action: jsii.String("getParameter"),
//   		parameters: map[string]interface{}{
//   			"Name": jsii.String("my-parameter"),
//   			"WithDecryption": jsii.Boolean(true),
//   		},
//   		physicalResourceId: cr.physicalResourceId.of(date.now().toString()),
//   	},
//   	policy: cr.awsCustomResourcePolicy.fromSdkCalls(&sdkCallsPolicyOptions{
//   		resources: cr.*awsCustomResourcePolicy_ANY_RESOURCE(),
//   	}),
//   })
//
//   // Use the value in another construct with
//   getParameter.getResponseField(jsii.String("Parameter.Value"))
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

