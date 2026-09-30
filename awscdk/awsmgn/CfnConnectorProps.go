package awsmgn

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnConnector`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectorProps := &CfnConnectorProps{
//   	Name: jsii.String("name"),
//   	SsmInstanceId: jsii.String("ssmInstanceId"),
//
//   	// the properties below are optional
//   	SsmCommandConfig: &ConnectorSsmCommandConfigProperty{
//   		CloudWatchOutputEnabled: jsii.Boolean(false),
//   		S3OutputEnabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		CloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   		OutputS3BucketName: jsii.String("outputS3BucketName"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-connector.html
//
type CfnConnectorProps struct {
	// The name of the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-connector.html#cfn-mgn-connector-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The SSM instance ID associated with this connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-connector.html#cfn-mgn-connector-ssminstanceid
	//
	SsmInstanceId *string `field:"required" json:"ssmInstanceId" yaml:"ssmInstanceId"`
	// SSM command configuration for the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-connector.html#cfn-mgn-connector-ssmcommandconfig
	//
	SsmCommandConfig interface{} `field:"optional" json:"ssmCommandConfig" yaml:"ssmCommandConfig"`
	// Tags to assign to the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-connector.html#cfn-mgn-connector-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

