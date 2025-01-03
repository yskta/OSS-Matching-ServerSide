package model

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

// UserSkill represents a row from 'public.user_skills'.
type UserSkill struct {
	ID        uuid.UUID      `json:"id"`         // id
	UserID    uuid.UUID      `json:"user_id"`    // user_id
	Name      string         `json:"name"`       // name
	Level     sql.NullString `json:"level"`      // level
	CreatedAt sql.NullTime   `json:"created_at"` // created_at
	UpdatedAt sql.NullTime   `json:"updated_at"` // updated_at
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [UserSkill] exists in the database.
func (us *UserSkill) Exists() bool {
	return us._exists
}

// Deleted returns true when the [UserSkill] has been marked for deletion
// from the database.
func (us *UserSkill) Deleted() bool {
	return us._deleted
}

// Insert inserts the [UserSkill] to the database.
func (us *UserSkill) Insert(ctx context.Context, db DB) error {
	switch {
	case us._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case us._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO public.user_skills (` +
		`id, user_id, name, level, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`)`
	// run
	logf(sqlstr, us.ID, us.UserID, us.Name, us.Level, us.CreatedAt, us.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, us.ID, us.UserID, us.Name, us.Level, us.CreatedAt, us.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	us._exists = true
	return nil
}

// Update updates a [UserSkill] in the database.
func (us *UserSkill) Update(ctx context.Context, db DB) error {
	switch {
	case !us._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case us._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.user_skills SET ` +
		`user_id = $1, name = $2, level = $3, created_at = $4, updated_at = $5 ` +
		`WHERE id = $6`
	// run
	logf(sqlstr, us.UserID, us.Name, us.Level, us.CreatedAt, us.UpdatedAt, us.ID)
	if _, err := db.ExecContext(ctx, sqlstr, us.UserID, us.Name, us.Level, us.CreatedAt, us.UpdatedAt, us.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [UserSkill] to the database.
func (us *UserSkill) Save(ctx context.Context, db DB) error {
	if us.Exists() {
		return us.Update(ctx, db)
	}
	return us.Insert(ctx, db)
}

// Upsert performs an upsert for [UserSkill].
func (us *UserSkill) Upsert(ctx context.Context, db DB) error {
	switch {
	case us._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.user_skills (` +
		`id, user_id, name, level, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`user_id = EXCLUDED.user_id, name = EXCLUDED.name, level = EXCLUDED.level, created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at `
	// run
	logf(sqlstr, us.ID, us.UserID, us.Name, us.Level, us.CreatedAt, us.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, us.ID, us.UserID, us.Name, us.Level, us.CreatedAt, us.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	us._exists = true
	return nil
}

// Delete deletes the [UserSkill] from the database.
func (us *UserSkill) Delete(ctx context.Context, db DB) error {
	switch {
	case !us._exists: // doesn't exist
		return nil
	case us._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.user_skills ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, us.ID)
	if _, err := db.ExecContext(ctx, sqlstr, us.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	us._deleted = true
	return nil
}

// UserSkillByID retrieves a row from 'public.user_skills' as a [UserSkill].
//
// Generated from index 'user_skills_pkey'.
func UserSkillByID(ctx context.Context, db DB, id uuid.UUID) (*UserSkill, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, user_id, name, level, created_at, updated_at ` +
		`FROM public.user_skills ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, id)
	us := UserSkill{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&us.ID, &us.UserID, &us.Name, &us.Level, &us.CreatedAt, &us.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &us, nil
}

// User returns the User associated with the [UserSkill]'s (UserID).
//
// Generated from foreign key 'user_skills_user_id_fkey'.
func (us *UserSkill) User(ctx context.Context, db DB) (*User, error) {
	return UserByID(ctx, db, us.UserID)
}