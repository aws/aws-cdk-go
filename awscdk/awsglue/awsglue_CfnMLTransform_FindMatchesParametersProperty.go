package awsglue


// The parameters to configure the find matches transform.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   findMatchesParametersProperty := &findMatchesParametersProperty{
//   	primaryKeyColumnName: jsii.String("primaryKeyColumnName"),
//
//   	// the properties below are optional
//   	accuracyCostTradeoff: jsii.Number(123),
//   	enforceProvidedLabels: jsii.Boolean(false),
//   	precisionRecallTradeoff: jsii.Number(123),
//   }
//
type CfnMLTransform_FindMatchesParametersProperty struct {
	// The name of a column that uniquely identifies rows in the source table.
	//
	// Used to help identify matching records.
	PrimaryKeyColumnName *string `field:"required" json:"primaryKeyColumnName" yaml:"primaryKeyColumnName"`
	// The value that is selected when tuning your transform for a balance between accuracy and cost.
	//
	// A value of 0.5 means that the system balances accuracy and cost concerns. A value of 1.0 means a bias purely for accuracy, which typically results in a higher cost, sometimes substantially higher. A value of 0.0 means a bias purely for cost, which results in a less accurate `FindMatches` transform, sometimes with unacceptable accuracy.
	//
	// Accuracy measures how well the transform finds true positives and true negatives. Increasing accuracy requires more machine resources and cost. But it also results in increased recall.
	//
	// Cost measures how many compute resources, and thus money, are consumed to run the transform.
	AccuracyCostTradeoff *float64 `field:"optional" json:"accuracyCostTradeoff" yaml:"accuracyCostTradeoff"`
	// The value to switch on or off to force the output to match the provided labels from users.
	//
	// If the value is `True` , the `find matches` transform forces the output to match the provided labels. The results override the normal conflation results. If the value is `False` , the `find matches` transform does not ensure all the labels provided are respected, and the results rely on the trained model.
	//
	// Note that setting this value to true may increase the conflation execution time.
	EnforceProvidedLabels interface{} `field:"optional" json:"enforceProvidedLabels" yaml:"enforceProvidedLabels"`
	// The value selected when tuning your transform for a balance between precision and recall.
	//
	// A value of 0.5 means no preference; a value of 1.0 means a bias purely for precision, and a value of 0.0 means a bias for recall. Because this is a tradeoff, choosing values close to 1.0 means very low recall, and choosing values close to 0.0 results in very low precision.
	//
	// The precision metric indicates how often your model is correct when it predicts a match.
	//
	// The recall metric indicates that for an actual match, how often your model predicts the match.
	PrecisionRecallTradeoff *float64 `field:"optional" json:"precisionRecallTradeoff" yaml:"precisionRecallTradeoff"`
}

