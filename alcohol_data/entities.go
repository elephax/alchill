package alcohol_data

//easyjson:json
type Spirit struct {
	/* Unique ID in MongoDB */
	ID			string
	/* True title of spirit */
	Name		string
	/* Bar's array */
	Bars		[]string
	/* Spirit's parent categories */
	Categories	[]Spirit
	/* Root category or not (flag) */
	IsCategory	bool
	/* Spirit's preferred units */
	PreferredUnits []PreferredUnit
	/* Image URL for spirit from origin URL */
	ImageURL	string
	/* Product URL (source) */
	ProductURL	string
	/* Strength of alcohol */
	Volume		float32
	/* Bottle capacity */
	Value		float32
	/* Caloricity (100ml) */
	Caloricity	float32
	/* Description */
	Description	string
}

//easyjson:json
type Spirits []Spirit

//easyjson:json
type SpiritWOCategory struct {
	/* Unique ID in MongoDB */
	ID			string
	/* True title of spirit */
	Name		string
	/* Bar's array */
	Bars		[]string
	/* Spirit's parent categories */
	Categories	[]string
	/* Root category or not (flag) */
	IsCategory	bool
	/* Spirit's preferred units */
	PreferredUnits []PreferredUnit
	/* Image URL for spirit from origin URL */
	ImageURL	string
	/* Product URL (source) */
	ProductURL	string
	/* Strength of alcohol */
	Volume		float32
	/* Bottle capacity */
	Value		float32
	/* Caloricity (100ml) */
	Caloricity	float32
	/* Description */
	Description	string
}

//easyjson:json
type SpiritsWOCategory []SpiritWOCategory

//easyjson:json
type PreferredUnit struct {
	/*  */
	Name		string
	/* How much */
	Value     	float32
	/* "ML", "peaces", etc */
	UOM        	string
}
