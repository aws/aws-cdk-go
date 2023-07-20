package awsdatabrew


// Parameters that are used as inputs for various recipe actions.
//
// The parameters are specific to the context in which they're used.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var input interface{}
//
//   recipeParametersProperty := &RecipeParametersProperty{
//   	AggregateFunction: jsii.String("aggregateFunction"),
//   	Base: jsii.String("base"),
//   	CaseStatement: jsii.String("caseStatement"),
//   	CategoryMap: jsii.String("categoryMap"),
//   	CharsToRemove: jsii.String("charsToRemove"),
//   	CollapseConsecutiveWhitespace: jsii.String("collapseConsecutiveWhitespace"),
//   	ColumnDataType: jsii.String("columnDataType"),
//   	ColumnRange: jsii.String("columnRange"),
//   	Count: jsii.String("count"),
//   	CustomCharacters: jsii.String("customCharacters"),
//   	CustomStopWords: jsii.String("customStopWords"),
//   	CustomValue: jsii.String("customValue"),
//   	DatasetsColumns: jsii.String("datasetsColumns"),
//   	DateAddValue: jsii.String("dateAddValue"),
//   	DateTimeFormat: jsii.String("dateTimeFormat"),
//   	DateTimeParameters: jsii.String("dateTimeParameters"),
//   	DeleteOtherRows: jsii.String("deleteOtherRows"),
//   	Delimiter: jsii.String("delimiter"),
//   	EndPattern: jsii.String("endPattern"),
//   	EndPosition: jsii.String("endPosition"),
//   	EndValue: jsii.String("endValue"),
//   	ExpandContractions: jsii.String("expandContractions"),
//   	Exponent: jsii.String("exponent"),
//   	FalseString: jsii.String("falseString"),
//   	GroupByAggFunctionOptions: jsii.String("groupByAggFunctionOptions"),
//   	GroupByColumns: jsii.String("groupByColumns"),
//   	HiddenColumns: jsii.String("hiddenColumns"),
//   	IgnoreCase: jsii.String("ignoreCase"),
//   	IncludeInSplit: jsii.String("includeInSplit"),
//   	Input: input,
//   	Interval: jsii.String("interval"),
//   	IsText: jsii.String("isText"),
//   	JoinKeys: jsii.String("joinKeys"),
//   	JoinType: jsii.String("joinType"),
//   	LeftColumns: jsii.String("leftColumns"),
//   	Limit: jsii.String("limit"),
//   	LowerBound: jsii.String("lowerBound"),
//   	MapType: jsii.String("mapType"),
//   	ModeType: jsii.String("modeType"),
//   	MultiLine: jsii.Boolean(false),
//   	NumRows: jsii.String("numRows"),
//   	NumRowsAfter: jsii.String("numRowsAfter"),
//   	NumRowsBefore: jsii.String("numRowsBefore"),
//   	OrderByColumn: jsii.String("orderByColumn"),
//   	OrderByColumns: jsii.String("orderByColumns"),
//   	Other: jsii.String("other"),
//   	Pattern: jsii.String("pattern"),
//   	PatternOption1: jsii.String("patternOption1"),
//   	PatternOption2: jsii.String("patternOption2"),
//   	PatternOptions: jsii.String("patternOptions"),
//   	Period: jsii.String("period"),
//   	Position: jsii.String("position"),
//   	RemoveAllPunctuation: jsii.String("removeAllPunctuation"),
//   	RemoveAllQuotes: jsii.String("removeAllQuotes"),
//   	RemoveAllWhitespace: jsii.String("removeAllWhitespace"),
//   	RemoveCustomCharacters: jsii.String("removeCustomCharacters"),
//   	RemoveCustomValue: jsii.String("removeCustomValue"),
//   	RemoveLeadingAndTrailingPunctuation: jsii.String("removeLeadingAndTrailingPunctuation"),
//   	RemoveLeadingAndTrailingQuotes: jsii.String("removeLeadingAndTrailingQuotes"),
//   	RemoveLeadingAndTrailingWhitespace: jsii.String("removeLeadingAndTrailingWhitespace"),
//   	RemoveLetters: jsii.String("removeLetters"),
//   	RemoveNumbers: jsii.String("removeNumbers"),
//   	RemoveSourceColumn: jsii.String("removeSourceColumn"),
//   	RemoveSpecialCharacters: jsii.String("removeSpecialCharacters"),
//   	RightColumns: jsii.String("rightColumns"),
//   	SampleSize: jsii.String("sampleSize"),
//   	SampleType: jsii.String("sampleType"),
//   	SecondaryInputs: []interface{}{
//   		&SecondaryInputProperty{
//   			DataCatalogInputDefinition: &DataCatalogInputDefinitionProperty{
//   				CatalogId: jsii.String("catalogId"),
//   				DatabaseName: jsii.String("databaseName"),
//   				TableName: jsii.String("tableName"),
//   				TempDirectory: &S3LocationProperty{
//   					Bucket: jsii.String("bucket"),
//
//   					// the properties below are optional
//   					Key: jsii.String("key"),
//   				},
//   			},
//   			S3InputDefinition: &S3LocationProperty{
//   				Bucket: jsii.String("bucket"),
//
//   				// the properties below are optional
//   				Key: jsii.String("key"),
//   			},
//   		},
//   	},
//   	SecondInput: jsii.String("secondInput"),
//   	SheetIndexes: []interface{}{
//   		jsii.Number(123),
//   	},
//   	SheetNames: []*string{
//   		jsii.String("sheetNames"),
//   	},
//   	SourceColumn: jsii.String("sourceColumn"),
//   	SourceColumn1: jsii.String("sourceColumn1"),
//   	SourceColumn2: jsii.String("sourceColumn2"),
//   	SourceColumns: jsii.String("sourceColumns"),
//   	StartColumnIndex: jsii.String("startColumnIndex"),
//   	StartPattern: jsii.String("startPattern"),
//   	StartPosition: jsii.String("startPosition"),
//   	StartValue: jsii.String("startValue"),
//   	StemmingMode: jsii.String("stemmingMode"),
//   	StepCount: jsii.String("stepCount"),
//   	StepIndex: jsii.String("stepIndex"),
//   	StopWordsMode: jsii.String("stopWordsMode"),
//   	Strategy: jsii.String("strategy"),
//   	TargetColumn: jsii.String("targetColumn"),
//   	TargetColumnNames: jsii.String("targetColumnNames"),
//   	TargetDateFormat: jsii.String("targetDateFormat"),
//   	TargetIndex: jsii.String("targetIndex"),
//   	TimeZone: jsii.String("timeZone"),
//   	TokenizerPattern: jsii.String("tokenizerPattern"),
//   	TrueString: jsii.String("trueString"),
//   	UdfLang: jsii.String("udfLang"),
//   	Units: jsii.String("units"),
//   	UnpivotColumn: jsii.String("unpivotColumn"),
//   	UpperBound: jsii.String("upperBound"),
//   	UseNewDataFrame: jsii.String("useNewDataFrame"),
//   	Value: jsii.String("value"),
//   	Value1: jsii.String("value1"),
//   	Value2: jsii.String("value2"),
//   	ValueColumn: jsii.String("valueColumn"),
//   	ViewFrame: jsii.String("viewFrame"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html
//
type CfnRecipe_RecipeParametersProperty struct {
	// The name of an aggregation function to apply.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-aggregatefunction
	//
	AggregateFunction *string `field:"optional" json:"aggregateFunction" yaml:"aggregateFunction"`
	// The number of digits used in a counting system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-base
	//
	Base *string `field:"optional" json:"base" yaml:"base"`
	// A case statement associated with a recipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-casestatement
	//
	CaseStatement *string `field:"optional" json:"caseStatement" yaml:"caseStatement"`
	// A category map used for one-hot encoding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-categorymap
	//
	CategoryMap *string `field:"optional" json:"categoryMap" yaml:"categoryMap"`
	// Characters to remove from a step that applies one-hot encoding or tokenization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-charstoremove
	//
	CharsToRemove *string `field:"optional" json:"charsToRemove" yaml:"charsToRemove"`
	// Remove any non-word non-punctuation character.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-collapseconsecutivewhitespace
	//
	CollapseConsecutiveWhitespace *string `field:"optional" json:"collapseConsecutiveWhitespace" yaml:"collapseConsecutiveWhitespace"`
	// The data type of the column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-columndatatype
	//
	ColumnDataType *string `field:"optional" json:"columnDataType" yaml:"columnDataType"`
	// A range of columns to which a step is applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-columnrange
	//
	ColumnRange *string `field:"optional" json:"columnRange" yaml:"columnRange"`
	// The number of times a string needs to be repeated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-count
	//
	Count *string `field:"optional" json:"count" yaml:"count"`
	// One or more characters that can be substituted or removed, depending on the context.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-customcharacters
	//
	CustomCharacters *string `field:"optional" json:"customCharacters" yaml:"customCharacters"`
	// A list of words to ignore in a step that applies word tokenization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-customstopwords
	//
	CustomStopWords *string `field:"optional" json:"customStopWords" yaml:"customStopWords"`
	// A list of custom values to use in a step that requires that you provide a value to finish the operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-customvalue
	//
	CustomValue *string `field:"optional" json:"customValue" yaml:"customValue"`
	// A list of the dataset columns included in a project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-datasetscolumns
	//
	DatasetsColumns *string `field:"optional" json:"datasetsColumns" yaml:"datasetsColumns"`
	// A value that specifies how many units of time to add or subtract for a date math operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-dateaddvalue
	//
	DateAddValue *string `field:"optional" json:"dateAddValue" yaml:"dateAddValue"`
	// A date format to apply to a date.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-datetimeformat
	//
	DateTimeFormat *string `field:"optional" json:"dateTimeFormat" yaml:"dateTimeFormat"`
	// A set of parameters associated with a datetime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-datetimeparameters
	//
	DateTimeParameters *string `field:"optional" json:"dateTimeParameters" yaml:"dateTimeParameters"`
	// Determines whether unmapped rows in a categorical mapping should be deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-deleteotherrows
	//
	DeleteOtherRows *string `field:"optional" json:"deleteOtherRows" yaml:"deleteOtherRows"`
	// The delimiter to use when parsing separated values in a text file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-delimiter
	//
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// The end pattern to locate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-endpattern
	//
	EndPattern *string `field:"optional" json:"endPattern" yaml:"endPattern"`
	// The end position to locate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-endposition
	//
	EndPosition *string `field:"optional" json:"endPosition" yaml:"endPosition"`
	// The end value to locate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-endvalue
	//
	EndValue *string `field:"optional" json:"endValue" yaml:"endValue"`
	// A list of word contractions and what they expand to.
	//
	// For eample: *can't* ; *cannot* ; *can not* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-expandcontractions
	//
	ExpandContractions *string `field:"optional" json:"expandContractions" yaml:"expandContractions"`
	// The exponent to apply in an exponential operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-exponent
	//
	Exponent *string `field:"optional" json:"exponent" yaml:"exponent"`
	// A value that represents `FALSE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-falsestring
	//
	FalseString *string `field:"optional" json:"falseString" yaml:"falseString"`
	// Specifies options to apply to the `GROUP BY` used in an aggregation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-groupbyaggfunctionoptions
	//
	GroupByAggFunctionOptions *string `field:"optional" json:"groupByAggFunctionOptions" yaml:"groupByAggFunctionOptions"`
	// The columns to use in the `GROUP BY` clause.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-groupbycolumns
	//
	GroupByColumns *string `field:"optional" json:"groupByColumns" yaml:"groupByColumns"`
	// A list of columns to hide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-hiddencolumns
	//
	HiddenColumns *string `field:"optional" json:"hiddenColumns" yaml:"hiddenColumns"`
	// Indicates that lower and upper case letters are treated equally.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-ignorecase
	//
	IgnoreCase *string `field:"optional" json:"ignoreCase" yaml:"ignoreCase"`
	// Indicates if this column is participating in a split transform.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-includeinsplit
	//
	IncludeInSplit *string `field:"optional" json:"includeInSplit" yaml:"includeInSplit"`
	// The input location to load the dataset from - Amazon S3 or AWS Glue Data Catalog .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-input
	//
	Input interface{} `field:"optional" json:"input" yaml:"input"`
	// The number of characters to split by.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-interval
	//
	Interval *string `field:"optional" json:"interval" yaml:"interval"`
	// Indicates if the content is text.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-istext
	//
	IsText *string `field:"optional" json:"isText" yaml:"isText"`
	// The keys or columns involved in a join.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-joinkeys
	//
	JoinKeys *string `field:"optional" json:"joinKeys" yaml:"joinKeys"`
	// The type of join to use, for example, `INNER JOIN` , `OUTER JOIN` , and so on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-jointype
	//
	JoinType *string `field:"optional" json:"joinType" yaml:"joinType"`
	// The columns on the left side of the join.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-leftcolumns
	//
	LeftColumns *string `field:"optional" json:"leftColumns" yaml:"leftColumns"`
	// The number of times to perform `split` or `replaceBy` in a string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-limit
	//
	Limit *string `field:"optional" json:"limit" yaml:"limit"`
	// The lower boundary for a value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-lowerbound
	//
	LowerBound *string `field:"optional" json:"lowerBound" yaml:"lowerBound"`
	// The type of mappings to apply to construct a new dynamic frame.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-maptype
	//
	MapType *string `field:"optional" json:"mapType" yaml:"mapType"`
	// Determines the manner in which mode value is calculated, in case there is more than one mode value.
	//
	// Valid values: `NONE` | `AVERAGE` | `MINIMUM` | `MAXIMUM`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-modetype
	//
	ModeType *string `field:"optional" json:"modeType" yaml:"modeType"`
	// Specifies whether JSON input contains embedded new line characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-multiline
	//
	MultiLine interface{} `field:"optional" json:"multiLine" yaml:"multiLine"`
	// The number of rows to consider in a window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-numrows
	//
	NumRows *string `field:"optional" json:"numRows" yaml:"numRows"`
	// The number of rows to consider after the current row in a window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-numrowsafter
	//
	NumRowsAfter *string `field:"optional" json:"numRowsAfter" yaml:"numRowsAfter"`
	// The number of rows to consider before the current row in a window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-numrowsbefore
	//
	NumRowsBefore *string `field:"optional" json:"numRowsBefore" yaml:"numRowsBefore"`
	// A column to sort the results by.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-orderbycolumn
	//
	OrderByColumn *string `field:"optional" json:"orderByColumn" yaml:"orderByColumn"`
	// The columns to sort the results by.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-orderbycolumns
	//
	OrderByColumns *string `field:"optional" json:"orderByColumns" yaml:"orderByColumns"`
	// The value to assign to unmapped cells, in categorical mapping.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-other
	//
	Other *string `field:"optional" json:"other" yaml:"other"`
	// The pattern to locate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-pattern
	//
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// The starting pattern to split between.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-patternoption1
	//
	PatternOption1 *string `field:"optional" json:"patternOption1" yaml:"patternOption1"`
	// The ending pattern to split between.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-patternoption2
	//
	PatternOption2 *string `field:"optional" json:"patternOption2" yaml:"patternOption2"`
	// For splitting by multiple delimiters: A JSON-encoded string that lists the patterns in the format.
	//
	// For example: `[{\"pattern\":\"1\",\"includeInSplit\":true}]`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-patternoptions
	//
	PatternOptions *string `field:"optional" json:"patternOptions" yaml:"patternOptions"`
	// The size of the rolling window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-period
	//
	Period *string `field:"optional" json:"period" yaml:"period"`
	// The character index within a string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-position
	//
	Position *string `field:"optional" json:"position" yaml:"position"`
	// If `true` , removes all of the following characters: `.` `.!` `.,` `.?`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-removeallpunctuation
	//
	RemoveAllPunctuation *string `field:"optional" json:"removeAllPunctuation" yaml:"removeAllPunctuation"`
	// If `true` , removes all single quotes and double quotes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-removeallquotes
	//
	RemoveAllQuotes *string `field:"optional" json:"removeAllQuotes" yaml:"removeAllQuotes"`
	// If `true` , removes all whitespaces from the value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-removeallwhitespace
	//
	RemoveAllWhitespace *string `field:"optional" json:"removeAllWhitespace" yaml:"removeAllWhitespace"`
	// If `true` , removes all chraracters specified by `CustomCharacters` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-removecustomcharacters
	//
	RemoveCustomCharacters *string `field:"optional" json:"removeCustomCharacters" yaml:"removeCustomCharacters"`
	// If `true` , removes all chraracters specified by `CustomValue` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-removecustomvalue
	//
	RemoveCustomValue *string `field:"optional" json:"removeCustomValue" yaml:"removeCustomValue"`
	// If `true` , removes the following characters if they occur at the start or end of the value: `.` `!` `,` `?`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-removeleadingandtrailingpunctuation
	//
	RemoveLeadingAndTrailingPunctuation *string `field:"optional" json:"removeLeadingAndTrailingPunctuation" yaml:"removeLeadingAndTrailingPunctuation"`
	// If `true` , removes single quotes and double quotes from the beginning and end of the value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-removeleadingandtrailingquotes
	//
	RemoveLeadingAndTrailingQuotes *string `field:"optional" json:"removeLeadingAndTrailingQuotes" yaml:"removeLeadingAndTrailingQuotes"`
	// If `true` , removes all whitespaces from the beginning and end of the value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-removeleadingandtrailingwhitespace
	//
	RemoveLeadingAndTrailingWhitespace *string `field:"optional" json:"removeLeadingAndTrailingWhitespace" yaml:"removeLeadingAndTrailingWhitespace"`
	// If `true` , removes all uppercase and lowercase alphabetic characters (A through Z;
	//
	// a through z).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-removeletters
	//
	RemoveLetters *string `field:"optional" json:"removeLetters" yaml:"removeLetters"`
	// If `true` , removes all numeric characters (0 through 9).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-removenumbers
	//
	RemoveNumbers *string `field:"optional" json:"removeNumbers" yaml:"removeNumbers"`
	// If `true` , the source column will be removed after un-nesting that column.
	//
	// (Used with nested column types, such as Map, Struct, or Array.)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-removesourcecolumn
	//
	RemoveSourceColumn *string `field:"optional" json:"removeSourceColumn" yaml:"removeSourceColumn"`
	// If `true` , removes all of the following characters: `!
	//
	// " # $ % & ' ( ) * + , - . / : ; < = > ? @ [ \ ] ^ _ ` { | } ~`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-removespecialcharacters
	//
	RemoveSpecialCharacters *string `field:"optional" json:"removeSpecialCharacters" yaml:"removeSpecialCharacters"`
	// The columns on the right side of a join.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-rightcolumns
	//
	RightColumns *string `field:"optional" json:"rightColumns" yaml:"rightColumns"`
	// The number of rows in the sample.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-samplesize
	//
	SampleSize *string `field:"optional" json:"sampleSize" yaml:"sampleSize"`
	// The sampling type to apply to the dataset.
	//
	// Valid values: `FIRST_N` | `LAST_N` | `RANDOM`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-sampletype
	//
	SampleType *string `field:"optional" json:"sampleType" yaml:"sampleType"`
	// A list of secondary inputs in a UNION transform.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-secondaryinputs
	//
	SecondaryInputs interface{} `field:"optional" json:"secondaryInputs" yaml:"secondaryInputs"`
	// A object value to indicate the second dataset used in a join.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-secondinput
	//
	SecondInput *string `field:"optional" json:"secondInput" yaml:"secondInput"`
	// One or more sheet numbers in the Excel file, which will be included in a dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-sheetindexes
	//
	SheetIndexes interface{} `field:"optional" json:"sheetIndexes" yaml:"sheetIndexes"`
	// Oone or more named sheets in the Excel file, which will be included in a dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-sheetnames
	//
	SheetNames *[]*string `field:"optional" json:"sheetNames" yaml:"sheetNames"`
	// A source column needed for an operation, step, or transform.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-sourcecolumn
	//
	SourceColumn *string `field:"optional" json:"sourceColumn" yaml:"sourceColumn"`
	// A source column needed for an operation, step, or transform.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-sourcecolumn1
	//
	SourceColumn1 *string `field:"optional" json:"sourceColumn1" yaml:"sourceColumn1"`
	// A source column needed for an operation, step, or transform.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-sourcecolumn2
	//
	SourceColumn2 *string `field:"optional" json:"sourceColumn2" yaml:"sourceColumn2"`
	// A list of source columns needed for an operation, step, or transform.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-sourcecolumns
	//
	SourceColumns *string `field:"optional" json:"sourceColumns" yaml:"sourceColumns"`
	// The index number of the first column used by an operation, step, or transform.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-startcolumnindex
	//
	StartColumnIndex *string `field:"optional" json:"startColumnIndex" yaml:"startColumnIndex"`
	// The starting pattern to locate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-startpattern
	//
	StartPattern *string `field:"optional" json:"startPattern" yaml:"startPattern"`
	// The starting position to locate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-startposition
	//
	StartPosition *string `field:"optional" json:"startPosition" yaml:"startPosition"`
	// The starting value to locate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-startvalue
	//
	StartValue *string `field:"optional" json:"startValue" yaml:"startValue"`
	// Indicates this operation uses stems and lemmas (base words) for word tokenization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-stemmingmode
	//
	StemmingMode *string `field:"optional" json:"stemmingMode" yaml:"stemmingMode"`
	// The total number of transforms in this recipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-stepcount
	//
	StepCount *string `field:"optional" json:"stepCount" yaml:"stepCount"`
	// The index ID of a step.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-stepindex
	//
	StepIndex *string `field:"optional" json:"stepIndex" yaml:"stepIndex"`
	// Indicates this operation uses stop words as part of word tokenization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-stopwordsmode
	//
	StopWordsMode *string `field:"optional" json:"stopWordsMode" yaml:"stopWordsMode"`
	// The resolution strategy to apply in resolving ambiguities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-strategy
	//
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
	// The column targeted by this operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-targetcolumn
	//
	TargetColumn *string `field:"optional" json:"targetColumn" yaml:"targetColumn"`
	// The names to give columns altered by this operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-targetcolumnnames
	//
	TargetColumnNames *string `field:"optional" json:"targetColumnNames" yaml:"targetColumnNames"`
	// The date format to convert to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-targetdateformat
	//
	TargetDateFormat *string `field:"optional" json:"targetDateFormat" yaml:"targetDateFormat"`
	// The index number of an object that is targeted by this operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-targetindex
	//
	TargetIndex *string `field:"optional" json:"targetIndex" yaml:"targetIndex"`
	// The current timezone that you want to use for dates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-timezone
	//
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
	// A regex expression to use when splitting text into terms, also called words or tokens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-tokenizerpattern
	//
	TokenizerPattern *string `field:"optional" json:"tokenizerPattern" yaml:"tokenizerPattern"`
	// A value to use to represent `TRUE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-truestring
	//
	TrueString *string `field:"optional" json:"trueString" yaml:"trueString"`
	// The language that's used in the user-defined function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-udflang
	//
	UdfLang *string `field:"optional" json:"udfLang" yaml:"udfLang"`
	// Specifies a unit of time.
	//
	// For example: `MINUTES` ; `SECONDS` ; `HOURS` ; etc.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-units
	//
	Units *string `field:"optional" json:"units" yaml:"units"`
	// Cast columns as rows, so that each value is a different row in a single column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-unpivotcolumn
	//
	UnpivotColumn *string `field:"optional" json:"unpivotColumn" yaml:"unpivotColumn"`
	// The upper boundary for a value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-upperbound
	//
	UpperBound *string `field:"optional" json:"upperBound" yaml:"upperBound"`
	// Create a new container to hold a dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-usenewdataframe
	//
	UseNewDataFrame *string `field:"optional" json:"useNewDataFrame" yaml:"useNewDataFrame"`
	// A static value that can be used in a comparison, a substitution, or in another context-specific way.
	//
	// A `Value` can be a number, string, or other datatype, depending on the recipe action in which it's used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
	// A value that's used by this operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-value1
	//
	Value1 *string `field:"optional" json:"value1" yaml:"value1"`
	// A value that's used by this operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-value2
	//
	Value2 *string `field:"optional" json:"value2" yaml:"value2"`
	// The column that is provided as a value that's used by this operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-valuecolumn
	//
	ValueColumn *string `field:"optional" json:"valueColumn" yaml:"valueColumn"`
	// The subset of rows currently available for viewing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipeparameters.html#cfn-databrew-recipe-recipeparameters-viewframe
	//
	ViewFrame *string `field:"optional" json:"viewFrame" yaml:"viewFrame"`
}

