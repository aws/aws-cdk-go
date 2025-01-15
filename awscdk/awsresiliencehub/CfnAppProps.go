package awsresiliencehub


// Properties for defining a `CfnApp`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAppProps := &CfnAppProps{
//   	AppTemplateBody: jsii.String("appTemplateBody"),
//   	Name: jsii.String("name"),
//   	ResourceMappings: []interface{}{
//   		&ResourceMappingProperty{
//   			MappingType: jsii.String("mappingType"),
//   			PhysicalResourceId: &PhysicalResourceIdProperty{
//   				Identifier: jsii.String("identifier"),
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				AwsAccountId: jsii.String("awsAccountId"),
//   				AwsRegion: jsii.String("awsRegion"),
//   			},
//
//   			// the properties below are optional
//   			EksSourceName: jsii.String("eksSourceName"),
//   			LogicalStackName: jsii.String("logicalStackName"),
//   			ResourceName: jsii.String("resourceName"),
//   			TerraformSourceName: jsii.String("terraformSourceName"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	AppAssessmentSchedule: jsii.String("appAssessmentSchedule"),
//   	Description: jsii.String("description"),
//   	EventSubscriptions: []interface{}{
//   		&EventSubscriptionProperty{
//   			EventType: jsii.String("eventType"),
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			SnsTopicArn: jsii.String("snsTopicArn"),
//   		},
//   	},
//   	PermissionModel: &PermissionModelProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		CrossAccountRoleArns: []*string{
//   			jsii.String("crossAccountRoleArns"),
//   		},
//   		InvokerRoleName: jsii.String("invokerRoleName"),
//   	},
//   	RegulatoryPolicyArn: jsii.String("regulatoryPolicyArn"),
//   	ResiliencyPolicyArn: jsii.String("resiliencyPolicyArn"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehub-app.html
//
type CfnAppProps struct {
	// A JSON string that provides information about your application structure.
	//
	// To learn more about the `appTemplateBody` template, see the sample template in [Sample appTemplateBody template](https://docs.aws.amazon.com//resilience-hub/latest/APIReference/API_PutDraftAppVersionTemplate.html#API_PutDraftAppVersionTemplate_Examples) .
	//
	// The `appTemplateBody` JSON string has the following structure:
	//
	// - *`resources`*
	//
	// The list of logical resources that needs to be included in the AWS Resilience Hub application.
	//
	// Type: Array
	//
	// > Don't add the resources that you want to exclude.
	//
	// Each `resources` array item includes the following fields:
	//
	// - *`logicalResourceId`*
	//
	// The logical identifier of the resource.
	//
	// Type: Object
	//
	// Each `logicalResourceId` object includes the following fields:
	//
	// - `identifier`
	//
	// Identifier of the resource.
	//
	// Type: String
	// - `logicalStackName`
	//
	// Name of the AWS CloudFormation stack this resource belongs to.
	//
	// Type: String
	// - `resourceGroupName`
	//
	// Name of the resource group this resource belongs to.
	//
	// Type: String
	// - `terraformSourceName`
	//
	// Name of the Terraform S3 state file this resource belongs to.
	//
	// Type: String
	// - `eksSourceName`
	//
	// Name of the Amazon Elastic Kubernetes Service cluster and namespace this resource belongs to.
	//
	// > This parameter accepts values in "eks-cluster/namespace" format.
	//
	// Type: String
	// - *`type`*
	//
	// The type of resource.
	//
	// Type: string
	// - *`name`*
	//
	// Name of the resource.
	//
	// Type: String
	// - `additionalInfo`
	//
	// Additional configuration parameters for an AWS Resilience Hub application. If you want to implement `additionalInfo` through the AWS Resilience Hub console rather than using an API call, see [Configure the application configuration parameters](https://docs.aws.amazon.com//resilience-hub/latest/userguide/app-config-param.html) .
	//
	// > Currently, this parameter accepts a key-value mapping (in a string format) of only one failover region and one associated account.
	// >
	// > Key: `"failover-regions"`
	// >
	// > Value: `"[{"region":"<REGION>", "accounts":[{"id":"<ACCOUNT_ID>"}]}]"`
	// - *`appComponents`*
	//
	// The list of Application Components (AppComponent) that this resource belongs to. If an AppComponent is not part of the AWS Resilience Hub application, it will be added.
	//
	// Type: Array
	//
	// Each `appComponents` array item includes the following fields:
	//
	// - `name`
	//
	// Name of the AppComponent.
	//
	// Type: String
	// - `type`
	//
	// The type of AppComponent. For more information about the types of AppComponent, see [Grouping resources in an AppComponent](https://docs.aws.amazon.com/resilience-hub/latest/userguide/AppComponent.grouping.html) .
	//
	// Type: String
	// - `resourceNames`
	//
	// The list of included resources that are assigned to the AppComponent.
	//
	// Type: Array of strings
	// - `additionalInfo`
	//
	// Additional configuration parameters for an AWS Resilience Hub application. If you want to implement `additionalInfo` through the AWS Resilience Hub console rather than using an API call, see [Configure the application configuration parameters](https://docs.aws.amazon.com//resilience-hub/latest/userguide/app-config-param.html) .
	//
	// > Currently, this parameter accepts a key-value mapping (in a string format) of only one failover region and one associated account.
	// >
	// > Key: `"failover-regions"`
	// >
	// > Value: `"[{"region":"<REGION>", "accounts":[{"id":"<ACCOUNT_ID>"}]}]"`
	// - *`excludedResources`*
	//
	// The list of logical resource identifiers to be excluded from the application.
	//
	// Type: Array
	//
	// > Don't add the resources that you want to include.
	//
	// Each `excludedResources` array item includes the following fields:
	//
	// - *`logicalResourceIds`*
	//
	// The logical identifier of the resource.
	//
	// Type: Object
	//
	// > You can configure only one of the following fields:
	// >
	// > - `logicalStackName`
	// > - `resourceGroupName`
	// > - `terraformSourceName`
	// > - `eksSourceName`
	//
	// Each `logicalResourceIds` object includes the following fields:
	//
	// - `identifier`
	//
	// The identifier of the resource.
	//
	// Type: String
	// - `logicalStackName`
	//
	// Name of the AWS CloudFormation stack this resource belongs to.
	//
	// Type: String
	// - `resourceGroupName`
	//
	// Name of the resource group this resource belongs to.
	//
	// Type: String
	// - `terraformSourceName`
	//
	// Name of the Terraform S3 state file this resource belongs to.
	//
	// Type: String
	// - `eksSourceName`
	//
	// Name of the Amazon Elastic Kubernetes Service cluster and namespace this resource belongs to.
	//
	// > This parameter accepts values in "eks-cluster/namespace" format.
	//
	// Type: String
	// - *`version`*
	//
	// The AWS Resilience Hub application version.
	// - `additionalInfo`
	//
	// Additional configuration parameters for an AWS Resilience Hub application. If you want to implement `additionalInfo` through the AWS Resilience Hub console rather than using an API call, see [Configure the application configuration parameters](https://docs.aws.amazon.com//resilience-hub/latest/userguide/app-config-param.html) .
	//
	// > Currently, this parameter accepts a key-value mapping (in a string format) of only one failover region and one associated account.
	// >
	// > Key: `"failover-regions"`
	// >
	// > Value: `"[{"region":"<REGION>", "accounts":[{"id":"<ACCOUNT_ID>"}]}]"`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehub-app.html#cfn-resiliencehub-app-apptemplatebody
	//
	AppTemplateBody *string `field:"required" json:"appTemplateBody" yaml:"appTemplateBody"`
	// Name for the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehub-app.html#cfn-resiliencehub-app-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// An array of `ResourceMapping` objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehub-app.html#cfn-resiliencehub-app-resourcemappings
	//
	ResourceMappings interface{} `field:"required" json:"resourceMappings" yaml:"resourceMappings"`
	// Assessment execution schedule with 'Daily' or 'Disabled' values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehub-app.html#cfn-resiliencehub-app-appassessmentschedule
	//
	AppAssessmentSchedule *string `field:"optional" json:"appAssessmentSchedule" yaml:"appAssessmentSchedule"`
	// Optional description for an application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehub-app.html#cfn-resiliencehub-app-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The list of events you would like to subscribe and get notification for.
	//
	// Currently, AWS Resilience Hub supports notifications only for *Drift detected* and *Scheduled assessment failure* events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehub-app.html#cfn-resiliencehub-app-eventsubscriptions
	//
	EventSubscriptions interface{} `field:"optional" json:"eventSubscriptions" yaml:"eventSubscriptions"`
	// Defines the roles and credentials that AWS Resilience Hub would use while creating the application, importing its resources, and running an assessment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehub-app.html#cfn-resiliencehub-app-permissionmodel
	//
	PermissionModel interface{} `field:"optional" json:"permissionModel" yaml:"permissionModel"`
	// Amazon Resource Name (ARN) of the Regulatory Policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehub-app.html#cfn-resiliencehub-app-regulatorypolicyarn
	//
	RegulatoryPolicyArn *string `field:"optional" json:"regulatoryPolicyArn" yaml:"regulatoryPolicyArn"`
	// The Amazon Resource Name (ARN) of the resiliency policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehub-app.html#cfn-resiliencehub-app-resiliencypolicyarn
	//
	ResiliencyPolicyArn *string `field:"optional" json:"resiliencyPolicyArn" yaml:"resiliencyPolicyArn"`
	// Tags assigned to the resource.
	//
	// A tag is a label that you assign to an AWS resource. Each tag consists of a key/value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehub-app.html#cfn-resiliencehub-app-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

