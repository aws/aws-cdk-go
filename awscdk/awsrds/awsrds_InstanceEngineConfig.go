package awsrds


// The type returned from the {@link IInstanceEngine.bind} method.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var optionGroup optionGroup
//
//   instanceEngineConfig := &instanceEngineConfig{
//   	features: &instanceEngineFeatures{
//   		s3Export: jsii.String("s3Export"),
//   		s3Import: jsii.String("s3Import"),
//   	},
//   	optionGroup: optionGroup,
//   }
//
type InstanceEngineConfig struct {
	// Features supported by the database engine.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBEngineVersion.html
	//
	Features *InstanceEngineFeatures `field:"optional" json:"features" yaml:"features"`
	// Option group of the database.
	OptionGroup IOptionGroup `field:"optional" json:"optionGroup" yaml:"optionGroup"`
}

