package previewawscustomerprofilesmixins


// Usage-specific statistics about the domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   domainStatsProperty := &DomainStatsProperty{
//   	MeteringProfileCount: jsii.Number(123),
//   	ObjectCount: jsii.Number(123),
//   	ProfileCount: jsii.Number(123),
//   	TotalSize: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-domainstats.html
//
type CfnDomainPropsMixin_DomainStatsProperty struct {
	// The number of profiles that you are currently paying for in the domain.
	//
	// If you have more than 100 objects associated with a single profile, that profile counts as two profiles. If you have more than 200 objects, that profile counts as three, and so on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-domainstats.html#cfn-customerprofiles-domain-domainstats-meteringprofilecount
	//
	MeteringProfileCount *float64 `field:"optional" json:"meteringProfileCount" yaml:"meteringProfileCount"`
	// The total number of objects in domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-domainstats.html#cfn-customerprofiles-domain-domainstats-objectcount
	//
	ObjectCount *float64 `field:"optional" json:"objectCount" yaml:"objectCount"`
	// The total number of profiles currently in the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-domainstats.html#cfn-customerprofiles-domain-domainstats-profilecount
	//
	ProfileCount *float64 `field:"optional" json:"profileCount" yaml:"profileCount"`
	// The total size, in bytes, of all objects in the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-domainstats.html#cfn-customerprofiles-domain-domainstats-totalsize
	//
	TotalSize *float64 `field:"optional" json:"totalSize" yaml:"totalSize"`
}

