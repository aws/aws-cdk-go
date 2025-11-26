package previewawsbatchmixins


// Specifies the weights for the share identifiers for the fair-share policy.
//
// Share identifiers that aren't included have a default weight of `1.0` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   shareAttributesProperty := &ShareAttributesProperty{
//   	ShareIdentifier: jsii.String("shareIdentifier"),
//   	WeightFactor: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-schedulingpolicy-shareattributes.html
//
type CfnSchedulingPolicyPropsMixin_ShareAttributesProperty struct {
	// A share identifier or share identifier prefix.
	//
	// If the string ends with an asterisk (*), this entry specifies the weight factor to use for share identifiers that start with that prefix. The list of share identifiers in a fair-share policy can't overlap. For example, you can't have one that specifies a `shareIdentifier` of `UserA*` and another that specifies a `shareIdentifier` of `UserA1` .
	//
	// There can be no more than 500 share identifiers active in a job queue.
	//
	// The string is limited to 255 alphanumeric characters, and can be followed by an asterisk (*).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-schedulingpolicy-shareattributes.html#cfn-batch-schedulingpolicy-shareattributes-shareidentifier
	//
	ShareIdentifier *string `field:"optional" json:"shareIdentifier" yaml:"shareIdentifier"`
	// The weight factor for the share identifier.
	//
	// The default value is 1.0. A lower value has a higher priority for compute resources. For example, jobs that use a share identifier with a weight factor of 0.125 (1/8) get 8 times the compute resources of jobs that use a share identifier with a weight factor of 1.
	//
	// The smallest supported value is 0.0001, and the largest supported value is 999.9999.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-schedulingpolicy-shareattributes.html#cfn-batch-schedulingpolicy-shareattributes-weightfactor
	//
	WeightFactor *float64 `field:"optional" json:"weightFactor" yaml:"weightFactor"`
}

