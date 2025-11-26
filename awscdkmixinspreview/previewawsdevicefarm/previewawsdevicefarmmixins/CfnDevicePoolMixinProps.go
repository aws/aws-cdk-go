package previewawsdevicefarmmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDevicePoolPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDevicePoolMixinProps := &CfnDevicePoolMixinProps{
//   	Description: jsii.String("description"),
//   	MaxDevices: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	ProjectArn: jsii.String("projectArn"),
//   	Rules: []interface{}{
//   		&RuleProperty{
//   			Attribute: jsii.String("attribute"),
//   			Operator: jsii.String("operator"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-devicepool.html
//
type CfnDevicePoolMixinProps struct {
	// The device pool's description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-devicepool.html#cfn-devicefarm-devicepool-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The number of devices that Device Farm can add to your device pool.
	//
	// Device Farm adds devices that are available and meet the criteria that you assign for the `rules` parameter. Depending on how many devices meet these constraints, your device pool might contain fewer devices than the value for this parameter.
	//
	// By specifying the maximum number of devices, you can control the costs that you incur by running tests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-devicepool.html#cfn-devicefarm-devicepool-maxdevices
	//
	MaxDevices *float64 `field:"optional" json:"maxDevices" yaml:"maxDevices"`
	// The device pool's name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-devicepool.html#cfn-devicefarm-devicepool-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ARN of the project for the device pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-devicepool.html#cfn-devicefarm-devicepool-projectarn
	//
	ProjectArn *string `field:"optional" json:"projectArn" yaml:"projectArn"`
	// The device pool's rules.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-devicepool.html#cfn-devicefarm-devicepool-rules
	//
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-devicepool.html#cfn-devicefarm-devicepool-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

