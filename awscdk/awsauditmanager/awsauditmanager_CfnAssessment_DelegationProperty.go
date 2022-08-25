package awsauditmanager


// The `Delegation` property type specifies the assignment of a control set to a delegate for review.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   delegationProperty := &delegationProperty{
//   	assessmentId: jsii.String("assessmentId"),
//   	assessmentName: jsii.String("assessmentName"),
//   	comment: jsii.String("comment"),
//   	controlSetId: jsii.String("controlSetId"),
//   	createdBy: jsii.String("createdBy"),
//   	creationTime: jsii.Number(123),
//   	id: jsii.String("id"),
//   	lastUpdated: jsii.Number(123),
//   	roleArn: jsii.String("roleArn"),
//   	roleType: jsii.String("roleType"),
//   	status: jsii.String("status"),
//   }
//
type CfnAssessment_DelegationProperty struct {
	// The identifier for the assessment that's associated with the delegation.
	AssessmentId *string `field:"optional" json:"assessmentId" yaml:"assessmentId"`
	// The name of the assessment that's associated with the delegation.
	AssessmentName *string `field:"optional" json:"assessmentName" yaml:"assessmentName"`
	// The comment that's related to the delegation.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The identifier for the control set that's associated with the delegation.
	ControlSetId *string `field:"optional" json:"controlSetId" yaml:"controlSetId"`
	// The IAM user or role that created the delegation.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `100`
	//
	// *Pattern* : `^[a-zA-Z0-9-_()\\[\\]\\s]+$`.
	CreatedBy *string `field:"optional" json:"createdBy" yaml:"createdBy"`
	// Specifies when the delegation was created.
	CreationTime *float64 `field:"optional" json:"creationTime" yaml:"creationTime"`
	// The unique identifier for the delegation.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies when the delegation was last updated.
	LastUpdated *float64 `field:"optional" json:"lastUpdated" yaml:"lastUpdated"`
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
	// The status of the delegation.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

