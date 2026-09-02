package awssagemaker


// A list of IP address ranges used to access your training data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceIpConfigProperty := &SourceIpConfigProperty{
//   	Cidrs: []*string{
//   		jsii.String("cidrs"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workforce-sourceipconfig.html
//
type CfnWorkforce_SourceIpConfigProperty struct {
	// A list of one to ten Classless Inter-Domain Routing (CIDR) values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workforce-sourceipconfig.html#cfn-sagemaker-workforce-sourceipconfig-cidrs
	//
	Cidrs *[]*string `field:"required" json:"cidrs" yaml:"cidrs"`
}

