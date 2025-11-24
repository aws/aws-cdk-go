package mixinsawsauditmanager


// The `Delegation` property type specifies the assignment of a control set to a delegate for review.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   delegationProperty := &DelegationProperty{
//   	AssessmentId: jsii.String("assessmentId"),
//   	AssessmentName: jsii.String("assessmentName"),
//   	Comment: jsii.String("comment"),
//   	ControlSetId: jsii.String("controlSetId"),
//   	CreatedBy: jsii.String("createdBy"),
//   	CreationTime: jsii.Number(123),
//   	Id: jsii.String("id"),
//   	LastUpdated: jsii.Number(123),
//   	RoleArn: jsii.String("roleArn"),
//   	RoleType: jsii.String("roleType"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-delegation.html
//
type CfnAssessmentPropsMixin_DelegationProperty struct {
	// The identifier for the assessment that's associated with the delegation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-delegation.html#cfn-auditmanager-assessment-delegation-assessmentid
	//
	AssessmentId *string `field:"optional" json:"assessmentId" yaml:"assessmentId"`
	// The name of the assessment that's associated with the delegation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-delegation.html#cfn-auditmanager-assessment-delegation-assessmentname
	//
	AssessmentName *string `field:"optional" json:"assessmentName" yaml:"assessmentName"`
	// The comment that's related to the delegation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-delegation.html#cfn-auditmanager-assessment-delegation-comment
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The identifier for the control set that's associated with the delegation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-delegation.html#cfn-auditmanager-assessment-delegation-controlsetid
	//
	ControlSetId *string `field:"optional" json:"controlSetId" yaml:"controlSetId"`
	// The user or role that created the delegation.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `100`
	//
	// *Pattern* : `^[a-zA-Z0-9-_()\\[\\]\\s]+$`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-delegation.html#cfn-auditmanager-assessment-delegation-createdby
	//
	CreatedBy *string `field:"optional" json:"createdBy" yaml:"createdBy"`
	// Specifies when the delegation was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-delegation.html#cfn-auditmanager-assessment-delegation-creationtime
	//
	CreationTime *float64 `field:"optional" json:"creationTime" yaml:"creationTime"`
	// The unique identifier for the delegation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-delegation.html#cfn-auditmanager-assessment-delegation-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies when the delegation was last updated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-delegation.html#cfn-auditmanager-assessment-delegation-lastupdated
	//
	LastUpdated *float64 `field:"optional" json:"lastUpdated" yaml:"lastUpdated"`
	// The Amazon Resource Name (ARN) of the IAM role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-delegation.html#cfn-auditmanager-assessment-delegation-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The type of customer persona.
	//
	// > In `CreateAssessment` , `roleType` can only be `PROCESS_OWNER` .
	// >
	// > In `UpdateSettings` , `roleType` can only be `PROCESS_OWNER` .
	// >
	// > In `BatchCreateDelegationByAssessment` , `roleType` can only be `RESOURCE_OWNER` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-delegation.html#cfn-auditmanager-assessment-delegation-roletype
	//
	RoleType *string `field:"optional" json:"roleType" yaml:"roleType"`
	// The status of the delegation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-delegation.html#cfn-auditmanager-assessment-delegation-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

