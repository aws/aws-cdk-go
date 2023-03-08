// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   rawBucket := s3.NewCfnBucket(this, jsii.String("Bucket"), &CfnBucketProps{
//   })
//   // -or-
//   rawBucketAlt := myBucket.Node.defaultChild.(cfnBucket)
//
//   // then
//   rawBucket.CfnOptions.Condition = awscdk.NewCfnCondition(this, jsii.String("EnableBucket"), &CfnConditionProps{
//   })
//   rawBucket.CfnOptions.Metadata = map[string]interface{}{
//   	"metadataKey": jsii.String("MetadataValue"),
//   }
//
type CfnConditionProps struct {
	// The expression that the condition will evaluate.
	Expression ICfnConditionExpression `field:"optional" json:"expression" yaml:"expression"`
}

