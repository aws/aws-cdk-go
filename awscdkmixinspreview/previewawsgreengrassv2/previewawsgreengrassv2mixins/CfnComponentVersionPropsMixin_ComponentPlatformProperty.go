package previewawsgreengrassv2mixins


// Contains information about a platform that a component supports.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   componentPlatformProperty := &ComponentPlatformProperty{
//   	Attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-componentplatform.html
//
type CfnComponentVersionPropsMixin_ComponentPlatformProperty struct {
	// A dictionary of attributes for the platform.
	//
	// The AWS IoT Greengrass Core software defines the `os` and `platform` by default. You can specify additional platform attributes for a core device when you deploy the AWS IoT Greengrass nucleus component. For more information, see the [AWS IoT Greengrass nucleus component](https://docs.aws.amazon.com/greengrass/v2/developerguide/greengrass-nucleus-component.html) in the *AWS IoT Greengrass V2 Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-componentplatform.html#cfn-greengrassv2-componentversion-componentplatform-attributes
	//
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// The friendly name of the platform. This name helps you identify the platform.
	//
	// If you omit this parameter, AWS IoT Greengrass creates a friendly name from the `os` and `architecture` of the platform.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-componentplatform.html#cfn-greengrassv2-componentversion-componentplatform-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

