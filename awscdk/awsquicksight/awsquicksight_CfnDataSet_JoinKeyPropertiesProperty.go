package awsquicksight


// Properties associated with the columns participating in a join.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   joinKeyPropertiesProperty := &joinKeyPropertiesProperty{
//   	uniqueKey: jsii.Boolean(false),
//   }
//
type CfnDataSet_JoinKeyPropertiesProperty struct {
	// A value that indicates that a row in a table is uniquely identified by the columns in a join key.
	//
	// This is used by Amazon QuickSight to optimize query performance.
	UniqueKey interface{} `field:"optional" json:"uniqueKey" yaml:"uniqueKey"`
}

