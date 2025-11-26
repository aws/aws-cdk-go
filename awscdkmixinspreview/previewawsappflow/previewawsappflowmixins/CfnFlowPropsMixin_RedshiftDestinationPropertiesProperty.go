package previewawsappflowmixins


// The properties that are applied when Amazon Redshift is being used as a destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   redshiftDestinationPropertiesProperty := &RedshiftDestinationPropertiesProperty{
//   	BucketPrefix: jsii.String("bucketPrefix"),
//   	ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   		BucketName: jsii.String("bucketName"),
//   		BucketPrefix: jsii.String("bucketPrefix"),
//   		FailOnFirstError: jsii.Boolean(false),
//   	},
//   	IntermediateBucketName: jsii.String("intermediateBucketName"),
//   	Object: jsii.String("object"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-redshiftdestinationproperties.html
//
type CfnFlowPropsMixin_RedshiftDestinationPropertiesProperty struct {
	// The object key for the bucket in which Amazon AppFlow places the destination files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-redshiftdestinationproperties.html#cfn-appflow-flow-redshiftdestinationproperties-bucketprefix
	//
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// The settings that determine how Amazon AppFlow handles an error when placing data in the Amazon Redshift destination.
	//
	// For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure. `ErrorHandlingConfig` is a part of the destination connector details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-redshiftdestinationproperties.html#cfn-appflow-flow-redshiftdestinationproperties-errorhandlingconfig
	//
	ErrorHandlingConfig interface{} `field:"optional" json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
	// The intermediate bucket that Amazon AppFlow uses when moving data into Amazon Redshift.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-redshiftdestinationproperties.html#cfn-appflow-flow-redshiftdestinationproperties-intermediatebucketname
	//
	IntermediateBucketName *string `field:"optional" json:"intermediateBucketName" yaml:"intermediateBucketName"`
	// The object specified in the Amazon Redshift flow destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-redshiftdestinationproperties.html#cfn-appflow-flow-redshiftdestinationproperties-object
	//
	Object *string `field:"optional" json:"object" yaml:"object"`
}

