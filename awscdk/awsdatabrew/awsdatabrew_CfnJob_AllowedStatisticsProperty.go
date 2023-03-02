package awsdatabrew


// Configuration of statistics that are allowed to be run on columns that contain detected entities.
//
// When undefined, no statistics will be computed on columns that contain detected entities.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   allowedStatisticsProperty := &allowedStatisticsProperty{
//   	statistics: []*string{
//   		jsii.String("statistics"),
//   	},
//   }
//
type CfnJob_AllowedStatisticsProperty struct {
	// One or more column statistics to allow for columns that contain detected entities.
	Statistics *[]*string `field:"required" json:"statistics" yaml:"statistics"`
}

