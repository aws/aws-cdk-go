package awscomprehend

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnEntityRecognizerEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEntityRecognizerEndpointProps := &CfnEntityRecognizerEndpointProps{
//   	DesiredInferenceUnits: jsii.Number(123),
//   	EndpointName: jsii.String("endpointName"),
//
//   	// the properties below are optional
//   	DataAccessRoleArn: jsii.String("dataAccessRoleArn"),
//   	FlywheelArn: jsii.String("flywheelArn"),
//   	ModelArn: jsii.String("modelArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-entityrecognizerendpoint.html
//
type CfnEntityRecognizerEndpointProps struct {
	// The desired number of inference units to be used by the model.
	//
	// Each inference unit represents throughput of 100 characters per second.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-entityrecognizerendpoint.html#cfn-comprehend-entityrecognizerendpoint-desiredinferenceunits
	//
	DesiredInferenceUnits *float64 `field:"required" json:"desiredInferenceUnits" yaml:"desiredInferenceUnits"`
	// The name of the endpoint.
	//
	// The name must be unique within the AWS Region and account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-entityrecognizerendpoint.html#cfn-comprehend-entityrecognizerendpoint-endpointname
	//
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend read access to trained custom models encrypted with a customer managed key (ModelKmsKeyId).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-entityrecognizerendpoint.html#cfn-comprehend-entityrecognizerendpoint-dataaccessrolearn
	//
	DataAccessRoleArn *string `field:"optional" json:"dataAccessRoleArn" yaml:"dataAccessRoleArn"`
	// The Amazon Resource Name (ARN) of the flywheel to which the endpoint is attached.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-entityrecognizerendpoint.html#cfn-comprehend-entityrecognizerendpoint-flywheelarn
	//
	FlywheelArn *string `field:"optional" json:"flywheelArn" yaml:"flywheelArn"`
	// The Amazon Resource Name (ARN) of the entity recognizer model to which the endpoint is attached.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-entityrecognizerendpoint.html#cfn-comprehend-entityrecognizerendpoint-modelarn
	//
	ModelArn *string `field:"optional" json:"modelArn" yaml:"modelArn"`
	// Tags associated with the endpoint being created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-entityrecognizerendpoint.html#cfn-comprehend-entityrecognizerendpoint-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

