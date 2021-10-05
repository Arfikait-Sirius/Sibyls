package sibyls

import (
	"strings"
)

type louise struct{}

var Louise louise

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
