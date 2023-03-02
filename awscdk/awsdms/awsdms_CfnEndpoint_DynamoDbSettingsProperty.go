package awsdms


// Provides information, including the Amazon Resource Name (ARN) of the IAM role used to define an Amazon DynamoDB target endpoint.
//
// This information also includes the output format of records applied to the endpoint and details of transaction and control table data information. For information about other available settings, see [Using object mapping to migrate data to DynamoDB](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.DynamoDB.html#CHAP_Target.DynamoDB.ObjectMapping) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dynamoDbSettingsProperty := &dynamoDbSettingsProperty{
//   	serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   }
//
type CfnEndpoint_DynamoDbSettingsProperty struct {
	// The Amazon Resource Name (ARN) used by the service to access the IAM role.
	//
	// The role must allow the `iam:PassRole` action.
	ServiceAccessRoleArn *string `field:"optional" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
}

