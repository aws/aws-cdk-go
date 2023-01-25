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
//   cfnAssessmentProps := &cfnAssessmentProps{
//   	assessmentReportsDestination: &assessmentReportsDestinationProperty{
//   		destination: jsii.String("destination"),
//   		destinationType: jsii.String("destinationType"),
//   	},
//   	awsAccount: &aWSAccountProperty{
//   		emailAddress: jsii.String("emailAddress"),
//   		id: jsii.String("id"),
//   		name: jsii.String("name"),
//   	},
//   	delegations: []interface{}{
//   		&delegationProperty{
//   			assessmentId: jsii.String("assessmentId"),
//   			assessmentName: jsii.String("assessmentName"),
//   			comment: jsii.String("comment"),
//   			controlSetId: jsii.String("controlSetId"),
//   			createdBy: jsii.String("createdBy"),
//   			creationTime: jsii.Number(123),
//   			id: jsii.String("id"),
//   			lastUpdated: jsii.Number(123),
//   			roleArn: jsii.String("roleArn"),
//   			roleType: jsii.String("roleType"),
//   			status: jsii.String("status"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	frameworkId: jsii.String("frameworkId"),
//   	name: jsii.String("name"),
//   	roles: []interface{}{
//   		&roleProperty{
//   			roleArn: jsii.String("roleArn"),
//   			roleType: jsii.String("roleType"),
//   		},
//   	},
//   	scope: &scopeProperty{
//   		awsAccounts: []interface{}{
//   			&aWSAccountProperty{
//   				emailAddress: jsii.String("emailAddress"),
//   				id: jsii.String("id"),
//   				name: jsii.String("name"),
//   			},
//   		},
//   		awsServices: []interface{}{
//   			&aWSServiceProperty{
//   				serviceName: jsii.String("serviceName"),
//   			},
//   		},
//   	},
//   	status: jsii.String("status"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAssessmentProps struct {
	// The destination that evidence reports are stored in for the assessment.
	AssessmentReportsDestination interface{} `field:"optional" json:"assessmentReportsDestination" yaml:"assessmentReportsDestination"`
	// The AWS account that's associated with the assessment.
	AwsAccount interface{} `field:"optional" json:"awsAccount" yaml:"awsAccount"`
	// The delegations that are associated with the assessment.
	Delegations interface{} `field:"optional" json:"delegations" yaml:"delegations"`
	// The description of the assessment.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The unique identifier for the framework.
	FrameworkId *string `field:"optional" json:"frameworkId" yaml:"frameworkId"`
	// The name of the assessment.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The roles that are associated with the assessment.
	Roles interface{} `field:"optional" json:"roles" yaml:"roles"`
	// The wrapper of AWS accounts and services that are in scope for the assessment.
	Scope interface{} `field:"optional" json:"scope" yaml:"scope"`
	// The overall status of the assessment.
	//
	// When you create a new assessment, the initial `Status` value is always `ACTIVE` . When you create an assessment, even if you specify the value as `INACTIVE` , the value overrides to `ACTIVE` .
	//
	// After you create an assessment, you can change the value of the `Status` property at any time. For example, when you want to stop collecting evidence for your assessment, you can change the assessment status to `INACTIVE` .
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The tags that are associated with the assessment.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

