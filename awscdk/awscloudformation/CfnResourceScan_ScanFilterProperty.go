package awscloudformation


// A filter that is used to specify which resource types to scan.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scanFilterProperty := &ScanFilterProperty{
//   	Types: []*string{
//   		jsii.String("types"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-resourcescan-scanfilter.html
//
type CfnResourceScan_ScanFilterProperty struct {
	// An array of strings where each string represents an AWS resource type to scan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-resourcescan-scanfilter.html#cfn-cloudformation-resourcescan-scanfilter-types
	//
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

