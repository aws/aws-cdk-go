package awsdynamodb


// Reference to ContributorInsightsSpecification.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   app := cdk.NewApp()
//   stack := cdk.NewStack(app, jsii.String("Stack"), &StackProps{
//   	Env: &Environment{
//   		Region: jsii.String("us-west-2"),
//   	},
//   })
//
//   globalTable := dynamodb.NewTableV2(stack, jsii.String("GlobalTable"), &TablePropsV2{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("pk"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	ContributorInsightsSpecification: &ContributorInsightsSpecification{
//   		Enabled: jsii.Boolean(true),
//   	},
//   	PointInTimeRecoverySpecification: &PointInTimeRecoverySpecification{
//   		PointInTimeRecoveryEnabled: jsii.Boolean(true),
//   	},
//   	Replicas: []replicaTableProps{
//   		&replicaTableProps{
//   			Region: jsii.String("us-east-1"),
//   			TableClass: dynamodb.TableClass_STANDARD_INFREQUENT_ACCESS,
//   			PointInTimeRecoverySpecification: &PointInTimeRecoverySpecification{
//   				PointInTimeRecoveryEnabled: jsii.Boolean(false),
//   			},
//   		},
//   		&replicaTableProps{
//   			Region: jsii.String("us-east-2"),
//   			ContributorInsightsSpecification: &ContributorInsightsSpecification{
//   				Enabled: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   })
//
type ContributorInsightsSpecification struct {
	// Indicates whether contributor insights is enabled.
	// Default: false.
	//
	Enabled *bool `field:"required" json:"enabled" yaml:"enabled"`
	// Indicates the type of metrics captured by contributor insights.
	// Default: ACCESSED_AND_THROTTLED_KEYS.
	//
	Mode ContributorInsightsMode `field:"optional" json:"mode" yaml:"mode"`
}

