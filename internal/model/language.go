package model

// Language represents a programming language.


const (
	LanguageUnknown Language = iota
	LanguageJava
	LanguageGo
	LanguagePython
	LanguageShell
	LanguageRuby
	LanguagePHP
	LanguageJavaScript
	LanguageTypeScript
	LanguageHTML
	LanguageCSS
)

func (l Language) String() string {
	switch l {
	case LanguageJava:
		return "Java"
	case LanguageGo:
		return "Go"
	case LanguagePython:
		return "Python"
	case LanguageShell:
		return "Shell"
	case LanguageRuby:
		return "Ruby"
	case LanguagePHP:
		return "PHP"
	case LanguageJavaScript:
		return "JavaScript"
	case LanguageTypeScript:
		return "TypeScript"
	case LanguageHTML:
		return "HTML"
	case LanguageCSS:
		return "CSS"
	default:
		return "Unknown"
	}
}

type Language int