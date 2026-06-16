package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   propagateTagsConfigProperty := &PropagateTagsConfigProperty{
//   	ExplicitTags: []interface{}{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Mode: jsii.String("mode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-propagatetagsconfig.html
//
type CfnCapacityProviderPropsMixin_PropagateTagsConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-propagatetagsconfig.html#cfn-lambda-capacityprovider-propagatetagsconfig-explicittags
	//
	ExplicitTags interface{} `field:"optional" json:"explicitTags" yaml:"explicitTags"`
	// The mode for tag propagation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-propagatetagsconfig.html#cfn-lambda-capacityprovider-propagatetagsconfig-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

