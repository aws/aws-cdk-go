package awsdeadline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMonitor`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMonitorProps := &CfnMonitorProps{
//   	DisplayName: jsii.String("displayName"),
//   	IdentityCenterInstanceArn: jsii.String("identityCenterInstanceArn"),
//   	RoleArn: jsii.String("roleArn"),
//   	Subdomain: jsii.String("subdomain"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-monitor.html
//
type CfnMonitorProps struct {
	// The name of the monitor that displays on the Deadline Cloud console.
	//
	// > This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-monitor.html#cfn-deadline-monitor-displayname
	//
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The Amazon Resource Name (ARN) of the IAM Identity Center instance responsible for authenticating monitor users.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-monitor.html#cfn-deadline-monitor-identitycenterinstancearn
	//
	IdentityCenterInstanceArn *string `field:"required" json:"identityCenterInstanceArn" yaml:"identityCenterInstanceArn"`
	// The Amazon Resource Name (ARN) of the IAM role for the monitor.
	//
	// Users of the monitor use this role to access Deadline Cloud resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-monitor.html#cfn-deadline-monitor-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The subdomain used for the monitor URL.
	//
	// The full URL of the monitor is subdomain.Region.deadlinecloud.amazonaws.com.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-monitor.html#cfn-deadline-monitor-subdomain
	//
	Subdomain *string `field:"required" json:"subdomain" yaml:"subdomain"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-monitor.html#cfn-deadline-monitor-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

