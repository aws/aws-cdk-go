package awsecs


// Firelens Configuration https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html#firelens-taskdef.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firelensConfig := &FirelensConfig{
//   	Type: awscdk.Aws_ecs.FirelensLogRouterType_FLUENTBIT,
//
//   	// the properties below are optional
//   	Options: &FirelensOptions{
//   		ConfigFileType: awscdk.*Aws_ecs.FirelensConfigFileType_S3,
//   		ConfigFileValue: jsii.String("configFileValue"),
//   		EnableECSLogMetadata: jsii.Boolean(false),
//   	},
//   }
//
// Experimental.
type FirelensConfig struct {
	// The log router to use.
	// Experimental.
	Type FirelensLogRouterType `field:"required" json:"type" yaml:"type"`
	// Firelens options.
	// Experimental.
	Options *FirelensOptions `field:"optional" json:"options" yaml:"options"`
}

