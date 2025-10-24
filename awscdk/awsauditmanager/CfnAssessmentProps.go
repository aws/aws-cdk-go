package awsauditmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAssessment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAssessmentProps := &CfnAssessmentProps{
//   	AssessmentReportsDestination: &AssessmentReportsDestinationProperty{
//   		Destination: jsii.String("destination"),
//   		DestinationType: jsii.String("destinationType"),
//   	},
//   	AwsAccount: &AWSAccountProperty{
//   		EmailAddress: jsii.String("emailAddress"),
//   		Id: jsii.String("id"),
//   		Name: jsii.String("name"),
//   	},
//   	Delegations: []interface{}{
//   		&DelegationProperty{
//   			AssessmentId: jsii.String("assessmentId"),
//   			AssessmentName: jsii.String("assessmentName"),
//   			Comment: jsii.String("comment"),
//   			ControlSetId: jsii.String("controlSetId"),
//   			CreatedBy: jsii.String("createdBy"),
//   			CreationTime: jsii.Number(123),
//   			Id: jsii.String("id"),
//   			LastUpdated: jsii.Number(123),
//   			RoleArn: jsii.String("roleArn"),
//   			RoleType: jsii.String("roleType"),
//   			Status: jsii.String("status"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	FrameworkId: jsii.String("frameworkId"),
//   	Name: jsii.String("name"),
//   	Roles: []interface{}{
//   		&RoleProperty{
//   			RoleArn: jsii.String("roleArn"),
//   			RoleType: jsii.String("roleType"),
//   		},
//   	},
//   	Scope: &ScopeProperty{
//   		AwsAccounts: []interface{}{
//   			&AWSAccountProperty{
//   				EmailAddress: jsii.String("emailAddress"),
//   				Id: jsii.String("id"),
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		AwsServices: []interface{}{
//   			&AWSServiceProperty{
//   				ServiceName: jsii.String("serviceName"),
//   			},
//   		},
//   	},
//   	Status: jsii.String("status"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html
//
type CfnAssessmentProps struct {
	// The destination that evidence reports are stored in for the assessment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-assessmentreportsdestination
	//
	AssessmentReportsDestination interface{} `field:"optional" json:"assessmentReportsDestination" yaml:"assessmentReportsDestination"`
	// The AWS account that's associated with the assessment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-awsaccount
	//
	AwsAccount interface{} `field:"optional" json:"awsAccount" yaml:"awsAccount"`
	// The delegations that are associated with the assessment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-delegations
	//
	Delegations interface{} `field:"optional" json:"delegations" yaml:"delegations"`
	// The description of the assessment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The unique identifier for the framework.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-frameworkid
	//
	FrameworkId *string `field:"optional" json:"frameworkId" yaml:"frameworkId"`
	// The name of the assessment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The roles that are associated with the assessment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-roles
	//
	Roles interface{} `field:"optional" json:"roles" yaml:"roles"`
	// The wrapper of AWS accounts and services that are in scope for the assessment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-scope
	//
	Scope interface{} `field:"optional" json:"scope" yaml:"scope"`
	// The overall status of the assessment.
	//
	// When you create a new assessment, the initial `Status` value is always `ACTIVE` . When you create an assessment, even if you specify the value as `INACTIVE` , the value overrides to `ACTIVE` .
	//
	// After you create an assessment, you can change the value of the `Status` property at any time. For example, when you want to stop collecting evidence for your assessment, you can change the assessment status to `INACTIVE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The tags that are associated with the assessment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

