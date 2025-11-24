package mixinsawsresourceexplorer2


// Information about an additional property that describes a resource, that you can optionally include in a view.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   includedPropertyProperty := &IncludedPropertyProperty{
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resourceexplorer2-view-includedproperty.html
//
type CfnViewPropsMixin_IncludedPropertyProperty struct {
	// The name of the property that is included in this view.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resourceexplorer2-view-includedproperty.html#cfn-resourceexplorer2-view-includedproperty-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

