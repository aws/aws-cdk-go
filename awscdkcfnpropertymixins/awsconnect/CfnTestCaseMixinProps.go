package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnTestCasePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnTestCaseMixinProps := &CfnTestCaseMixinProps{
//   	Content: jsii.String("content"),
//   	Description: jsii.String("description"),
//   	EntryPoint: &EntryPointProperty{
//   		ChatEntryPointParameters: &ChatEntryPointParametersProperty{
//   			FlowId: jsii.String("flowId"),
//   		},
//   		Type: jsii.String("type"),
//   		VoiceCallEntryPointParameters: &VoiceCallEntryPointParametersProperty{
//   			DestinationPhoneNumber: jsii.String("destinationPhoneNumber"),
//   			FlowId: jsii.String("flowId"),
//   			SourcePhoneNumber: jsii.String("sourcePhoneNumber"),
//   		},
//   	},
//   	InitializationData: jsii.String("initializationData"),
//   	InstanceArn: jsii.String("instanceArn"),
//   	Name: jsii.String("name"),
//   	Status: jsii.String("status"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-testcase.html
//
type CfnTestCaseMixinProps struct {
	// The content of the test case.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-testcase.html#cfn-connect-testcase-content
	//
	Content *string `field:"optional" json:"content" yaml:"content"`
	// The description of the test case.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-testcase.html#cfn-connect-testcase-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Entry Point associated with the test case.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-testcase.html#cfn-connect-testcase-entrypoint
	//
	EntryPoint interface{} `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// The initialization data of the test case.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-testcase.html#cfn-connect-testcase-initializationdata
	//
	InitializationData *string `field:"optional" json:"initializationData" yaml:"initializationData"`
	// The identifier of the Amazon Connect instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-testcase.html#cfn-connect-testcase-instancearn
	//
	InstanceArn *string `field:"optional" json:"instanceArn" yaml:"instanceArn"`
	// The name of the test case.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-testcase.html#cfn-connect-testcase-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The status of the test case.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-testcase.html#cfn-connect-testcase-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// One or more tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-testcase.html#cfn-connect-testcase-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

