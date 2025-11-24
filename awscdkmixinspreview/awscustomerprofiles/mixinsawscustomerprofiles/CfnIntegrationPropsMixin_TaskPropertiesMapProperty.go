package mixinsawscustomerprofiles


// A map used to store task-related information.
//
// The execution service looks for particular information based on the `TaskType` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   taskPropertiesMapProperty := &TaskPropertiesMapProperty{
//   	OperatorPropertyKey: jsii.String("operatorPropertyKey"),
//   	Property: jsii.String("property"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-taskpropertiesmap.html
//
type CfnIntegrationPropsMixin_TaskPropertiesMapProperty struct {
	// The task property key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-taskpropertiesmap.html#cfn-customerprofiles-integration-taskpropertiesmap-operatorpropertykey
	//
	OperatorPropertyKey *string `field:"optional" json:"operatorPropertyKey" yaml:"operatorPropertyKey"`
	// The task property value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-taskpropertiesmap.html#cfn-customerprofiles-integration-taskpropertiesmap-property
	//
	Property *string `field:"optional" json:"property" yaml:"property"`
}

