package awsresiliencehub


// Defines a resource mapping.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceMappingProperty := &ResourceMappingProperty{
//   	MappingType: jsii.String("mappingType"),
//   	PhysicalResourceId: &PhysicalResourceIdProperty{
//   		Identifier: jsii.String("identifier"),
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		AwsAccountId: jsii.String("awsAccountId"),
//   		AwsRegion: jsii.String("awsRegion"),
//   	},
//
//   	// the properties below are optional
//   	EksSourceName: jsii.String("eksSourceName"),
//   	LogicalStackName: jsii.String("logicalStackName"),
//   	ResourceName: jsii.String("resourceName"),
//   	TerraformSourceName: jsii.String("terraformSourceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-resourcemapping.html
//
type CfnApp_ResourceMappingProperty struct {
	// Specifies the type of resource mapping.
	//
	// Valid Values: CfnStack | Resource | AppRegistryApp | ResourceGroup | Terraform
	//
	// - **AppRegistryApp** - The resource is mapped to another application. The name of the application is contained in the `appRegistryAppName` property.
	// - **CfnStack** - The resource is mapped to a CloudFormation stack. The name of the CloudFormation stack is contained in the `logicalStackName` property.
	// - **Resource** - The resource is mapped to another resource. The name of the resource is contained in the `resourceName` property.
	// - **ResourceGroup** - The resource is mapped to a resource group. The name of the resource group is contained in the `resourceGroupName` property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-resourcemapping.html#cfn-resiliencehub-app-resourcemapping-mappingtype
	//
	MappingType *string `field:"required" json:"mappingType" yaml:"mappingType"`
	// The identifier of this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-resourcemapping.html#cfn-resiliencehub-app-resourcemapping-physicalresourceid
	//
	PhysicalResourceId interface{} `field:"required" json:"physicalResourceId" yaml:"physicalResourceId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-resourcemapping.html#cfn-resiliencehub-app-resourcemapping-ekssourcename
	//
	EksSourceName *string `field:"optional" json:"eksSourceName" yaml:"eksSourceName"`
	// The name of the CloudFormation stack this resource is mapped to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-resourcemapping.html#cfn-resiliencehub-app-resourcemapping-logicalstackname
	//
	LogicalStackName *string `field:"optional" json:"logicalStackName" yaml:"logicalStackName"`
	// The name of the resource this resource is mapped to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-resourcemapping.html#cfn-resiliencehub-app-resourcemapping-resourcename
	//
	ResourceName *string `field:"optional" json:"resourceName" yaml:"resourceName"`
	// The short name of the Terraform source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-resourcemapping.html#cfn-resiliencehub-app-resourcemapping-terraformsourcename
	//
	TerraformSourceName *string `field:"optional" json:"terraformSourceName" yaml:"terraformSourceName"`
}

