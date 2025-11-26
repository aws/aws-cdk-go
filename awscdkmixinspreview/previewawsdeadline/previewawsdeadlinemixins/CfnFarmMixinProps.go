package previewawsdeadlinemixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnFarmPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFarmMixinProps := &CfnFarmMixinProps{
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-farm.html
//
type CfnFarmMixinProps struct {
	// A description of the farm that helps identify what the farm is used for.
	//
	// > This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-farm.html#cfn-deadline-farm-description
	//
	// Default: - "".
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name of the farm.
	//
	// > This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-farm.html#cfn-deadline-farm-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The ARN for the KMS key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-farm.html#cfn-deadline-farm-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The tags to add to your farm.
	//
	// Each tag consists of a tag key and a tag value. Tag keys and values are both required, but tag values can be empty strings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-farm.html#cfn-deadline-farm-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

