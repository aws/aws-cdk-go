package awsauditmanager


// A reference to an existing control by ID.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   controlSetControlProperty := &ControlSetControlProperty{
//   	Id: jsii.String("id"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessmentframework-controlsetcontrol.html
//
type CfnAssessmentFramework_ControlSetControlProperty struct {
	// The unique identifier of the control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessmentframework-controlsetcontrol.html#cfn-auditmanager-assessmentframework-controlsetcontrol-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
}

