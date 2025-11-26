package previewawscustomerprofilesmixins


// Progress information for data store setup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   readinessProperty := &ReadinessProperty{
//   	Message: jsii.String("message"),
//   	ProgressPercentage: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-readiness.html
//
type CfnDomainPropsMixin_ReadinessProperty struct {
	// A message describing the current progress.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-readiness.html#cfn-customerprofiles-domain-readiness-message
	//
	Message *string `field:"optional" json:"message" yaml:"message"`
	// The percentage of progress completed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-readiness.html#cfn-customerprofiles-domain-readiness-progresspercentage
	//
	ProgressPercentage *float64 `field:"optional" json:"progressPercentage" yaml:"progressPercentage"`
}

