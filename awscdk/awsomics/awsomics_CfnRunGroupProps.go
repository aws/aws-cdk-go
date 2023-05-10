package awsomics


// Properties for defining a `CfnRunGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRunGroupProps := &CfnRunGroupProps{
//   	MaxCpus: jsii.Number(123),
//   	MaxDuration: jsii.Number(123),
//   	MaxRuns: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnRunGroupProps struct {
	// The group's maximum CPU count setting.
	MaxCpus *float64 `field:"optional" json:"maxCpus" yaml:"maxCpus"`
	// The group's maximum duration setting in minutes.
	MaxDuration *float64 `field:"optional" json:"maxDuration" yaml:"maxDuration"`
	// The group's maximum concurrent run setting.
	MaxRuns *float64 `field:"optional" json:"maxRuns" yaml:"maxRuns"`
	// The group's name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Tags for the group.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

