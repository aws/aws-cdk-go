package awsmediatailor


// <p>Slate VOD source configuration.</p>.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slateSourceProperty := &SlateSourceProperty{
//   	SourceLocationName: jsii.String("sourceLocationName"),
//   	VodSourceName: jsii.String("vodSourceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-slatesource.html
//
type CfnChannel_SlateSourceProperty struct {
	// <p>The name of the source location where the slate VOD source is stored.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-slatesource.html#cfn-mediatailor-channel-slatesource-sourcelocationname
	//
	SourceLocationName *string `field:"optional" json:"sourceLocationName" yaml:"sourceLocationName"`
	// <p>The slate VOD source name.
	//
	// The VOD source must already exist in a source location before it can be used for slate.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-slatesource.html#cfn-mediatailor-channel-slatesource-vodsourcename
	//
	VodSourceName *string `field:"optional" json:"vodSourceName" yaml:"vodSourceName"`
}

