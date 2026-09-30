package awsmgn

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnConnectorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnConnectorMixinProps := &CfnConnectorMixinProps{
//   	Name: jsii.String("name"),
//   	SsmCommandConfig: &ConnectorSsmCommandConfigProperty{
//   		CloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   		CloudWatchOutputEnabled: jsii.Boolean(false),
//   		OutputS3BucketName: jsii.String("outputS3BucketName"),
//   		S3OutputEnabled: jsii.Boolean(false),
//   	},
//   	SsmInstanceId: jsii.String("ssmInstanceId"),
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
type CfnConnectorMixinProps struct {
	// The name of the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-connector.html#cfn-mgn-connector-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// SSM command configuration for the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-connector.html#cfn-mgn-connector-ssmcommandconfig
	//
	SsmCommandConfig interface{} `field:"optional" json:"ssmCommandConfig" yaml:"ssmCommandConfig"`
	// The SSM instance ID associated with this connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-connector.html#cfn-mgn-connector-ssminstanceid
	//
	SsmInstanceId *string `field:"optional" json:"ssmInstanceId" yaml:"ssmInstanceId"`
	// Tags to assign to the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-connector.html#cfn-mgn-connector-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

