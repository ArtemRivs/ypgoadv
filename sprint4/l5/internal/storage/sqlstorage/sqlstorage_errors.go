package sqlstorage

import (
	"errors"
)

var ErrExistedURL = errors.New(`existed URL`)
var ErrDeletedURL = errors.New(`deleted URL`)
