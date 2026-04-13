package awssecurityagent

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnApplicationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnApplicationMixinProps := &CfnApplicationMixinProps{
//   	DefaultKmsKeyId: jsii.String("defaultKmsKeyId"),
//   	IdCConfiguration: &IdCConfigurationProperty{
//   		IdCApplicationArn: jsii.String("idCApplicationArn"),
//   		IdCInstanceArn: jsii.String("idCInstanceArn"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-application.html
//
type CfnApplicationMixinProps struct {
	// Identifier of a KMS key.
	//
	// Can be a key ID, key ARN, alias name, or alias ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-application.html#cfn-securityagent-application-defaultkmskeyid
	//
	DefaultKmsKeyId *string `field:"optional" json:"defaultKmsKeyId" yaml:"defaultKmsKeyId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-application.html#cfn-securityagent-application-idcconfiguration
	//
	IdCConfiguration interface{} `field:"optional" json:"idCConfiguration" yaml:"idCConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-application.html#cfn-securityagent-application-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Tags for the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-application.html#cfn-securityagent-application-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

