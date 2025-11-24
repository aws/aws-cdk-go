package mixinsawsmsk


// Describes a configuration revision.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   latestRevisionProperty := &LatestRevisionProperty{
//   	CreationTime: jsii.String("creationTime"),
//   	Description: jsii.String("description"),
//   	Revision: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-configuration-latestrevision.html
//
type CfnConfigurationPropsMixin_LatestRevisionProperty struct {
	// The time when the configuration revision was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-configuration-latestrevision.html#cfn-msk-configuration-latestrevision-creationtime
	//
	CreationTime *string `field:"optional" json:"creationTime" yaml:"creationTime"`
	// The description of the configuration revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-configuration-latestrevision.html#cfn-msk-configuration-latestrevision-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The revision number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-configuration-latestrevision.html#cfn-msk-configuration-latestrevision-revision
	//
	Revision *float64 `field:"optional" json:"revision" yaml:"revision"`
}

