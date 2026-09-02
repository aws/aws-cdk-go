package awsbedrockagentcore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnHarnessEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnHarnessEndpointProps := &CfnHarnessEndpointProps{
//   	EndpointName: jsii.String("endpointName"),
//   	HarnessId: jsii.String("harnessId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetVersion: jsii.String("targetVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-harnessendpoint.html
//
type CfnHarnessEndpointProps struct {
	// The name of the endpoint.
	//
	// Must start with a letter and contain only alphanumeric characters and underscores.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-harnessendpoint.html#cfn-bedrockagentcore-harnessendpoint-endpointname
	//
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// The ID of the harness that the endpoint belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-harnessendpoint.html#cfn-bedrockagentcore-harnessendpoint-harnessid
	//
	HarnessId *string `field:"required" json:"harnessId" yaml:"harnessId"`
	// The description of the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-harnessendpoint.html#cfn-bedrockagentcore-harnessendpoint-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags to apply to the harness endpoint resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-harnessendpoint.html#cfn-bedrockagentcore-harnessendpoint-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The harness version that the endpoint points to and serves invocations from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-harnessendpoint.html#cfn-bedrockagentcore-harnessendpoint-targetversion
	//
	TargetVersion *string `field:"optional" json:"targetVersion" yaml:"targetVersion"`
}

