package mixinsawsdax


// Properties for CfnParameterGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var parameterNameValues interface{}
//
//   cfnParameterGroupMixinProps := &CfnParameterGroupMixinProps{
//   	Description: jsii.String("description"),
//   	ParameterGroupName: jsii.String("parameterGroupName"),
//   	ParameterNameValues: parameterNameValues,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-parametergroup.html
//
type CfnParameterGroupMixinProps struct {
	// A description of the parameter group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-parametergroup.html#cfn-dax-parametergroup-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the parameter group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-parametergroup.html#cfn-dax-parametergroup-parametergroupname
	//
	ParameterGroupName *string `field:"optional" json:"parameterGroupName" yaml:"parameterGroupName"`
	// An array of name-value pairs for the parameters in the group.
	//
	// Each element in the array represents a single parameter.
	//
	// > `record-ttl-millis` and `query-ttl-millis` are the only supported parameter names. For more details, see [Configuring TTL Settings](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DAX.cluster-management.html#DAX.cluster-management.custom-settings.ttl) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-parametergroup.html#cfn-dax-parametergroup-parameternamevalues
	//
	ParameterNameValues interface{} `field:"optional" json:"parameterNameValues" yaml:"parameterNameValues"`
}

