package awscdk


// Properties to provide a Lambda-backed custom resource.
//
// Example:
//   stack := awscdk.Newstack()
//   durToken := awscdk.NewCfnParameter(stack, jsii.String("MyParameter"), &CfnParameterProps{
//   	Type: jsii.String("Number"),
//   	Default: jsii.Number(60),
//   })
//   awscdk.NewCustomResource(stack, jsii.String("MyCustomResource"), &CustomResourceProps{
//   	ServiceToken: jsii.String("MyServiceToken"),
//   	ServiceTimeout: awscdk.Duration_Seconds(durToken.valueAsNumber),
//   })
//
type CustomResourceProps struct {
	// The ARN of the provider which implements this custom resource type.
	//
	// You can implement a provider by listening to raw AWS CloudFormation events
	// and specify the ARN of an SNS topic (`topic.topicArn`) or the ARN of an AWS
	// Lambda function (`lambda.functionArn`) or use the CDK's custom [resource
	// provider framework] which makes it easier to implement robust providers.
	//
	// [resource provider framework]:
	// https://docs.aws.amazon.com/cdk/api/latest/docs/custom-resources-readme.html
	//
	// Provider framework:
	//
	// ```ts
	// // use the provider framework from aws-cdk/custom-resources:
	// const provider = new customresources.Provider(this, 'ResourceProvider', {
	//   onEventHandler,
	//   isCompleteHandler, // optional
	// });
	//
	// new CustomResource(this, 'MyResource', {
	//   serviceToken: provider.serviceToken,
	// });
	// ```
	//
	// AWS Lambda function (not recommended to use AWS Lambda Functions directly,
	// see the module README):
	//
	// ```ts
	// // invoke an AWS Lambda function when a lifecycle event occurs:
	// new CustomResource(this, 'MyResource', {
	//   serviceToken: myFunction.functionArn,
	// });
	// ```
	//
	// SNS topic (not recommended to use AWS Lambda Functions directly, see the
	// module README):
	//
	// ```ts
	// // publish lifecycle events to an SNS topic:
	// new CustomResource(this, 'MyResource', {
	//   serviceToken: myTopic.topicArn,
	// });
	// ```
	//
	// Maps to [ServiceToken](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-customresource.html#cfn-cloudformation-customresource-servicetoken) property for the `AWS::CloudFormation::CustomResource` resource
	ServiceToken *string `field:"required" json:"serviceToken" yaml:"serviceToken"`
	// Convert all property keys to pascal case.
	// Default: false.
	//
	PascalCaseProperties *bool `field:"optional" json:"pascalCaseProperties" yaml:"pascalCaseProperties"`
	// Properties to pass to the Lambda.
	//
	// Values in this `properties` dictionary can possibly overwrite other values in `CustomResourceProps`
	// E.g. `ServiceToken` and `ServiceTimeout`
	// It is recommended to avoid using same keys that exist in `CustomResourceProps`.
	// Default: - No properties.
	//
	Properties *map[string]interface{} `field:"optional" json:"properties" yaml:"properties"`
	// The policy to apply when this resource is removed from the application.
	// Default: cdk.RemovalPolicy.Destroy
	//
	RemovalPolicy RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// For custom resources, you can specify AWS::CloudFormation::CustomResource (the default) as the resource type, or you can specify your own resource type name.
	//
	// For example, you can use "Custom::MyCustomResourceTypeName".
	//
	// Custom resource type names must begin with "Custom::" and can include
	// alphanumeric characters and the following characters: _@-. You can specify
	// a custom resource type name up to a maximum length of 60 characters. You
	// cannot change the type during an update.
	//
	// Using your own resource type names helps you quickly differentiate the
	// types of custom resources in your stack. For example, if you had two custom
	// resources that conduct two different ping tests, you could name their type
	// as Custom::PingTester to make them easily identifiable as ping testers
	// (instead of using AWS::CloudFormation::CustomResource).
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cfn-customresource.html#aws-cfn-resource-type-name
	//
	// Default: - AWS::CloudFormation::CustomResource.
	//
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// The maximum time that can elapse before a custom resource operation times out.
	//
	// The value must be between 1 second and 3600 seconds.
	//
	// Maps to [ServiceTimeout](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-customresource.html#cfn-cloudformation-customresource-servicetimeout) property for the `AWS::CloudFormation::CustomResource` resource
	//
	// A token can be specified for this property, but it must be specified with `Duration.seconds()`.
	//
	// Example:
	//   stack := awscdk.Newstack()
	//   durToken := awscdk.NewCfnParameter(stack, jsii.String("MyParameter"), &CfnParameterProps{
	//   	Type: jsii.String("Number"),
	//   	Default: jsii.Number(60),
	//   })
	//   awscdk.NewCustomResource(stack, jsii.String("MyCustomResource"), &CustomResourceProps{
	//   	ServiceToken: jsii.String("MyServiceToken"),
	//   	ServiceTimeout: awscdk.Duration_Seconds(durToken.valueAsNumber),
	//   })
	//
	// Default: Duration.seconds(3600)
	//
	ServiceTimeout Duration `field:"optional" json:"serviceTimeout" yaml:"serviceTimeout"`
}

