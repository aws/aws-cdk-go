package awsresiliencehub


// Defines a resource mapping.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceMappingProperty := &resourceMappingProperty{
//   	mappingType: jsii.String("mappingType"),
//   	physicalResourceId: &physicalResourceIdProperty{
//   		identifier: jsii.String("identifier"),
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		awsAccountId: jsii.String("awsAccountId"),
//   		awsRegion: jsii.String("awsRegion"),
//   	},
//
//   	// the properties below are optional
//   	logicalStackName: jsii.String("logicalStackName"),
//   	resourceName: jsii.String("resourceName"),
//   	terraformSourceName: jsii.String("terraformSourceName"),
//   }
//
type CfnApp_ResourceMappingProperty struct {
	// Specifies the type of resource mapping.
	//
	// - **AppRegistryApp** - The resource is mapped to another application. The name of the application is contained in the `appRegistryAppName` property.
	// - **CfnStack** - The resource is mapped to a CloudFormation stack. The name of the CloudFormation stack is contained in the `logicalStackName` property.
	// - **Resource** - The resource is mapped to another resource. The name of the resource is contained in the `resourceName` property.
	// - **ResourceGroup** - The resource is mapped to a resource group. The name of the resource group is contained in the `resourceGroupName` property.
	MappingType *string `field:"required" json:"mappingType" yaml:"mappingType"`
	// The identifier of this resource.
	PhysicalResourceId interface{} `field:"required" json:"physicalResourceId" yaml:"physicalResourceId"`
	// The name of the CloudFormation stack this resource is mapped to.
	LogicalStackName *string `field:"optional" json:"logicalStackName" yaml:"logicalStackName"`
	// The name of the resource this resource is mapped to.
	ResourceName *string `field:"optional" json:"resourceName" yaml:"resourceName"`
	// `CfnApp.ResourceMappingProperty.TerraformSourceName`.
	TerraformSourceName *string `field:"optional" json:"terraformSourceName" yaml:"terraformSourceName"`
}

