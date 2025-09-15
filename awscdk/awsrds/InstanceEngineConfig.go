package awsrds


// The type returned from the `IInstanceEngine.bind` method.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var optionGroup optionGroup
//
//   instanceEngineConfig := &InstanceEngineConfig{
//   	Features: &InstanceEngineFeatures{
//   		S3Export: jsii.String("s3Export"),
//   		S3Import: jsii.String("s3Import"),
//   	},
//   	OptionGroup: optionGroup,
//   }
//
type InstanceEngineConfig struct {
	// Features supported by the database engine.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBEngineVersion.html
	//
	// Default: - no features.
	//
	Features *InstanceEngineFeatures `field:"optional" json:"features" yaml:"features"`
	// Option group of the database.
	// Default: - none.
	//
	OptionGroup IOptionGroup `field:"optional" json:"optionGroup" yaml:"optionGroup"`
}

