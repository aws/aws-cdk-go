package mixinsawsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDimensionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDimensionMixinProps := &CfnDimensionMixinProps{
//   	Name: jsii.String("name"),
//   	StringValues: []*string{
//   		jsii.String("stringValues"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-dimension.html
//
type CfnDimensionMixinProps struct {
	// A unique identifier for the dimension.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-dimension.html#cfn-iot-dimension-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the value or list of values for the dimension.
	//
	// For `TOPIC_FILTER` dimensions, this is a pattern used to match the MQTT topic (for example, "admin/#").
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-dimension.html#cfn-iot-dimension-stringvalues
	//
	StringValues *[]*string `field:"optional" json:"stringValues" yaml:"stringValues"`
	// Metadata that can be used to manage the dimension.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-dimension.html#cfn-iot-dimension-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies the type of dimension.
	//
	// Supported types: `TOPIC_FILTER.`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-dimension.html#cfn-iot-dimension-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

