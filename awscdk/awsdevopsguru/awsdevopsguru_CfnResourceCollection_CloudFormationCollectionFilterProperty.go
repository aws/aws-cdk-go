package awsdevopsguru


// Information about AWS CloudFormation stacks.
//
// You can use up to 500 stacks to specify which AWS resources in your account to analyze. For more information, see [Stacks](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacks.html) in the *AWS CloudFormation User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudFormationCollectionFilterProperty := &cloudFormationCollectionFilterProperty{
//   	stackNames: []*string{
//   		jsii.String("stackNames"),
//   	},
//   }
//
type CfnResourceCollection_CloudFormationCollectionFilterProperty struct {
	// An array of CloudFormation stack names.
	StackNames *[]*string `field:"optional" json:"stackNames" yaml:"stackNames"`
}

