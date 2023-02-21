package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableStreamSAMPTProperty := &tableStreamSAMPTProperty{
//   	streamName: jsii.String("streamName"),
//   	tableName: jsii.String("tableName"),
//   }
//
type CfnFunction_TableStreamSAMPTProperty struct {
	// `CfnFunction.TableStreamSAMPTProperty.StreamName`.
	StreamName *string `field:"required" json:"streamName" yaml:"streamName"`
	// `CfnFunction.TableStreamSAMPTProperty.TableName`.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

