package awsdeadline


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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-monitor.html
//
type CfnMonitorProps struct {
	// The name of the monitor that displays on the Deadline Cloud console.
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
}

