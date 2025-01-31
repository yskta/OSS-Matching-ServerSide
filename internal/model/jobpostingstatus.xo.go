package model

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql/driver"
	"fmt"
)

// JobPostingStatus is the 'job_posting_status' enum type from schema 'public'.
type JobPostingStatus uint16

// JobPostingStatus values.
const (
	// JobPostingStatusDraft is the 'draft' job_posting_status.
	JobPostingStatusDraft JobPostingStatus = 1
	// JobPostingStatusOpen is the 'open' job_posting_status.
	JobPostingStatusOpen JobPostingStatus = 2
	// JobPostingStatusClosed is the 'closed' job_posting_status.
	JobPostingStatusClosed JobPostingStatus = 3
	// JobPostingStatusCancelled is the 'cancelled' job_posting_status.
	JobPostingStatusCancelled JobPostingStatus = 4
)

// String satisfies the [fmt.Stringer] interface.
func (jps JobPostingStatus) String() string {
	switch jps {
	case JobPostingStatusDraft:
		return "draft"
	case JobPostingStatusOpen:
		return "open"
	case JobPostingStatusClosed:
		return "closed"
	case JobPostingStatusCancelled:
		return "cancelled"
	}
	return fmt.Sprintf("JobPostingStatus(%d)", jps)
}

// MarshalText marshals [JobPostingStatus] into text.
func (jps JobPostingStatus) MarshalText() ([]byte, error) {
	return []byte(jps.String()), nil
}

// UnmarshalText unmarshals [JobPostingStatus] from text.
func (jps *JobPostingStatus) UnmarshalText(buf []byte) error {
	switch str := string(buf); str {
	case "draft":
		*jps = JobPostingStatusDraft
	case "open":
		*jps = JobPostingStatusOpen
	case "closed":
		*jps = JobPostingStatusClosed
	case "cancelled":
		*jps = JobPostingStatusCancelled
	default:
		return ErrInvalidJobPostingStatus(str)
	}
	return nil
}

// Value satisfies the [driver.Valuer] interface.
func (jps JobPostingStatus) Value() (driver.Value, error) {
	return jps.String(), nil
}

// Scan satisfies the [sql.Scanner] interface.
func (jps *JobPostingStatus) Scan(v interface{}) error {
	switch x := v.(type) {
	case []byte:
		return jps.UnmarshalText(x)
	case string:
		return jps.UnmarshalText([]byte(x))
	}
	return ErrInvalidJobPostingStatus(fmt.Sprintf("%T", v))
}

// NullJobPostingStatus represents a null 'job_posting_status' enum for schema 'public'.
type NullJobPostingStatus struct {
	JobPostingStatus JobPostingStatus
	// Valid is true if [JobPostingStatus] is not null.
	Valid bool
}

// Value satisfies the [driver.Valuer] interface.
func (njps NullJobPostingStatus) Value() (driver.Value, error) {
	if !njps.Valid {
		return nil, nil
	}
	return njps.JobPostingStatus.Value()
}

// Scan satisfies the [sql.Scanner] interface.
func (njps *NullJobPostingStatus) Scan(v interface{}) error {
	if v == nil {
		njps.JobPostingStatus, njps.Valid = 0, false
		return nil
	}
	err := njps.JobPostingStatus.Scan(v)
	njps.Valid = err == nil
	return err
}

// ErrInvalidJobPostingStatus is the invalid [JobPostingStatus] error.
type ErrInvalidJobPostingStatus string

// Error satisfies the error interface.
func (err ErrInvalidJobPostingStatus) Error() string {
	return fmt.Sprintf("invalid JobPostingStatus(%s)", string(err))
}
