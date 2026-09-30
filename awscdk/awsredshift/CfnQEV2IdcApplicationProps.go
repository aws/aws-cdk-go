package awsredshift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnQEV2IdcApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnQEV2IdcApplicationProps := &CfnQEV2IdcApplicationProps{
//   	IdcDisplayName: jsii.String("idcDisplayName"),
//   	IdcInstanceArn: jsii.String("idcInstanceArn"),
//   	Qev2IdcApplicationName: jsii.String("qev2IdcApplicationName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-qev2idcapplication.html
//
type CfnQEV2IdcApplicationProps struct {
	// The display name for the Amazon Redshift Query Editor (QEV2) IAM Identity Center application.
	//
	// It appears in the console.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-qev2idcapplication.html#cfn-redshift-qev2idcapplication-idcdisplayname
	//
	IdcDisplayName *string `field:"required" json:"idcDisplayName" yaml:"idcDisplayName"`
	// The Amazon Resource Name (ARN) of the IAM Identity Center instance used to create the Amazon Redshift Query Editor (QEV2) managed application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-qev2idcapplication.html#cfn-redshift-qev2idcapplication-idcinstancearn
	//
	IdcInstanceArn *string `field:"required" json:"idcInstanceArn" yaml:"idcInstanceArn"`
	// The name of the Amazon Redshift Query Editor (QEV2) application in IAM Identity Center.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-qev2idcapplication.html#cfn-redshift-qev2idcapplication-qev2idcapplicationname
	//
	Qev2IdcApplicationName *string `field:"required" json:"qev2IdcApplicationName" yaml:"qev2IdcApplicationName"`
	// A list of tags associated with the application.
	//
	// Tags are key-value pairs that you can use to organize and identify your resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-qev2idcapplication.html#cfn-redshift-qev2idcapplication-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

