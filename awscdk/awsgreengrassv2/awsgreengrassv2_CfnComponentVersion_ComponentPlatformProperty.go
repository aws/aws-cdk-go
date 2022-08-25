package awsgreengrassv2


// Contains information about a platform that a component supports.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   componentPlatformProperty := &componentPlatformProperty{
//   	attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   	name: jsii.String("name"),
//   }
//
type CfnComponentVersion_ComponentPlatformProperty struct {
	// A dictionary of attributes for the platform.
	//
	// The  software defines the `os` and `platform` by default. You can specify additional platform attributes for a core device when you deploy the Greengrass nucleus component. For more information, see the [Greengrass nucleus component](https://docs.aws.amazon.com/greengrass/v2/developerguide/greengrass-nucleus-component.html) in the *AWS IoT Greengrass V2 Developer Guide* .
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// The friendly name of the platform. This name helps you identify the platform.
	//
	// If you omit this parameter, AWS IoT Greengrass creates a friendly name from the `os` and `architecture` of the platform.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

