package awscleanrooms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCollaboration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCollaborationProps := &CfnCollaborationProps{
//   	CreatorDisplayName: jsii.String("creatorDisplayName"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	QueryLogStatus: jsii.String("queryLogStatus"),
//
//   	// the properties below are optional
//   	AnalyticsEngine: jsii.String("analyticsEngine"),
//   	CreatorMemberAbilities: []*string{
//   		jsii.String("creatorMemberAbilities"),
//   	},
//   	CreatorMlMemberAbilities: &MLMemberAbilitiesProperty{
//   		CustomMlMemberAbilities: []*string{
//   			jsii.String("customMlMemberAbilities"),
//   		},
//   	},
//   	CreatorPaymentConfiguration: &PaymentConfigurationProperty{
//   		QueryCompute: &QueryComputePaymentConfigProperty{
//   			IsResponsible: jsii.Boolean(false),
//   		},
//
//   		// the properties below are optional
//   		JobCompute: &JobComputePaymentConfigProperty{
//   			IsResponsible: jsii.Boolean(false),
//   		},
//   		MachineLearning: &MLPaymentConfigProperty{
//   			ModelInference: &ModelInferencePaymentConfigProperty{
//   				IsResponsible: jsii.Boolean(false),
//   			},
//   			ModelTraining: &ModelTrainingPaymentConfigProperty{
//   				IsResponsible: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	DataEncryptionMetadata: &DataEncryptionMetadataProperty{
//   		AllowCleartext: jsii.Boolean(false),
//   		AllowDuplicates: jsii.Boolean(false),
//   		AllowJoinsOnColumnsWithDifferentNames: jsii.Boolean(false),
//   		PreserveNulls: jsii.Boolean(false),
//   	},
//   	JobLogStatus: jsii.String("jobLogStatus"),
//   	Members: []interface{}{
//   		&MemberSpecificationProperty{
//   			AccountId: jsii.String("accountId"),
//   			DisplayName: jsii.String("displayName"),
//
//   			// the properties below are optional
//   			MemberAbilities: []*string{
//   				jsii.String("memberAbilities"),
//   			},
//   			MlMemberAbilities: &MLMemberAbilitiesProperty{
//   				CustomMlMemberAbilities: []*string{
//   					jsii.String("customMlMemberAbilities"),
//   				},
//   			},
//   			PaymentConfiguration: &PaymentConfigurationProperty{
//   				QueryCompute: &QueryComputePaymentConfigProperty{
//   					IsResponsible: jsii.Boolean(false),
//   				},
//
//   				// the properties below are optional
//   				JobCompute: &JobComputePaymentConfigProperty{
//   					IsResponsible: jsii.Boolean(false),
//   				},
//   				MachineLearning: &MLPaymentConfigProperty{
//   					ModelInference: &ModelInferencePaymentConfigProperty{
//   						IsResponsible: jsii.Boolean(false),
//   					},
//   					ModelTraining: &ModelTrainingPaymentConfigProperty{
//   						IsResponsible: jsii.Boolean(false),
//   					},
//   				},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-collaboration.html
//
type CfnCollaborationProps struct {
	// A display name of the collaboration creator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-collaboration.html#cfn-cleanrooms-collaboration-creatordisplayname
	//
	CreatorDisplayName *string `field:"required" json:"creatorDisplayName" yaml:"creatorDisplayName"`
	// A description of the collaboration provided by the collaboration owner.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-collaboration.html#cfn-cleanrooms-collaboration-description
	//
	Description *string `field:"required" json:"description" yaml:"description"`
	// A human-readable identifier provided by the collaboration owner.
	//
	// Display names are not unique.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-collaboration.html#cfn-cleanrooms-collaboration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// An indicator as to whether query logging has been enabled or disabled for the collaboration.
	//
	// When `ENABLED` , AWS Clean Rooms logs details about queries run within this collaboration and those logs can be viewed in Amazon CloudWatch Logs. The default value is `DISABLED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-collaboration.html#cfn-cleanrooms-collaboration-querylogstatus
	//
	QueryLogStatus *string `field:"required" json:"queryLogStatus" yaml:"queryLogStatus"`
	// The analytics engine for the collaboration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-collaboration.html#cfn-cleanrooms-collaboration-analyticsengine
	//
	AnalyticsEngine *string `field:"optional" json:"analyticsEngine" yaml:"analyticsEngine"`
	// The abilities granted to the collaboration creator.
	//
	// *Allowed values* `CAN_QUERY` | `CAN_RECEIVE_RESULTS` | `CAN_RUN_JOB`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-collaboration.html#cfn-cleanrooms-collaboration-creatormemberabilities
	//
	CreatorMemberAbilities *[]*string `field:"optional" json:"creatorMemberAbilities" yaml:"creatorMemberAbilities"`
	// The ML member abilities for a collaboration member.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-collaboration.html#cfn-cleanrooms-collaboration-creatormlmemberabilities
	//
	CreatorMlMemberAbilities interface{} `field:"optional" json:"creatorMlMemberAbilities" yaml:"creatorMlMemberAbilities"`
	// An object representing the collaboration member's payment responsibilities set by the collaboration creator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-collaboration.html#cfn-cleanrooms-collaboration-creatorpaymentconfiguration
	//
	CreatorPaymentConfiguration interface{} `field:"optional" json:"creatorPaymentConfiguration" yaml:"creatorPaymentConfiguration"`
	// The settings for client-side encryption for cryptographic computing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-collaboration.html#cfn-cleanrooms-collaboration-dataencryptionmetadata
	//
	DataEncryptionMetadata interface{} `field:"optional" json:"dataEncryptionMetadata" yaml:"dataEncryptionMetadata"`
	// An indicator as to whether job logging has been enabled or disabled for the collaboration.
	//
	// When `ENABLED` , AWS Clean Rooms logs details about jobs run within this collaboration and those logs can be viewed in Amazon CloudWatch Logs. The default value is `DISABLED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-collaboration.html#cfn-cleanrooms-collaboration-joblogstatus
	//
	JobLogStatus *string `field:"optional" json:"jobLogStatus" yaml:"jobLogStatus"`
	// A list of initial members, not including the creator.
	//
	// This list is immutable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-collaboration.html#cfn-cleanrooms-collaboration-members
	//
	Members interface{} `field:"optional" json:"members" yaml:"members"`
	// An optional label that you can assign to a resource when you create it.
	//
	// Each tag consists of a key and an optional value, both of which you define. When you use tagging, you can also use tag-based access control in IAM policies to control access to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-collaboration.html#cfn-cleanrooms-collaboration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

