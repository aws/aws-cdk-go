// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   rawBucket := s3.NewCfnBucket(this, jsii.String("Bucket"), &cfnBucketProps{
//   })
//   // -or-
//   rawBucketAlt := myBucket.node.defaultChild.(cfnBucket)
//
//   // then
//   rawBucket.cfnOptions.condition = awscdk.NewCfnCondition(this, jsii.String("EnableBucket"), &cfnConditionProps{
//   })
//   rawBucket.cfnOptions.metadata = map[string]interface{}{
//   	"metadataKey": jsii.String("MetadataValue"),
//   }
//
type CfnConditionProps struct {
	// The expression that the condition will evaluate.
	Expression ICfnConditionExpression `field:"optional" json:"expression" yaml:"expression"`
}

