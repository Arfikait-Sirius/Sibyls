package sibyls

import (
	"strings"
)

type louise struct{}

var Louise louise

//------------------------
// :[ NAME ]:
//     FnCount
//
// :[ CATEGORY ]:
//     Skill
//------------------------
func ( l louise ) FnCount( base string, target string ) int {
	return strings.Count( base, target )
}

//------------------------
// :[ NAME ]:
//     FnCopy
//
// :[ CATEGORY ]:
//     Skill
//------------------------
func ( l louise ) FnCopy( base string ) string {
	return base
}


//------------------------
// :[ NAME ]:
//     FnReplace
//
// :[ CATEGORY ]:
//     Skill
//------------------------
func ( l louise ) FnReplace( base string, target string, replacement string ) string {
	return strings.Replace( base, target, replacement, 1 )
}

//------------------------
// :[ NAME ]:
//     FnSplit
//
// :[ CATEGORY ]:
//     Skill
//------------------------
func ( l louise ) FnSplit( base string, splitter string, index int ) string {
	list := strings.Split( base, splitter )
	
	if len( list ) > index {
		return ""
	}
	
	return list[index]
}

//------------------------
// :[ NAME ]:
//     FnUpperAll
//
// :[ CATEGORY ]:
//     Skill
//------------------------
func ( l louise ) FnUpperAll( base string ) string {
	return strings.ToUpper( base )
}

//------------------------
// :[ NAME ]:
//     FnLowerAll
//
// :[ CATEGORY ]:
//     Skill
//------------------------
func ( l louise ) FnLowerAll( base string ) string {
	return strings.ToLower( base )
}

//------------------------
// :[ NAME ]:
//     FnUpperFirst
//
// :[ CATEGORY ]:
//     Skill
//------------------------
func ( l louise ) FnUpperFirst( base string ) string {
	upper := strings.ToUpper( base )
	lower := strings.ToLower( base )

	return upper[:1] + lower[1:]
}

