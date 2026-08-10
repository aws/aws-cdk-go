package awsauditmanager


// A control set entity that represents a collection of controls in Audit Manager.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   controlSetProperty := &ControlSetProperty{
//   	Controls: []interface{}{
//   		&ControlSetControlProperty{
//   			Id: jsii.String("id"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessmentframework-controlset.html
//
type CfnAssessmentFramework_ControlSetProperty struct {
	// The list of controls within the control set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessmentframework-controlset.html#cfn-auditmanager-assessmentframework-controlset-controls
	//
	Controls interface{} `field:"required" json:"controls" yaml:"controls"`
	// The name of the control set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessmentframework-controlset.html#cfn-auditmanager-assessmentframework-controlset-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
}

