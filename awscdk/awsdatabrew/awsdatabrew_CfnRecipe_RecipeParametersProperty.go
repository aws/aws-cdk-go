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
//   recipeParametersProperty := &recipeParametersProperty{
//   	aggregateFunction: jsii.String("aggregateFunction"),
//   	base: jsii.String("base"),
//   	caseStatement: jsii.String("caseStatement"),
//   	categoryMap: jsii.String("categoryMap"),
//   	charsToRemove: jsii.String("charsToRemove"),
//   	collapseConsecutiveWhitespace: jsii.String("collapseConsecutiveWhitespace"),
//   	columnDataType: jsii.String("columnDataType"),
//   	columnRange: jsii.String("columnRange"),
//   	count: jsii.String("count"),
//   	customCharacters: jsii.String("customCharacters"),
//   	customStopWords: jsii.String("customStopWords"),
//   	customValue: jsii.String("customValue"),
//   	datasetsColumns: jsii.String("datasetsColumns"),
//   	dateAddValue: jsii.String("dateAddValue"),
//   	dateTimeFormat: jsii.String("dateTimeFormat"),
//   	dateTimeParameters: jsii.String("dateTimeParameters"),
//   	deleteOtherRows: jsii.String("deleteOtherRows"),
//   	delimiter: jsii.String("delimiter"),
//   	endPattern: jsii.String("endPattern"),
//   	endPosition: jsii.String("endPosition"),
//   	endValue: jsii.String("endValue"),
//   	expandContractions: jsii.String("expandContractions"),
//   	exponent: jsii.String("exponent"),
//   	falseString: jsii.String("falseString"),
//   	groupByAggFunctionOptions: jsii.String("groupByAggFunctionOptions"),
//   	groupByColumns: jsii.String("groupByColumns"),
//   	hiddenColumns: jsii.String("hiddenColumns"),
//   	ignoreCase: jsii.String("ignoreCase"),
//   	includeInSplit: jsii.String("includeInSplit"),
//   	input: input,
//   	interval: jsii.String("interval"),
//   	isText: jsii.String("isText"),
//   	joinKeys: jsii.String("joinKeys"),
//   	joinType: jsii.String("joinType"),
//   	leftColumns: jsii.String("leftColumns"),
//   	limit: jsii.String("limit"),
//   	lowerBound: jsii.String("lowerBound"),
//   	mapType: jsii.String("mapType"),
//   	modeType: jsii.String("modeType"),
//   	multiLine: jsii.Boolean(false),
//   	numRows: jsii.String("numRows"),
//   	numRowsAfter: jsii.String("numRowsAfter"),
//   	numRowsBefore: jsii.String("numRowsBefore"),
//   	orderByColumn: jsii.String("orderByColumn"),
//   	orderByColumns: jsii.String("orderByColumns"),
//   	other: jsii.String("other"),
//   	pattern: jsii.String("pattern"),
//   	patternOption1: jsii.String("patternOption1"),
//   	patternOption2: jsii.String("patternOption2"),
//   	patternOptions: jsii.String("patternOptions"),
//   	period: jsii.String("period"),
//   	position: jsii.String("position"),
//   	removeAllPunctuation: jsii.String("removeAllPunctuation"),
//   	removeAllQuotes: jsii.String("removeAllQuotes"),
//   	removeAllWhitespace: jsii.String("removeAllWhitespace"),
//   	removeCustomCharacters: jsii.String("removeCustomCharacters"),
//   	removeCustomValue: jsii.String("removeCustomValue"),
//   	removeLeadingAndTrailingPunctuation: jsii.String("removeLeadingAndTrailingPunctuation"),
//   	removeLeadingAndTrailingQuotes: jsii.String("removeLeadingAndTrailingQuotes"),
//   	removeLeadingAndTrailingWhitespace: jsii.String("removeLeadingAndTrailingWhitespace"),
//   	removeLetters: jsii.String("removeLetters"),
//   	removeNumbers: jsii.String("removeNumbers"),
//   	removeSourceColumn: jsii.String("removeSourceColumn"),
//   	removeSpecialCharacters: jsii.String("removeSpecialCharacters"),
//   	rightColumns: jsii.String("rightColumns"),
//   	sampleSize: jsii.String("sampleSize"),
//   	sampleType: jsii.String("sampleType"),
//   	secondaryInputs: []interface{}{
//   		&secondaryInputProperty{
//   			dataCatalogInputDefinition: &dataCatalogInputDefinitionProperty{
//   				catalogId: jsii.String("catalogId"),
//   				databaseName: jsii.String("databaseName"),
//   				tableName: jsii.String("tableName"),
//   				tempDirectory: &s3LocationProperty{
//   					bucket: jsii.String("bucket"),
//
//   					// the properties below are optional
//   					key: jsii.String("key"),
//   				},
//   			},
//   			s3InputDefinition: &s3LocationProperty{
//   				bucket: jsii.String("bucket"),
//
//   				// the properties below are optional
//   				key: jsii.String("key"),
//   			},
//   		},
//   	},
//   	secondInput: jsii.String("secondInput"),
//   	sheetIndexes: []interface{}{
//   		jsii.Number(123),
//   	},
//   	sheetNames: []*string{
//   		jsii.String("sheetNames"),
//   	},
//   	sourceColumn: jsii.String("sourceColumn"),
//   	sourceColumn1: jsii.String("sourceColumn1"),
//   	sourceColumn2: jsii.String("sourceColumn2"),
//   	sourceColumns: jsii.String("sourceColumns"),
//   	startColumnIndex: jsii.String("startColumnIndex"),
//   	startPattern: jsii.String("startPattern"),
//   	startPosition: jsii.String("startPosition"),
//   	startValue: jsii.String("startValue"),
//   	stemmingMode: jsii.String("stemmingMode"),
//   	stepCount: jsii.String("stepCount"),
//   	stepIndex: jsii.String("stepIndex"),
//   	stopWordsMode: jsii.String("stopWordsMode"),
//   	strategy: jsii.String("strategy"),
//   	targetColumn: jsii.String("targetColumn"),
//   	targetColumnNames: jsii.String("targetColumnNames"),
//   	targetDateFormat: jsii.String("targetDateFormat"),
//   	targetIndex: jsii.String("targetIndex"),
//   	timeZone: jsii.String("timeZone"),
//   	tokenizerPattern: jsii.String("tokenizerPattern"),
//   	trueString: jsii.String("trueString"),
//   	udfLang: jsii.String("udfLang"),
//   	units: jsii.String("units"),
//   	unpivotColumn: jsii.String("unpivotColumn"),
//   	upperBound: jsii.String("upperBound"),
//   	useNewDataFrame: jsii.String("useNewDataFrame"),
//   	value: jsii.String("value"),
//   	value1: jsii.String("value1"),
//   	value2: jsii.String("value2"),
//   	valueColumn: jsii.String("valueColumn"),
//   	viewFrame: jsii.String("viewFrame"),
//   }
//
type CfnRecipe_RecipeParametersProperty struct {
	// The name of an aggregation function to apply.
	AggregateFunction *string `field:"optional" json:"aggregateFunction" yaml:"aggregateFunction"`
	// The number of digits used in a counting system.
	Base *string `field:"optional" json:"base" yaml:"base"`
	// A case statement associated with a recipe.
	CaseStatement *string `field:"optional" json:"caseStatement" yaml:"caseStatement"`
	// A category map used for one-hot encoding.
	CategoryMap *string `field:"optional" json:"categoryMap" yaml:"categoryMap"`
	// Characters to remove from a step that applies one-hot encoding or tokenization.
	CharsToRemove *string `field:"optional" json:"charsToRemove" yaml:"charsToRemove"`
	// Remove any non-word non-punctuation character.
	CollapseConsecutiveWhitespace *string `field:"optional" json:"collapseConsecutiveWhitespace" yaml:"collapseConsecutiveWhitespace"`
	// The data type of the column.
	ColumnDataType *string `field:"optional" json:"columnDataType" yaml:"columnDataType"`
	// A range of columns to which a step is applied.
	ColumnRange *string `field:"optional" json:"columnRange" yaml:"columnRange"`
	// The number of times a string needs to be repeated.
	Count *string `field:"optional" json:"count" yaml:"count"`
	// One or more characters that can be substituted or removed, depending on the context.
	CustomCharacters *string `field:"optional" json:"customCharacters" yaml:"customCharacters"`
	// A list of words to ignore in a step that applies word tokenization.
	CustomStopWords *string `field:"optional" json:"customStopWords" yaml:"customStopWords"`
	// A list of custom values to use in a step that requires that you provide a value to finish the operation.
	CustomValue *string `field:"optional" json:"customValue" yaml:"customValue"`
	// A list of the dataset columns included in a project.
	DatasetsColumns *string `field:"optional" json:"datasetsColumns" yaml:"datasetsColumns"`
	// A value that specifies how many units of time to add or subtract for a date math operation.
	DateAddValue *string `field:"optional" json:"dateAddValue" yaml:"dateAddValue"`
	// A date format to apply to a date.
	DateTimeFormat *string `field:"optional" json:"dateTimeFormat" yaml:"dateTimeFormat"`
	// A set of parameters associated with a datetime.
	DateTimeParameters *string `field:"optional" json:"dateTimeParameters" yaml:"dateTimeParameters"`
	// Determines whether unmapped rows in a categorical mapping should be deleted.
	DeleteOtherRows *string `field:"optional" json:"deleteOtherRows" yaml:"deleteOtherRows"`
	// The delimiter to use when parsing separated values in a text file.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// The end pattern to locate.
	EndPattern *string `field:"optional" json:"endPattern" yaml:"endPattern"`
	// The end position to locate.
	EndPosition *string `field:"optional" json:"endPosition" yaml:"endPosition"`
	// The end value to locate.
	EndValue *string `field:"optional" json:"endValue" yaml:"endValue"`
	// A list of word contractions and what they expand to.
	//
	// For eample: *can't* ; *cannot* ; *can not* .
	ExpandContractions *string `field:"optional" json:"expandContractions" yaml:"expandContractions"`
	// The exponent to apply in an exponential operation.
	Exponent *string `field:"optional" json:"exponent" yaml:"exponent"`
	// A value that represents `FALSE` .
	FalseString *string `field:"optional" json:"falseString" yaml:"falseString"`
	// Specifies options to apply to the `GROUP BY` used in an aggregation.
	GroupByAggFunctionOptions *string `field:"optional" json:"groupByAggFunctionOptions" yaml:"groupByAggFunctionOptions"`
	// The columns to use in the `GROUP BY` clause.
	GroupByColumns *string `field:"optional" json:"groupByColumns" yaml:"groupByColumns"`
	// A list of columns to hide.
	HiddenColumns *string `field:"optional" json:"hiddenColumns" yaml:"hiddenColumns"`
	// Indicates that lower and upper case letters are treated equally.
	IgnoreCase *string `field:"optional" json:"ignoreCase" yaml:"ignoreCase"`
	// Indicates if this column is participating in a split transform.
	IncludeInSplit *string `field:"optional" json:"includeInSplit" yaml:"includeInSplit"`
	// The input location to load the dataset from - Amazon S3 or AWS Glue Data Catalog .
	Input interface{} `field:"optional" json:"input" yaml:"input"`
	// The number of characters to split by.
	Interval *string `field:"optional" json:"interval" yaml:"interval"`
	// Indicates if the content is text.
	IsText *string `field:"optional" json:"isText" yaml:"isText"`
	// The keys or columns involved in a join.
	JoinKeys *string `field:"optional" json:"joinKeys" yaml:"joinKeys"`
	// The type of join to use, for example, `INNER JOIN` , `OUTER JOIN` , and so on.
	JoinType *string `field:"optional" json:"joinType" yaml:"joinType"`
	// The columns on the left side of the join.
	LeftColumns *string `field:"optional" json:"leftColumns" yaml:"leftColumns"`
	// The number of times to perform `split` or `replaceBy` in a string.
	Limit *string `field:"optional" json:"limit" yaml:"limit"`
	// The lower boundary for a value.
	LowerBound *string `field:"optional" json:"lowerBound" yaml:"lowerBound"`
	// The type of mappings to apply to construct a new dynamic frame.
	MapType *string `field:"optional" json:"mapType" yaml:"mapType"`
	// Determines the manner in which mode value is calculated, in case there is more than one mode value.
	//
	// Valid values: `NONE` | `AVERAGE` | `MINIMUM` | `MAXIMUM`.
	ModeType *string `field:"optional" json:"modeType" yaml:"modeType"`
	// Specifies whether JSON input contains embedded new line characters.
	MultiLine interface{} `field:"optional" json:"multiLine" yaml:"multiLine"`
	// The number of rows to consider in a window.
	NumRows *string `field:"optional" json:"numRows" yaml:"numRows"`
	// The number of rows to consider after the current row in a window.
	NumRowsAfter *string `field:"optional" json:"numRowsAfter" yaml:"numRowsAfter"`
	// The number of rows to consider before the current row in a window.
	NumRowsBefore *string `field:"optional" json:"numRowsBefore" yaml:"numRowsBefore"`
	// A column to sort the results by.
	OrderByColumn *string `field:"optional" json:"orderByColumn" yaml:"orderByColumn"`
	// The columns to sort the results by.
	OrderByColumns *string `field:"optional" json:"orderByColumns" yaml:"orderByColumns"`
	// The value to assign to unmapped cells, in categorical mapping.
	Other *string `field:"optional" json:"other" yaml:"other"`
	// The pattern to locate.
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// The starting pattern to split between.
	PatternOption1 *string `field:"optional" json:"patternOption1" yaml:"patternOption1"`
	// The ending pattern to split between.
	PatternOption2 *string `field:"optional" json:"patternOption2" yaml:"patternOption2"`
	// For splitting by multiple delimiters: A JSON-encoded string that lists the patterns in the format.
	//
	// For example: `[{\"pattern\":\"1\",\"includeInSplit\":true}]`.
	PatternOptions *string `field:"optional" json:"patternOptions" yaml:"patternOptions"`
	// The size of the rolling window.
	Period *string `field:"optional" json:"period" yaml:"period"`
	// The character index within a string.
	Position *string `field:"optional" json:"position" yaml:"position"`
	// If `true` , removes all of the following characters: `.` `.!` `.,` `.?`.
	RemoveAllPunctuation *string `field:"optional" json:"removeAllPunctuation" yaml:"removeAllPunctuation"`
	// If `true` , removes all single quotes and double quotes.
	RemoveAllQuotes *string `field:"optional" json:"removeAllQuotes" yaml:"removeAllQuotes"`
	// If `true` , removes all whitespaces from the value.
	RemoveAllWhitespace *string `field:"optional" json:"removeAllWhitespace" yaml:"removeAllWhitespace"`
	// If `true` , removes all chraracters specified by `CustomCharacters` .
	RemoveCustomCharacters *string `field:"optional" json:"removeCustomCharacters" yaml:"removeCustomCharacters"`
	// If `true` , removes all chraracters specified by `CustomValue` .
	RemoveCustomValue *string `field:"optional" json:"removeCustomValue" yaml:"removeCustomValue"`
	// If `true` , removes the following characters if they occur at the start or end of the value: `.` `!` `,` `?`.
	RemoveLeadingAndTrailingPunctuation *string `field:"optional" json:"removeLeadingAndTrailingPunctuation" yaml:"removeLeadingAndTrailingPunctuation"`
	// If `true` , removes single quotes and double quotes from the beginning and end of the value.
	RemoveLeadingAndTrailingQuotes *string `field:"optional" json:"removeLeadingAndTrailingQuotes" yaml:"removeLeadingAndTrailingQuotes"`
	// If `true` , removes all whitespaces from the beginning and end of the value.
	RemoveLeadingAndTrailingWhitespace *string `field:"optional" json:"removeLeadingAndTrailingWhitespace" yaml:"removeLeadingAndTrailingWhitespace"`
	// If `true` , removes all uppercase and lowercase alphabetic characters (A through Z;
	//
	// a through z).
	RemoveLetters *string `field:"optional" json:"removeLetters" yaml:"removeLetters"`
	// If `true` , removes all numeric characters (0 through 9).
	RemoveNumbers *string `field:"optional" json:"removeNumbers" yaml:"removeNumbers"`
	// If `true` , the source column will be removed after un-nesting that column.
	//
	// (Used with nested column types, such as Map, Struct, or Array.)
	RemoveSourceColumn *string `field:"optional" json:"removeSourceColumn" yaml:"removeSourceColumn"`
	// If `true` , removes all of the following characters: `!
	//
	// " # $ % & ' ( ) * + , - . / : ; < = > ? @ [ \ ] ^ _ ` { | } ~`
	RemoveSpecialCharacters *string `field:"optional" json:"removeSpecialCharacters" yaml:"removeSpecialCharacters"`
	// The columns on the right side of a join.
	RightColumns *string `field:"optional" json:"rightColumns" yaml:"rightColumns"`
	// The number of rows in the sample.
	SampleSize *string `field:"optional" json:"sampleSize" yaml:"sampleSize"`
	// The sampling type to apply to the dataset.
	//
	// Valid values: `FIRST_N` | `LAST_N` | `RANDOM`.
	SampleType *string `field:"optional" json:"sampleType" yaml:"sampleType"`
	// A list of secondary inputs in a UNION transform.
	SecondaryInputs interface{} `field:"optional" json:"secondaryInputs" yaml:"secondaryInputs"`
	// A object value to indicate the second dataset used in a join.
	SecondInput *string `field:"optional" json:"secondInput" yaml:"secondInput"`
	// One or more sheet numbers in the Excel file, which will be included in a dataset.
	SheetIndexes interface{} `field:"optional" json:"sheetIndexes" yaml:"sheetIndexes"`
	// Oone or more named sheets in the Excel file, which will be included in a dataset.
	SheetNames *[]*string `field:"optional" json:"sheetNames" yaml:"sheetNames"`
	// A source column needed for an operation, step, or transform.
	SourceColumn *string `field:"optional" json:"sourceColumn" yaml:"sourceColumn"`
	// A source column needed for an operation, step, or transform.
	SourceColumn1 *string `field:"optional" json:"sourceColumn1" yaml:"sourceColumn1"`
	// A source column needed for an operation, step, or transform.
	SourceColumn2 *string `field:"optional" json:"sourceColumn2" yaml:"sourceColumn2"`
	// A list of source columns needed for an operation, step, or transform.
	SourceColumns *string `field:"optional" json:"sourceColumns" yaml:"sourceColumns"`
	// The index number of the first column used by an operation, step, or transform.
	StartColumnIndex *string `field:"optional" json:"startColumnIndex" yaml:"startColumnIndex"`
	// The starting pattern to locate.
	StartPattern *string `field:"optional" json:"startPattern" yaml:"startPattern"`
	// The starting position to locate.
	StartPosition *string `field:"optional" json:"startPosition" yaml:"startPosition"`
	// The starting value to locate.
	StartValue *string `field:"optional" json:"startValue" yaml:"startValue"`
	// Indicates this operation uses stems and lemmas (base words) for word tokenization.
	StemmingMode *string `field:"optional" json:"stemmingMode" yaml:"stemmingMode"`
	// The total number of transforms in this recipe.
	StepCount *string `field:"optional" json:"stepCount" yaml:"stepCount"`
	// The index ID of a step.
	StepIndex *string `field:"optional" json:"stepIndex" yaml:"stepIndex"`
	// Indicates this operation uses stop words as part of word tokenization.
	StopWordsMode *string `field:"optional" json:"stopWordsMode" yaml:"stopWordsMode"`
	// The resolution strategy to apply in resolving ambiguities.
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
	// The column targeted by this operation.
	TargetColumn *string `field:"optional" json:"targetColumn" yaml:"targetColumn"`
	// The names to give columns altered by this operation.
	TargetColumnNames *string `field:"optional" json:"targetColumnNames" yaml:"targetColumnNames"`
	// The date format to convert to.
	TargetDateFormat *string `field:"optional" json:"targetDateFormat" yaml:"targetDateFormat"`
	// The index number of an object that is targeted by this operation.
	TargetIndex *string `field:"optional" json:"targetIndex" yaml:"targetIndex"`
	// The current timezone that you want to use for dates.
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
	// A regex expression to use when splitting text into terms, also called words or tokens.
	TokenizerPattern *string `field:"optional" json:"tokenizerPattern" yaml:"tokenizerPattern"`
	// A value to use to represent `TRUE` .
	TrueString *string `field:"optional" json:"trueString" yaml:"trueString"`
	// The language that's used in the user-defined function.
	UdfLang *string `field:"optional" json:"udfLang" yaml:"udfLang"`
	// Specifies a unit of time.
	//
	// For example: `MINUTES` ; `SECONDS` ; `HOURS` ; etc.
	Units *string `field:"optional" json:"units" yaml:"units"`
	// Cast columns as rows, so that each value is a different row in a single column.
	UnpivotColumn *string `field:"optional" json:"unpivotColumn" yaml:"unpivotColumn"`
	// The upper boundary for a value.
	UpperBound *string `field:"optional" json:"upperBound" yaml:"upperBound"`
	// Create a new container to hold a dataset.
	UseNewDataFrame *string `field:"optional" json:"useNewDataFrame" yaml:"useNewDataFrame"`
	// A static value that can be used in a comparison, a substitution, or in another context-specific way.
	//
	// A `Value` can be a number, string, or other datatype, depending on the recipe action in which it's used.
	Value *string `field:"optional" json:"value" yaml:"value"`
	// A value that's used by this operation.
	Value1 *string `field:"optional" json:"value1" yaml:"value1"`
	// A value that's used by this operation.
	Value2 *string `field:"optional" json:"value2" yaml:"value2"`
	// The column that is provided as a value that's used by this operation.
	ValueColumn *string `field:"optional" json:"valueColumn" yaml:"valueColumn"`
	// The subset of rows currently available for viewing.
	ViewFrame *string `field:"optional" json:"viewFrame" yaml:"viewFrame"`
}

