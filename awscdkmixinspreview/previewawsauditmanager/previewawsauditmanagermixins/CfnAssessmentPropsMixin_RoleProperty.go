package previewawsauditmanagermixins


// The `Role` property type specifies the wrapper that contains AWS Audit Manager role information, such as the role type and IAM Amazon Resource Name (ARN).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   roleProperty := &RoleProperty{
//   	RoleArn: jsii.String("roleArn"),
//   	RoleType: jsii.String("roleType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-role.html
//
type CfnAssessmentPropsMixin_RoleProperty struct {
	// The Amazon Resource Name (ARN) of the IAM role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-role.html#cfn-auditmanager-assessment-role-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The type of customer persona.
	//
	// > In `CreateAssessment` , `roleType` can only be `PROCESS_OWNER` .
	// >
	// > In `UpdateSettings` , `roleType` can only be `PROCESS_OWNER` .
	// >
	// > In `BatchCreateDelegationByAssessment` , `roleType` can only be `RESOURCE_OWNER` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-role.html#cfn-auditmanager-assessment-role-roletype
	//
	RoleType *string `field:"optional" json:"roleType" yaml:"roleType"`
}

