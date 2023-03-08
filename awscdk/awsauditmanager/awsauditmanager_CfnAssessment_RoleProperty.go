package awsauditmanager


// The `Role` property type specifies the wrapper that contains AWS Audit Manager role information, such as the role type and IAM Amazon Resource Name (ARN).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   roleProperty := &roleProperty{
//   	roleArn: jsii.String("roleArn"),
//   	roleType: jsii.String("roleType"),
//   }
//
type CfnAssessment_RoleProperty struct {
	// The Amazon Resource Name (ARN) of the IAM role.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The type of customer persona.
	//
	// > In `CreateAssessment` , `roleType` can only be `PROCESS_OWNER` .
	// >
	// > In `UpdateSettings` , `roleType` can only be `PROCESS_OWNER` .
	// >
	// > In `BatchCreateDelegationByAssessment` , `roleType` can only be `RESOURCE_OWNER` .
	RoleType *string `field:"optional" json:"roleType" yaml:"roleType"`
}

