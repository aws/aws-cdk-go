package mixinsawsmedialive


// A direct source or destination neighbor to an Amazon Web Services media resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mediaResourceNeighborProperty := &MediaResourceNeighborProperty{
//   	Arn: jsii.String("arn"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-signalmap-mediaresourceneighbor.html
//
type CfnSignalMapPropsMixin_MediaResourceNeighborProperty struct {
	// The ARN of a resource used in Amazon Web Services media workflows.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-signalmap-mediaresourceneighbor.html#cfn-medialive-signalmap-mediaresourceneighbor-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The logical name of an Amazon Web Services media resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-signalmap-mediaresourceneighbor.html#cfn-medialive-signalmap-mediaresourceneighbor-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

