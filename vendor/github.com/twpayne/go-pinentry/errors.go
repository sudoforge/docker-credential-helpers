package pinentry

import "errors"

// combineErrors combines all non-nil errors in errs into one. If there are no
// non-nil errors, it returns nil. If there is exactly one non-nil error then it
// returns that error. Otherwise, it returns the non-nil errors combined with
// errors.Join.
func combineErrors(errs ...error) error {
	nonNilErrs := make([]error, 0, len(errs))
	for _, err := range errs {
		if err != nil {
			nonNilErrs = append(nonNilErrs, err)
		}
	}
	switch len(nonNilErrs) {
	case 0:
		return nil
	case 1:
		return nonNilErrs[0]
	default:
		return errors.Join(nonNilErrs...)
	}
}

// combineErrorFunc combines the error pointed to by errp with the result of
// calling f.
func combineErrorFunc(errp *error, f func() error) {
	if err := f(); err != nil {
		*errp = combineErrors(*errp, err)
	}
}
