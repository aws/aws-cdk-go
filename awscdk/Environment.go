package awscdk


// The deployment environment for a stack.
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
//   	Replicas: []replicaTableProps{
//   		&replicaTableProps{
//   			Region: jsii.String("us-east-1"),
//   		},
//   	},
//   })
//
//   globalTable.AddReplica(&replicaTableProps{
//   	Region: jsii.String("us-east-2"),
//   	DeletionProtection: jsii.Boolean(true),
//   })
//
type Environment struct {
	// The AWS account ID for this environment.
	//
	// This can be either a concrete value such as `585191031104` or `Aws.ACCOUNT_ID` which
	// indicates that account ID will only be determined during deployment (it
	// will resolve to the CloudFormation intrinsic `{"Ref":"AWS::AccountId"}`).
	// Note that certain features, such as cross-stack references and
	// environmental context providers require concrete region information and
	// will cause this stack to emit synthesis errors.
	// Default: Aws.ACCOUNT_ID which means that the stack will be account-agnostic.
	//
	Account *string `field:"optional" json:"account" yaml:"account"`
	// The AWS region for this environment.
	//
	// This can be either a concrete value such as `eu-west-2` or `Aws.REGION`
	// which indicates that account ID will only be determined during deployment
	// (it will resolve to the CloudFormation intrinsic `{"Ref":"AWS::Region"}`).
	// Note that certain features, such as cross-stack references and
	// environmental context providers require concrete region information and
	// will cause this stack to emit synthesis errors.
	// Default: Aws.REGION which means that the stack will be region-agnostic.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}

