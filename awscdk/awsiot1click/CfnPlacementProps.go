package awsiot1click


// Properties for defining a `CfnPlacement`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var associatedDevices interface{}
//   var attributes interface{}
//
//   cfnPlacementProps := &CfnPlacementProps{
//   	ProjectName: jsii.String("projectName"),
//
//   	// the properties below are optional
//   	AssociatedDevices: associatedDevices,
//   	Attributes: attributes,
//   	PlacementName: jsii.String("placementName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-placement.html
//
type CfnPlacementProps struct {
	// The name of the project containing the placement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-placement.html#cfn-iot1click-placement-projectname
	//
	ProjectName *string `field:"required" json:"projectName" yaml:"projectName"`
	// The devices to associate with the placement, as defined by a mapping of zero or more key-value pairs wherein the key is a template name and the value is a device ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-placement.html#cfn-iot1click-placement-associateddevices
	//
	AssociatedDevices interface{} `field:"optional" json:"associatedDevices" yaml:"associatedDevices"`
	// The user-defined attributes associated with the placement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-placement.html#cfn-iot1click-placement-attributes
	//
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// The name of the placement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-placement.html#cfn-iot1click-placement-placementname
	//
	PlacementName *string `field:"optional" json:"placementName" yaml:"placementName"`
}

