package awscleanrooms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMembership`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMembershipProps := &CfnMembershipProps{
//   	CollaborationIdentifier: jsii.String("collaborationIdentifier"),
//   	QueryLogStatus: jsii.String("queryLogStatus"),
//
//   	// the properties below are optional
//   	DefaultResultConfiguration: &MembershipProtectedQueryResultConfigurationProperty{
//   		OutputConfiguration: &MembershipProtectedQueryOutputConfigurationProperty{
//   			S3: &ProtectedQueryS3OutputConfigurationProperty{
//   				Bucket: jsii.String("bucket"),
//   				ResultFormat: jsii.String("resultFormat"),
//
//   				// the properties below are optional
//   				KeyPrefix: jsii.String("keyPrefix"),
//   				SingleFileOutput: jsii.Boolean(false),
//   			},
//   		},
//
//   		// the properties below are optional
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	PaymentConfiguration: &MembershipPaymentConfigurationProperty{
//   		QueryCompute: &MembershipQueryComputePaymentConfigProperty{
//   			IsResponsible: jsii.Boolean(false),
//   		},
//
//   		// the properties below are optional
//   		MachineLearning: &MembershipMLPaymentConfigProperty{
//   			ModelInference: &MembershipModelInferencePaymentConfigProperty{
//   				IsResponsible: jsii.Boolean(false),
//   			},
//   			ModelTraining: &MembershipModelTrainingPaymentConfigProperty{
//   				IsResponsible: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-membership.html
//
type CfnMembershipProps struct {
	// The unique ID for the associated collaboration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-membership.html#cfn-cleanrooms-membership-collaborationidentifier
	//
	CollaborationIdentifier *string `field:"required" json:"collaborationIdentifier" yaml:"collaborationIdentifier"`
	// An indicator as to whether query logging has been enabled or disabled for the membership.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-membership.html#cfn-cleanrooms-membership-querylogstatus
	//
	QueryLogStatus *string `field:"required" json:"queryLogStatus" yaml:"queryLogStatus"`
	// The default protected query result configuration as specified by the member who can receive results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-membership.html#cfn-cleanrooms-membership-defaultresultconfiguration
	//
	DefaultResultConfiguration interface{} `field:"optional" json:"defaultResultConfiguration" yaml:"defaultResultConfiguration"`
	// The payment responsibilities accepted by the collaboration member.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-membership.html#cfn-cleanrooms-membership-paymentconfiguration
	//
	PaymentConfiguration interface{} `field:"optional" json:"paymentConfiguration" yaml:"paymentConfiguration"`
	// An optional label that you can assign to a resource when you create it.
	//
	// Each tag consists of a key and an optional value, both of which you define. When you use tagging, you can also use tag-based access control in IAM policies to control access to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-membership.html#cfn-cleanrooms-membership-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

