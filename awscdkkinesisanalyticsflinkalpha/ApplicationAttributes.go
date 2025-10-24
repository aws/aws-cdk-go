package awscdkkinesisanalyticsflinkalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Attributes used for importing an Application with Application.fromApplicationAttributes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import kinesisanalytics_flink_alpha "github.com/aws/aws-cdk-go/awscdkkinesisanalyticsflinkalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup SecurityGroup
//
//   applicationAttributes := &ApplicationAttributes{
//   	ApplicationArn: jsii.String("applicationArn"),
//
//   	// the properties below are optional
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   }
//
// Experimental.
type ApplicationAttributes struct {
	// The ARN of the Flink application.
	//
	// Format: arn:<partition>:kinesisanalytics:<region>:<account-id>:application/<application-name>.
	// Experimental.
	ApplicationArn *string `field:"required" json:"applicationArn" yaml:"applicationArn"`
	// The security groups for this Flink application if deployed in a VPC.
	// Default: - no security groups.
	//
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

