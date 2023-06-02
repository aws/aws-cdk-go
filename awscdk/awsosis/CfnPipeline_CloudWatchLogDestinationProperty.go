package awsosis


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogDestinationProperty := &CloudWatchLogDestinationProperty{
//   	LogGroup: jsii.String("logGroup"),
//   }
//
type CfnPipeline_CloudWatchLogDestinationProperty struct {
	// `CfnPipeline.CloudWatchLogDestinationProperty.LogGroup`.
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}

