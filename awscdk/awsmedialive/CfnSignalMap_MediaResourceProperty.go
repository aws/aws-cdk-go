package awsmedialive


// An Amazon Web Services resource used in media workflows.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mediaResourceProperty := &MediaResourceProperty{
//   	Destinations: []interface{}{
//   		&MediaResourceNeighborProperty{
//   			Arn: jsii.String("arn"),
//
//   			// the properties below are optional
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Sources: []interface{}{
//   		&MediaResourceNeighborProperty{
//   			Arn: jsii.String("arn"),
//
//   			// the properties below are optional
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-signalmap-mediaresource.html
//
type CfnSignalMap_MediaResourceProperty struct {
	// A direct destination neighbor to an Amazon Web Services media resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-signalmap-mediaresource.html#cfn-medialive-signalmap-mediaresource-destinations
	//
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// The logical name of an Amazon Web Services media resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-signalmap-mediaresource.html#cfn-medialive-signalmap-mediaresource-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A direct source neighbor to an Amazon Web Services media resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-signalmap-mediaresource.html#cfn-medialive-signalmap-mediaresource-sources
	//
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
}

