package awsresiliencehub


// Properties for defining a `CfnApp`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAppProps := &cfnAppProps{
//   	appTemplateBody: jsii.String("appTemplateBody"),
//   	name: jsii.String("name"),
//   	resourceMappings: []interface{}{
//   		&resourceMappingProperty{
//   			mappingType: jsii.String("mappingType"),
//   			physicalResourceId: &physicalResourceIdProperty{
//   				identifier: jsii.String("identifier"),
//   				type: jsii.String("type"),
//
//   				// the properties below are optional
//   				awsAccountId: jsii.String("awsAccountId"),
//   				awsRegion: jsii.String("awsRegion"),
//   			},
//
//   			// the properties below are optional
//   			logicalStackName: jsii.String("logicalStackName"),
//   			resourceName: jsii.String("resourceName"),
//   			terraformSourceName: jsii.String("terraformSourceName"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	appAssessmentSchedule: jsii.String("appAssessmentSchedule"),
//   	description: jsii.String("description"),
//   	resiliencyPolicyArn: jsii.String("resiliencyPolicyArn"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnAppProps struct {
	// A string containing a full Resilience Hub app template body.
	AppTemplateBody *string `field:"required" json:"appTemplateBody" yaml:"appTemplateBody"`
	// The name for the application.
	Name *string `field:"required" json:"name" yaml:"name"`
	// An array of ResourceMapping objects.
	ResourceMappings interface{} `field:"required" json:"resourceMappings" yaml:"resourceMappings"`
	// `AWS::ResilienceHub::App.AppAssessmentSchedule`.
	AppAssessmentSchedule *string `field:"optional" json:"appAssessmentSchedule" yaml:"appAssessmentSchedule"`
	// The optional description for an app.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of the resiliency policy.
	ResiliencyPolicyArn *string `field:"optional" json:"resiliencyPolicyArn" yaml:"resiliencyPolicyArn"`
	// The tags assigned to the resource.
	//
	// A tag is a label that you assign to an AWS resource. Each tag consists of a key/value pair.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

