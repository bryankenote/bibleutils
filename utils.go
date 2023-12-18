package utils

func GetTotalChapters() map[string]int {
	chapters := make(map[string]int)

	chapters["Genesis"] = 50
	chapters["Exodus"] = 40
	chapters["Leviticus"] = 27
	chapters["Numbers"] = 36
	chapters["Deuteronomy"] = 34
	chapters["Joshua"] = 24
	chapters["Judges"] = 21
	chapters["Ruth"] = 4
	chapters["1 Samuel"] = 31
	chapters["2 Samuel"] = 24
	chapters["1 Kings"] = 22
	chapters["2 Kings"] = 25
	chapters["1 Chronicles"] = 29
	chapters["2 Chronicles"] = 36
	chapters["Ezra"] = 10
	chapters["Nehemiah"] = 13
	chapters["Esther"] = 10
	chapters["Job"] = 42
	chapters["Psalms"] = 150
	chapters["Proverbs"] = 31
	chapters["Ecclesiastes"] = 12
	chapters["Song of Solomon"] = 8
	chapters["Isaiah"] = 66
	chapters["Jeremiah"] = 52
	chapters["Lamentations"] = 5
	chapters["Ezekiel"] = 48
	chapters["Daniel"] = 12
	chapters["Hosea"] = 14
	chapters["Joel"] = 6
	chapters["Amos"] = 9
	chapters["Obadiah"] = 1
	chapters["Jonah"] = 4
	chapters["Micah"] = 7
	chapters["Nahum"] = 3
	chapters["Habakkuk"] = 3
	chapters["Zephaniah"] = 3
	chapters["Haggai"] = 2
	chapters["Zechariah"] = 14
	chapters["Malachi"] = 4

	chapters["Matthew"] = 28
	chapters["Mark"] = 16
	chapters["Luke"] = 24
	chapters["John"] = 21
	chapters["Acts"] = 28
	chapters["Romans"] = 16
	chapters["1 Corinthians"] = 16
	chapters["2 Corinthians"] = 13
	chapters["Galatians"] = 6
	chapters["Ephesians"] = 6
	chapters["Philippians"] = 4
	chapters["Colossians"] = 4
	chapters["1 Thessalonians"] = 5
	chapters["2 Thessalonians"] = 3
	chapters["1 Timothy"] = 6
	chapters["2 Timothy"] = 4
	chapters["Titus"] = 3
	chapters["Philemon"] = 1
	chapters["Hebrews"] = 13
	chapters["James"] = 5
	chapters["1 Peter"] = 5
	chapters["2 Peter"] = 3
	chapters["1 John"] = 5
	chapters["2 John"] = 1
	chapters["3 John"] = 1
	chapters["Jude"] = 1
	chapters["Revelation"] = 22

	return chapters
}

func GetBookNames() []string {
	return []string{
		"Genesis",
		"Exodus",
		"Leviticus",
		"Numbers",
		"Deuteronomy",
		"Joshua",
		"Judges",
		"Ruth",
		"1 Samuel",
		"2 Samuel",
		"1 Kings",
		"2 Kings",
		"1 Chronicles",
		"2 Chronicles",
		"Ezra",
		"Nehemiah",
		"Esther",
		"Job",
		"Psalms",
		"Proverbs",
		"Ecclesiastes",
		"Song of Solomon",
		"Isaiah",
		"Jeremiah",
		"Lamentations",
		"Ezekiel",
		"Daniel",
		"Hosea",
		"Joel",
		"Amos",
		"Obadiah",
		"Jonah",
		"Micah",
		"Nahum",
		"Habakkuk",
		"Zephaniah",
		"Haggai",
		"Zechariah",
		"Malachi",

		"Matthew",
		"Mark",
		"Luke",
		"John",
		"Acts",
		"Romans",
		"1 Corinthians",
		"2 Corinthians",
		"Galatians",
		"Ephesians",
		"Philippians",
		"Colossians",
		"1 Thessalonians",
		"2 Thessalonians",
		"1 Timothy",
		"2 Timothy",
		"Titus",
		"Philemon",
		"Hebrews",
		"James",
		"1 Peter",
		"2 Peter",
		"1 John",
		"2 John",
		"3 John",
		"Jude",
		"Revelation",
	}
}
