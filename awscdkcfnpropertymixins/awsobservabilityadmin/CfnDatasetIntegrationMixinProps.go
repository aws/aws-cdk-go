package awsobservabilityadmin

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDatasetIntegrationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnDatasetIntegrationMixinProps := &CfnDatasetIntegrationMixinProps{
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-datasetintegration.html
//
type CfnDatasetIntegrationMixinProps struct {
	// The ARN of the IAM role that the service assumes to forward telemetry to the customer's CloudWatch dataset.
	//
	// The role must live in the caller's account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-datasetintegration.html#cfn-observabilityadmin-datasetintegration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-datasetintegration.html#cfn-observabilityadmin-datasetintegration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

