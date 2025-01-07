package awscleanrooms


// Basic metadata used to construct a new member.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memberSpecificationProperty := &MemberSpecificationProperty{
//   	AccountId: jsii.String("accountId"),
//   	DisplayName: jsii.String("displayName"),
//   	MemberAbilities: []*string{
//   		jsii.String("memberAbilities"),
//   	},
//
//   	// the properties below are optional
//   	MlMemberAbilities: &MLMemberAbilitiesProperty{
//   		CustomMlMemberAbilities: []*string{
//   			jsii.String("customMlMemberAbilities"),
//   		},
//   	},
//   	PaymentConfiguration: &PaymentConfigurationProperty{
//   		QueryCompute: &QueryComputePaymentConfigProperty{
//   			IsResponsible: jsii.Boolean(false),
//   		},
//
//   		// the properties below are optional
//   		MachineLearning: &MLPaymentConfigProperty{
//   			ModelInference: &ModelInferencePaymentConfigProperty{
//   				IsResponsible: jsii.Boolean(false),
//   			},
//   			ModelTraining: &ModelTrainingPaymentConfigProperty{
//   				IsResponsible: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-memberspecification.html
//
type CfnCollaboration_MemberSpecificationProperty struct {
	// The identifier used to reference members of the collaboration.
	//
	// Currently only supports AWS account ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-memberspecification.html#cfn-cleanrooms-collaboration-memberspecification-accountid
	//
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The member's display name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-memberspecification.html#cfn-cleanrooms-collaboration-memberspecification-displayname
	//
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The abilities granted to the collaboration member.
	//
	// *Allowed Values* : `CAN_QUERY` | `CAN_RECEIVE_RESULTS`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-memberspecification.html#cfn-cleanrooms-collaboration-memberspecification-memberabilities
	//
	MemberAbilities *[]*string `field:"required" json:"memberAbilities" yaml:"memberAbilities"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-memberspecification.html#cfn-cleanrooms-collaboration-memberspecification-mlmemberabilities
	//
	MlMemberAbilities interface{} `field:"optional" json:"mlMemberAbilities" yaml:"mlMemberAbilities"`
	// The collaboration member's payment responsibilities set by the collaboration creator.
	//
	// If the collaboration creator hasn't speciﬁed anyone as the member paying for query compute costs, then the member who can query is the default payer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-memberspecification.html#cfn-cleanrooms-collaboration-memberspecification-paymentconfiguration
	//
	PaymentConfiguration interface{} `field:"optional" json:"paymentConfiguration" yaml:"paymentConfiguration"`
}

