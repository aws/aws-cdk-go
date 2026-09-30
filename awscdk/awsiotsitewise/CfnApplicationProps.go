package awsiotsitewise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &CfnApplicationProps{
//   	Name: jsii.String("name"),
//   	WorkspaceName: jsii.String("workspaceName"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	IdcInstanceArn: jsii.String("idcInstanceArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-application.html
//
type CfnApplicationProps struct {
	// The name of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-application.html#cfn-iotsitewise-application-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the workspace that the application belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-application.html#cfn-iotsitewise-application-workspacename
	//
	WorkspaceName *string `field:"required" json:"workspaceName" yaml:"workspaceName"`
	// A description of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-application.html#cfn-iotsitewise-application-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ARN of the IAM Identity Center instance used to create the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-application.html#cfn-iotsitewise-application-idcinstancearn
	//
	IdcInstanceArn *string `field:"optional" json:"idcInstanceArn" yaml:"idcInstanceArn"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-application.html#cfn-iotsitewise-application-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

