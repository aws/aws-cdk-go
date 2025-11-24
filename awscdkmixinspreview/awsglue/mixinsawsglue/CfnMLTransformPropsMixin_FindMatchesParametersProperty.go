package mixinsawsglue


// The parameters to configure the find matches transform.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   findMatchesParametersProperty := &FindMatchesParametersProperty{
//   	AccuracyCostTradeoff: jsii.Number(123),
//   	EnforceProvidedLabels: jsii.Boolean(false),
//   	PrecisionRecallTradeoff: jsii.Number(123),
//   	PrimaryKeyColumnName: jsii.String("primaryKeyColumnName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-mltransform-findmatchesparameters.html
//
type CfnMLTransformPropsMixin_FindMatchesParametersProperty struct {
	// The value that is selected when tuning your transform for a balance between accuracy and cost.
	//
	// A value of 0.5 means that the system balances accuracy and cost concerns. A value of 1.0 means a bias purely for accuracy, which typically results in a higher cost, sometimes substantially higher. A value of 0.0 means a bias purely for cost, which results in a less accurate `FindMatches` transform, sometimes with unacceptable accuracy.
	//
	// Accuracy measures how well the transform finds true positives and true negatives. Increasing accuracy requires more machine resources and cost. But it also results in increased recall.
	//
	// Cost measures how many compute resources, and thus money, are consumed to run the transform.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-mltransform-findmatchesparameters.html#cfn-glue-mltransform-findmatchesparameters-accuracycosttradeoff
	//
	AccuracyCostTradeoff *float64 `field:"optional" json:"accuracyCostTradeoff" yaml:"accuracyCostTradeoff"`
	// The value to switch on or off to force the output to match the provided labels from users.
	//
	// If the value is `True` , the `find matches` transform forces the output to match the provided labels. The results override the normal conflation results. If the value is `False` , the `find matches` transform does not ensure all the labels provided are respected, and the results rely on the trained model.
	//
	// Note that setting this value to true may increase the conflation execution time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-mltransform-findmatchesparameters.html#cfn-glue-mltransform-findmatchesparameters-enforceprovidedlabels
	//
	EnforceProvidedLabels interface{} `field:"optional" json:"enforceProvidedLabels" yaml:"enforceProvidedLabels"`
	// The value selected when tuning your transform for a balance between precision and recall.
	//
	// A value of 0.5 means no preference; a value of 1.0 means a bias purely for precision, and a value of 0.0 means a bias for recall. Because this is a tradeoff, choosing values close to 1.0 means very low recall, and choosing values close to 0.0 results in very low precision.
	//
	// The precision metric indicates how often your model is correct when it predicts a match.
	//
	// The recall metric indicates that for an actual match, how often your model predicts the match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-mltransform-findmatchesparameters.html#cfn-glue-mltransform-findmatchesparameters-precisionrecalltradeoff
	//
	PrecisionRecallTradeoff *float64 `field:"optional" json:"precisionRecallTradeoff" yaml:"precisionRecallTradeoff"`
	// The name of a column that uniquely identifies rows in the source table.
	//
	// Used to help identify matching records.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-mltransform-findmatchesparameters.html#cfn-glue-mltransform-findmatchesparameters-primarykeycolumnname
	//
	PrimaryKeyColumnName *string `field:"optional" json:"primaryKeyColumnName" yaml:"primaryKeyColumnName"`
}

