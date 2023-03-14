package awsiot1click


// Properties for defining a `CfnProject`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var callbackOverrides interface{}
//   var defaultAttributes interface{}
//
//   cfnProjectProps := &cfnProjectProps{
//   	placementTemplate: &placementTemplateProperty{
//   		defaultAttributes: defaultAttributes,
//   		deviceTemplates: map[string]interface{}{
//   			"deviceTemplatesKey": &DeviceTemplateProperty{
//   				"callbackOverrides": callbackOverrides,
//   				"deviceType": jsii.String("deviceType"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	projectName: jsii.String("projectName"),
//   }
//
type CfnProjectProps struct {
	// An object describing the project's placement specifications.
	PlacementTemplate interface{} `field:"required" json:"placementTemplate" yaml:"placementTemplate"`
	// The description of the project.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the project from which to obtain information.
	ProjectName *string `field:"optional" json:"projectName" yaml:"projectName"`
}

