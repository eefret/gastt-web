package lang

//Lang is the language type
type Lang string

//Languages contains all available languages
var Languages = struct {
	//Afrikaans language
	Afrikaans Lang
	//AF language
	AF Lang
	//Albanian language
	Albanian Lang
	//SQ language
	SQ Lang
	//Arabic language
	Arabic Lang
	//AR language
	AR Lang
	//Armenian language
	Armenian Lang
	//HY language
	HY Lang
	//Azerbaijani language
	Azerbaijani Lang
	//AZ language
	AZ Lang
	//Basque language
	Basque Lang
	//EU language
	EU Lang
	//Belarusian language
	Belarusian Lang
	//BE language
	BE Lang
	//Bengali language
	Bengali Lang
	//BN language
	BN Lang
	//Bosnian language
	Bosnian Lang
	//BS language
	BS Lang
	//Bulgarian language
	Bulgarian Lang
	//BG language
	BG Lang
	//Catalan language
	Catalan Lang
	//CA language
	CA Lang
	//Cebuano language
	Cebuano Lang
	//CEB language
	CEB Lang
	//ChineseS language
	ChineseS Lang
	//ZHCN language
	ZHCN Lang
	//ChineseT language
	ChineseT Lang
	//ZHTW language
	ZHTW Lang
	//Croatian language
	Croatian Lang
	//HR language
	HR Lang
	//Czech language
	Czech Lang
	//CS language
	CS Lang
	//Danish language
	Danish Lang
	//DA language
	DA Lang
	//Dutch language
	Dutch Lang
	//NL language
	NL Lang
	//English language
	English Lang
	//EN language
	EN Lang
	//Esperanto language
	Esperanto Lang
	//EO language
	EO Lang
	//Estonian language
	Estonian Lang
	//ET language
	ET Lang
	//Fillipino language
	Fillipino Lang
	//TL language
	TL Lang
	//Finnish language
	Finnish Lang
	//FI language
	FI Lang
	//French language
	French Lang
	//FR language
	FR Lang
	//Galician language
	Galician Lang
	//GL language
	GL Lang
	//Georgian language
	Georgian Lang
	//KA language
	KA Lang
	//German language
	German Lang
	//DE language
	DE Lang
	//Greek language
	Greek Lang
	//EL language
	EL Lang
	//Gujarati language
	Gujarati Lang
	//GU language
	GU Lang
	//Haitian language
	Haitian Lang
	//HT language
	HT Lang
	//Hausa language
	Hausa Lang
	//HA language
	HA Lang
	//Hebrew language
	Hebrew Lang
	//IW language
	IW Lang
	//Hindi language
	Hindi Lang
	//HI language
	HI Lang
	//Hmong language
	Hmong Lang
	//HMN language
	HMN Lang
	//Hungarian language
	Hungarian Lang
	//HU language
	HU Lang
	//Icelandic language
	Icelandic Lang
	//IS language
	IS Lang
	//Igbo language
	Igbo Lang
	//IG language
	IG Lang
	//Indonesian language
	Indonesian Lang
	//ID language
	ID Lang
	//Irish language
	Irish Lang
	//GA language
	GA Lang
	//Italian language
	Italian Lang
	//IT language
	IT Lang
	//Japanese language
	Japanese Lang
	//JA language
	JA Lang
	//Javanese language
	Javanese Lang
	//JW language
	JW Lang
	//Kannada language
	Kannada Lang
	//KN language
	KN Lang
	//Khmer language
	Khmer Lang
	//KM language
	KM Lang
	//Korean language
	Korean Lang
	//KO language
	KO Lang
	//Lao language
	Lao Lang
	//LO language
	LO Lang
	//Latin language
	Latin Lang
	//LA language
	LA Lang
	//Latvian language
	Latvian Lang
	//LV language
	LV Lang
	//Lithuanian language
	Lithuanian Lang
	//LT language
	LT Lang
	//Macedonian language
	Macedonian Lang
	//MK language
	MK Lang
	//Malay language
	Malay Lang
	//MS language
	MS Lang
	//Maltese language
	Maltese Lang
	//MT language
	MT Lang
	//Maori language
	Maori Lang
	//MI language
	MI Lang
	//Marathi language
	Marathi Lang
	//MR language
	MR Lang
	//Mongolian language
	Mongolian Lang
	//MN language
	MN Lang
	//Nepali language
	Nepali Lang
	//NE language
	NE Lang
	//Norwegian language
	Norwegian Lang
	//NO language
	NO Lang
	//Persian language
	Persian Lang
	//FA language
	FA Lang
	//Polish language
	Polish Lang
	//PL language
	PL Lang
	//Portuguese language
	Portuguese Lang
	//PT language
	PT Lang
	//Punjabi language
	Punjabi Lang
	//PA language
	PA Lang
	//Romanian language
	Romanian Lang
	//RO language
	RO Lang
	//Russian language
	Russian Lang
	//RU language
	RU Lang
	//Serbian language
	Serbian Lang
	//SR language
	SR Lang
	//Slovak language
	Slovak Lang
	//SK language
	SK Lang
	//Slovenian language
	Slovenian Lang
	//SL language
	SL Lang
	//Somali language
	Somali Lang
	//SO language
	SO Lang
	//Spanish language
	Spanish Lang
	//ES language
	ES Lang
	//Swahili language
	Swahili Lang
	//SW language
	SW Lang
	//Swedish language
	Swedish Lang
	//SV language
	SV Lang
	//Tamil language
	Tamil Lang
	//TA language
	TA Lang
	//Telugu language
	Telugu Lang
	//TE language
	TE Lang
	//Thai language
	Thai Lang
	//TH language
	TH Lang
	//Turkish language
	Turkish Lang
	//TR language
	TR Lang
	//Ukranian language
	Ukranian Lang
	//UK language
	UK Lang
	//Urdu language
	Urdu Lang
	//UR language
	UR Lang
	//Vietnamese language
	Vietnamese Lang
	//VI language
	VI Lang
	//Weish language
	Weish Lang
	//CY language
	CY Lang
	//Yiddish language
	Yiddish Lang
	//YI language
	YI Lang
	//Yoruba language
	Yoruba Lang
	//YO language
	YO Lang
	//Zulu language
	Zulu Lang
	//ZU language
	ZU Lang
	//AUTO Automatically recognized lang
	AUTO Lang
	//Automatic recognized Lang
	Automatic Lang
}{
	"af",
	"af",
	"sq",
	"sq",
	"ar",
	"ar",
	"hy",
	"hy",
	"az",
	"az",
	"eu",
	"eu",
	"be",
	"be",
	"bn",
	"bn",
	"bs",
	"bs",
	"bg",
	"bg",
	"ca",
	"ca",
	"ceb",
	"ceb",
	"zh-CN",
	"zh-CN",
	"zh-TW",
	"zh-TW",
	"hr",
	"hr",
	"cs",
	"cs",
	"da",
	"da",
	"nl",
	"nl",
	"en",
	"en",
	"eo",
	"eo",
	"et",
	"et",
	"tl",
	"tl",
	"fi",
	"fi",
	"fr",
	"fr",
	"gl",
	"gl",
	"ka",
	"ka",
	"de",
	"de",
	"el",
	"el",
	"gu",
	"gu",
	"ht",
	"ht",
	"ha",
	"ha",
	"iw",
	"iw",
	"hi",
	"hi",
	"hmn",
	"hmn",
	"hu",
	"hu",
	"is",
	"is",
	"ig",
	"ig",
	"id",
	"id",
	"ga",
	"ga",
	"it",
	"it",
	"ja",
	"ja",
	"jw",
	"jw",
	"kn",
	"kn",
	"km",
	"km",
	"ko",
	"ko",
	"lo",
	"lo",
	"la",
	"la",
	"lv",
	"lv",
	"lt",
	"lt",
	"mk",
	"mk",
	"ms",
	"ms",
	"mt",
	"mt",
	"mi",
	"mi",
	"mr",
	"mr",
	"mn",
	"mn",
	"ne",
	"ne",
	"no",
	"no",
	"fa",
	"fa",
	"pl",
	"pl",
	"pt",
	"pt",
	"pa",
	"pa",
	"ro",
	"ro",
	"ru",
	"ru",
	"sr",
	"sr",
	"sk",
	"sk",
	"sl",
	"sl",
	"so",
	"so",
	"es",
	"es",
	"sw",
	"sw",
	"sv",
	"sv",
	"ta",
	"ta",
	"te",
	"te",
	"th",
	"th",
	"tr",
	"tr",
	"uk",
	"uk",
	"ur",
	"ur",
	"vi",
	"vi",
	"cy",
	"cy",
	"yi",
	"yi",
	"yo",
	"yo",
	"zu",
	"zu",
	"auto",
	"auto"}
