package previewawsmemorydbmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnParameterGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var parameters interface{}
//
//   cfnParameterGroupMixinProps := &CfnParameterGroupMixinProps{
//   	Description: jsii.String("description"),
//   	Family: jsii.String("family"),
//   	ParameterGroupName: jsii.String("parameterGroupName"),
//   	Parameters: parameters,
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-parametergroup.html
//
type CfnParameterGroupMixinProps struct {
	// A description of the parameter group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-parametergroup.html#cfn-memorydb-parametergroup-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the parameter group family that this parameter group is compatible with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-parametergroup.html#cfn-memorydb-parametergroup-family
	//
	Family *string `field:"optional" json:"family" yaml:"family"`
	// The name of the parameter group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-parametergroup.html#cfn-memorydb-parametergroup-parametergroupname
	//
	ParameterGroupName *string `field:"optional" json:"parameterGroupName" yaml:"parameterGroupName"`
	// Returns the detailed parameter list for the parameter group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-parametergroup.html#cfn-memorydb-parametergroup-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-parametergroup.html#cfn-memorydb-parametergroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

