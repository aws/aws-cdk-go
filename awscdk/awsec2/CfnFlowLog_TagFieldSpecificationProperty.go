package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagFieldSpecificationProperty := &TagFieldSpecificationProperty{
//   	ResourceType: jsii.String("resourceType"),
//   	TagKeys: []*string{
//   		jsii.String("tagKeys"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-flowlog-tagfieldspecification.html
//
type CfnFlowLog_TagFieldSpecificationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-flowlog-tagfieldspecification.html#cfn-ec2-flowlog-tagfieldspecification-resourcetype
	//
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-flowlog-tagfieldspecification.html#cfn-ec2-flowlog-tagfieldspecification-tagkeys
	//
	TagKeys *[]*string `field:"required" json:"tagKeys" yaml:"tagKeys"`
}

