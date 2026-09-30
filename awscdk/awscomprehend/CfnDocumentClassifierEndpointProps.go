package awscomprehend

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDocumentClassifierEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDocumentClassifierEndpointProps := &CfnDocumentClassifierEndpointProps{
//   	DesiredInferenceUnits: jsii.Number(123),
//   	EndpointName: jsii.String("endpointName"),
//   	ModelArn: jsii.String("modelArn"),
//
//   	// the properties below are optional
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-documentclassifierendpoint.html
//
type CfnDocumentClassifierEndpointProps struct {
	// The desired number of inference units to be used by the model.
	//
	// Each inference unit represents throughput of 100 characters per second.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-documentclassifierendpoint.html#cfn-comprehend-documentclassifierendpoint-desiredinferenceunits
	//
	DesiredInferenceUnits *float64 `field:"required" json:"desiredInferenceUnits" yaml:"desiredInferenceUnits"`
	// The name of the endpoint.
	//
	// The name must be unique within the AWS Region and account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-documentclassifierendpoint.html#cfn-comprehend-documentclassifierendpoint-endpointname
	//
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// The Amazon Resource Name (ARN) of the document classifier model to which the endpoint is attached.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-documentclassifierendpoint.html#cfn-comprehend-documentclassifierendpoint-modelarn
	//
	ModelArn *string `field:"required" json:"modelArn" yaml:"modelArn"`
	// Tags associated with the endpoint being created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-documentclassifierendpoint.html#cfn-comprehend-documentclassifierendpoint-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

