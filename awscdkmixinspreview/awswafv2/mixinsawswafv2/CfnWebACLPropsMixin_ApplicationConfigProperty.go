package mixinsawswafv2


// A list of `ApplicationAttribute` s that contains information about the application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   applicationConfigProperty := &ApplicationConfigProperty{
//   	Attributes: []interface{}{
//   		&ApplicationAttributeProperty{
//   			Name: jsii.String("name"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-applicationconfig.html
//
type CfnWebACLPropsMixin_ApplicationConfigProperty struct {
	// Contains the attribute name and a list of values for that attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-applicationconfig.html#cfn-wafv2-webacl-applicationconfig-attributes
	//
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
}

