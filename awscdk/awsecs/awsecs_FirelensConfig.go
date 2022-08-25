package awsecs


// Firelens Configuration https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html#firelens-taskdef.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firelensConfig := &firelensConfig{
//   	type: awscdk.Aws_ecs.firelensLogRouterType_FLUENTBIT,
//
//   	// the properties below are optional
//   	options: &firelensOptions{
//   		configFileType: awscdk.*Aws_ecs.firelensConfigFileType_S3,
//   		configFileValue: jsii.String("configFileValue"),
//   		enableECSLogMetadata: jsii.Boolean(false),
//   	},
//   }
//
type FirelensConfig struct {
	// The log router to use.
	Type FirelensLogRouterType `field:"required" json:"type" yaml:"type"`
	// Firelens options.
	Options *FirelensOptions `field:"optional" json:"options" yaml:"options"`
}

