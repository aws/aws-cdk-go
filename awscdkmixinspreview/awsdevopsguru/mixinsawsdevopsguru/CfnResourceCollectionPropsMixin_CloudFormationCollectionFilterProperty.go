package mixinsawsdevopsguru


// Information about AWS CloudFormation stacks.
//
// You can use up to 1000 stacks to specify which AWS resources in your account to analyze. For more information, see [Stacks](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacks.html) in the *AWS CloudFormation User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cloudFormationCollectionFilterProperty := &CloudFormationCollectionFilterProperty{
//   	StackNames: []*string{
//   		jsii.String("stackNames"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsguru-resourcecollection-cloudformationcollectionfilter.html
//
type CfnResourceCollectionPropsMixin_CloudFormationCollectionFilterProperty struct {
	// An array of CloudFormation stack names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsguru-resourcecollection-cloudformationcollectionfilter.html#cfn-devopsguru-resourcecollection-cloudformationcollectionfilter-stacknames
	//
	StackNames *[]*string `field:"optional" json:"stackNames" yaml:"stackNames"`
}

