package awsdevicefarm

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDevicePool`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDevicePoolProps := &CfnDevicePoolProps{
//   	Name: jsii.String("name"),
//   	ProjectArn: jsii.String("projectArn"),
//   	Rules: []interface{}{
//   		&RuleProperty{
//   			Attribute: jsii.String("attribute"),
//   			Operator: jsii.String("operator"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	MaxDevices: jsii.Number(123),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDevicePoolProps struct {
	// The device pool's name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of the project for the device pool.
	ProjectArn *string `field:"required" json:"projectArn" yaml:"projectArn"`
	// The device pool's rules.
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// The device pool's description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The number of devices that Device Farm can add to your device pool.
	//
	// Device Farm adds devices that are available and meet the criteria that you assign for the `rules` parameter. Depending on how many devices meet these constraints, your device pool might contain fewer devices than the value for this parameter.
	//
	// By specifying the maximum number of devices, you can control the costs that you incur by running tests.
	MaxDevices *float64 `field:"optional" json:"maxDevices" yaml:"maxDevices"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

