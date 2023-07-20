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
//   cfnProjectProps := &CfnProjectProps{
//   	PlacementTemplate: &PlacementTemplateProperty{
//   		DefaultAttributes: defaultAttributes,
//   		DeviceTemplates: map[string]interface{}{
//   			"deviceTemplatesKey": &DeviceTemplateProperty{
//   				"callbackOverrides": callbackOverrides,
//   				"deviceType": jsii.String("deviceType"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	ProjectName: jsii.String("projectName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-project.html
//
type CfnProjectProps struct {
	// An object describing the project's placement specifications.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-project.html#cfn-iot1click-project-placementtemplate
	//
	PlacementTemplate interface{} `field:"required" json:"placementTemplate" yaml:"placementTemplate"`
	// The description of the project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-project.html#cfn-iot1click-project-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the project from which to obtain information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-project.html#cfn-iot1click-project-projectname
	//
	ProjectName *string `field:"optional" json:"projectName" yaml:"projectName"`
}

