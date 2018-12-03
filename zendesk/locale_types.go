package zendesk

// DO NOT EDIT
// This file is generated by script/create_locale_types.rb
//
// 64 locales are supported

const (
	// LocaleENUS English
	LocaleENUS = 1
	// LocaleES Spanish - español
	LocaleES = 2
	// LocaleDE German - Deutsch
	LocaleDE = 8
	// LocaleZHTW Traditional Chinese - 繁體中文
	LocaleZHTW = 9
	// LocaleZHCN Simplified Chinese - 简体中文
	LocaleZHCN = 10
	// LocalePL Polish - polski
	LocalePL = 13
	// LocaleFR French - français
	LocaleFR = 16
	// LocalePTBR Brazilian Portuguese - português (Brasil)
	LocalePTBR = 19
	// LocaleIT Italian - italiano
	LocaleIT = 22
	// LocaleRO Romanian - română
	LocaleRO = 23
	// LocaleIS Icelandic - íslenska
	LocaleIS = 24
	// LocaleVI Vietnamese - Tiếng Việt
	LocaleVI = 26
	// LocaleRU Russian - русский
	LocaleRU = 27
	// LocaleHE Hebrew - עברית
	LocaleHE = 30
	// LocaleNO Norwegian - norsk
	LocaleNO = 34
	// LocaleFIL Filipino
	LocaleFIL = 47
	// LocaleAR Arabic - العربية
	LocaleAR = 66
	// LocaleJA Japanese - 日本語
	LocaleJA = 67
	// LocaleKO Korean - 한국어
	LocaleKO = 69
	// LocaleSL Slovenian - slovenščina
	LocaleSL = 72
	// LocaleHR Croatian - hrvatski
	LocaleHR = 74
	// LocaleID Indonesian - Indonesia
	LocaleID = 77
	// LocaleCS Czech - čeština
	LocaleCS = 78
	// LocaleTH Thai - ไทย
	LocaleTH = 81
	// LocaleFI Finnish - suomi
	LocaleFI = 84
	// LocaleTR Turkish - Türkçe
	LocaleTR = 88
	// LocaleSV Swedish - svenska
	LocaleSV = 92
	// LocaleEL Greek - Ελληνικά
	LocaleEL = 93
	// LocaleBG Bulgarian - български
	LocaleBG = 94
	// LocaleET Estonian - eesti
	LocaleET = 101
	// LocaleDA Danish - dansk
	LocaleDA = 1000
	// LocaleSK Slovak - slovenčina
	LocaleSK = 1003
	// LocaleNL Dutch - Nederlands
	LocaleNL = 1005
	// LocaleHU Hungarian - magyar
	LocaleHU = 1009
	// LocalePT Portuguese - português
	LocalePT = 1011
	// LocaleFA Persian - فارسی
	LocaleFA = 1016
	// LocaleCA Catalan - català
	LocaleCA = 1075
	// LocaleLT Lithuanian - lietuvių
	LocaleLT = 1092
	// LocaleLV Latvian - latviešu
	LocaleLV = 1101
	// LocaleSR Serbian - српски
	LocaleSR = 1150
	// LocaleUK Ukrainian - українська
	LocaleUK = 1173
	// LocaleENGB British English
	LocaleENGB = 1176
	// LocaleENCA Canadian English
	LocaleENCA = 1181
	// LocaleESES European Spanish - español de España
	LocaleESES = 1186
	// LocaleFRCA Canadian French - français canadien
	LocaleFRCA = 1187
	// LocaleES419 Latin American Spanish - español latinoamericano
	LocaleES419 = 1194
	// LocaleENAU Australian English
	LocaleENAU = 1277
	// LocaleENIN English (India)
	LocaleENIN = 1278
	// LocaleENIE English (Ireland)
	LocaleENIE = 1279
	// LocaleENSG English (Singapore)
	LocaleENSG = 1281
	// LocaleENNZ English (New Zealand)
	LocaleENNZ = 1288
	// LocaleENZA English (South Africa)
	LocaleENZA = 1289
	// LocaleFRBE French (Belgium) - français (Belgique)
	LocaleFRBE = 1291
	// LocaleFRCH Swiss French - français suisse
	LocaleFRCH = 1292
	// LocaleNLBE Flemish - Nederlands (België)
	LocaleNLBE = 1293
	// LocaleDEAT Austrian German - Österreichisches Deutsch
	LocaleDEAT = 1294
	// LocaleDECH Swiss High German - Schweizer Hochdeutsch
	LocaleDECH = 1295
	// LocaleSRME Serbian (Montenegro) - srpski (Crna Gora)
	LocaleSRME = 1298
	// LocaleHI Hindi - हिन्दी
	LocaleHI = 1303
	// LocaleMS Malay - Bahasa Melayu
	LocaleMS = 1307
	// LocaleENBE English (Belgium)
	LocaleENBE = 1350
	// LocaleESMX Mexican Spanish - español de México
	LocaleESMX = 1364
	// LocaleFRFR French (France) - français (France)
	LocaleFRFR = 1365
	// LocaleENPH English (Philippines)
	LocaleENPH = 1392
)

var localeTypeText = map[int]string{
	// LocaleENUS English
	LocaleENUS: "en-US",
	// LocaleES Spanish - español
	LocaleES: "es",
	// LocaleDE German - Deutsch
	LocaleDE: "de",
	// LocaleZHTW Traditional Chinese - 繁體中文
	LocaleZHTW: "zh-tw",
	// LocaleZHCN Simplified Chinese - 简体中文
	LocaleZHCN: "zh-cn",
	// LocalePL Polish - polski
	LocalePL: "pl",
	// LocaleFR French - français
	LocaleFR: "fr",
	// LocalePTBR Brazilian Portuguese - português (Brasil)
	LocalePTBR: "pt-br",
	// LocaleIT Italian - italiano
	LocaleIT: "it",
	// LocaleRO Romanian - română
	LocaleRO: "ro",
	// LocaleIS Icelandic - íslenska
	LocaleIS: "is",
	// LocaleVI Vietnamese - Tiếng Việt
	LocaleVI: "vi",
	// LocaleRU Russian - русский
	LocaleRU: "ru",
	// LocaleHE Hebrew - עברית
	LocaleHE: "he",
	// LocaleNO Norwegian - norsk
	LocaleNO: "no",
	// LocaleFIL Filipino
	LocaleFIL: "fil",
	// LocaleAR Arabic - العربية
	LocaleAR: "ar",
	// LocaleJA Japanese - 日本語
	LocaleJA: "ja",
	// LocaleKO Korean - 한국어
	LocaleKO: "ko",
	// LocaleSL Slovenian - slovenščina
	LocaleSL: "sl",
	// LocaleHR Croatian - hrvatski
	LocaleHR: "hr",
	// LocaleID Indonesian - Indonesia
	LocaleID: "id",
	// LocaleCS Czech - čeština
	LocaleCS: "cs",
	// LocaleTH Thai - ไทย
	LocaleTH: "th",
	// LocaleFI Finnish - suomi
	LocaleFI: "fi",
	// LocaleTR Turkish - Türkçe
	LocaleTR: "tr",
	// LocaleSV Swedish - svenska
	LocaleSV: "sv",
	// LocaleEL Greek - Ελληνικά
	LocaleEL: "el",
	// LocaleBG Bulgarian - български
	LocaleBG: "bg",
	// LocaleET Estonian - eesti
	LocaleET: "et",
	// LocaleDA Danish - dansk
	LocaleDA: "da",
	// LocaleSK Slovak - slovenčina
	LocaleSK: "sk",
	// LocaleNL Dutch - Nederlands
	LocaleNL: "nl",
	// LocaleHU Hungarian - magyar
	LocaleHU: "hu",
	// LocalePT Portuguese - português
	LocalePT: "pt",
	// LocaleFA Persian - فارسی
	LocaleFA: "fa",
	// LocaleCA Catalan - català
	LocaleCA: "ca",
	// LocaleLT Lithuanian - lietuvių
	LocaleLT: "lt",
	// LocaleLV Latvian - latviešu
	LocaleLV: "lv",
	// LocaleSR Serbian - српски
	LocaleSR: "sr",
	// LocaleUK Ukrainian - українська
	LocaleUK: "uk",
	// LocaleENGB British English
	LocaleENGB: "en-gb",
	// LocaleENCA Canadian English
	LocaleENCA: "en-ca",
	// LocaleESES European Spanish - español de España
	LocaleESES: "es-es",
	// LocaleFRCA Canadian French - français canadien
	LocaleFRCA: "fr-ca",
	// LocaleES419 Latin American Spanish - español latinoamericano
	LocaleES419: "es-419",
	// LocaleENAU Australian English
	LocaleENAU: "en-au",
	// LocaleENIN English (India)
	LocaleENIN: "en-in",
	// LocaleENIE English (Ireland)
	LocaleENIE: "en-ie",
	// LocaleENSG English (Singapore)
	LocaleENSG: "en-sg",
	// LocaleENNZ English (New Zealand)
	LocaleENNZ: "en-nz",
	// LocaleENZA English (South Africa)
	LocaleENZA: "en-za",
	// LocaleFRBE French (Belgium) - français (Belgique)
	LocaleFRBE: "fr-be",
	// LocaleFRCH Swiss French - français suisse
	LocaleFRCH: "fr-ch",
	// LocaleNLBE Flemish - Nederlands (België)
	LocaleNLBE: "nl-be",
	// LocaleDEAT Austrian German - Österreichisches Deutsch
	LocaleDEAT: "de-at",
	// LocaleDECH Swiss High German - Schweizer Hochdeutsch
	LocaleDECH: "de-ch",
	// LocaleSRME Serbian (Montenegro) - srpski (Crna Gora)
	LocaleSRME: "sr-me",
	// LocaleHI Hindi - हिन्दी
	LocaleHI: "hi",
	// LocaleMS Malay - Bahasa Melayu
	LocaleMS: "ms",
	// LocaleENBE English (Belgium)
	LocaleENBE: "en-be",
	// LocaleESMX Mexican Spanish - español de México
	LocaleESMX: "es-mx",
	// LocaleFRFR French (France) - français (France)
	LocaleFRFR: "fr-fr",
	// LocaleENPH English (Philippines)
	LocaleENPH: "en-ph",
}

// LocaleTypeText returns locale type text
func LocaleTypeText(loc int) string {
	return localeTypeText[loc]
}
