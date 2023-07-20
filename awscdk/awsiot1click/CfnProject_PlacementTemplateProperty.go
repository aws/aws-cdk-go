package awsiot1click


// In AWS CloudFormation , use the `PlacementTemplate` property type to define the template for an AWS IoT 1-Click project.
//
// `PlacementTemplate` is a property of the `AWS::IoT1Click::Project` resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var callbackOverrides interface{}
//   var defaultAttributes interface{}
//
//   placementTemplateProperty := &PlacementTemplateProperty{
//   	DefaultAttributes: defaultAttributes,
//   	DeviceTemplates: map[string]interface{}{
//   		"deviceTemplatesKey": &DeviceTemplateProperty{
//   			"callbackOverrides": callbackOverrides,
//   			"deviceType": jsii.String("deviceType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot1click-project-placementtemplate.html
//
type CfnProject_PlacementTemplateProperty struct {
	// The default attributes (key-value pairs) to be applied to all placements using this template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot1click-project-placementtemplate.html#cfn-iot1click-project-placementtemplate-defaultattributes
	//
	DefaultAttributes interface{} `field:"optional" json:"defaultAttributes" yaml:"defaultAttributes"`
	// An object specifying the [DeviceTemplate](https://docs.aws.amazon.com/iot-1-click/latest/projects-apireference/API_DeviceTemplate.html) for all placements using this ( [PlacementTemplate](https://docs.aws.amazon.com/iot-1-click/latest/projects-apireference/API_PlacementTemplate.html) ) template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot1click-project-placementtemplate.html#cfn-iot1click-project-placementtemplate-devicetemplates
	//
	DeviceTemplates interface{} `field:"optional" json:"deviceTemplates" yaml:"deviceTemplates"`
}

