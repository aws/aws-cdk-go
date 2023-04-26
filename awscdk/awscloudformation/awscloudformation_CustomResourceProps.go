package awscloudformation

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties to provide a Lambda-backed custom resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var customResourceProvider customResourceProvider
//   var properties interface{}
//
//   customResourceProps := &CustomResourceProps{
//   	Provider: customResourceProvider,
//
//   	// the properties below are optional
//   	Properties: map[string]interface{}{
//   		"propertiesKey": properties,
//   	},
//   	RemovalPolicy: monocdk.RemovalPolicy_DESTROY,
//   	ResourceType: jsii.String("resourceType"),
//   }
//
// Deprecated: use `core.CustomResourceProps`
type CustomResourceProps struct {
	// The provider which implements the custom resource.
	//
	// You can implement a provider by listening to raw AWS CloudFormation events
	// through an SNS topic or an AWS Lambda function or use the CDK's custom
	// [resource provider framework] which makes it easier to implement robust
	// providers.
	//
	// [resource provider framework]: https://docs.aws.amazon.com/cdk/api/latest/docs/custom-resources-readme.html
	//
	// ```ts
	// import * as custom_resources from '@aws-cdk/custom-resources';
	// import * as lambda from '@aws-cdk/aws-lambda';
	// import { Stack } from '@aws-cdk/core';
	// declare const myOnEventLambda: lambda.Function;
	// declare const myIsCompleteLambda: lambda.Function;
	// const stack = new Stack();
	//
	// const provider = new custom_resources.Provider(stack, 'myProvider', {
	//    onEventHandler: myOnEventLambda,
	//    isCompleteHandler: myIsCompleteLambda, // optional
	// });
	// ```
	//
	// ```ts
	// import * as cloudformation from '@aws-cdk/aws-cloudformation';
	// import * as lambda from '@aws-cdk/aws-lambda';
	// declare const myFunction: lambda.Function;
	//
	// // invoke an AWS Lambda function when a lifecycle event occurs:
	// const provider = cloudformation.CustomResourceProvider.fromLambda(myFunction);
	// ```
	//
	// ```ts
	// import * as cloudformation from '@aws-cdk/aws-cloudformation';
	// import * as sns from '@aws-cdk/aws-sns';
	// declare const myTopic: sns.Topic;
	//
	// // publish lifecycle events to an SNS topic:
	// const provider = cloudformation.CustomResourceProvider.fromTopic(myTopic);
	// ```.
	// Deprecated: use `core.CustomResourceProps`
	Provider ICustomResourceProvider `field:"required" json:"provider" yaml:"provider"`
	// Properties to pass to the Lambda.
	// Deprecated: use `core.CustomResourceProps`
	Properties *map[string]interface{} `field:"optional" json:"properties" yaml:"properties"`
	// The policy to apply when this resource is removed from the application.
	// Deprecated: use `core.CustomResourceProps`
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
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
	// Deprecated: use `core.CustomResourceProps`
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
}

