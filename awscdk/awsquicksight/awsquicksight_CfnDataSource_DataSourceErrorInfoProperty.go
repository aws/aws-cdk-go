package awsquicksight


// Error information for the data source creation or update.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceErrorInfoProperty := &dataSourceErrorInfoProperty{
//   	message: jsii.String("message"),
//   	type: jsii.String("type"),
//   }
//
type CfnDataSource_DataSourceErrorInfoProperty struct {
	// Error message.
	Message *string `field:"optional" json:"message" yaml:"message"`
	// Error type.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

