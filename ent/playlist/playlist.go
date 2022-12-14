// Code generated by ent, DO NOT EDIT.

package playlist

const (
	// Label holds the string label denoting the playlist type in the database.
	Label = "playlist"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeUploads holds the string denoting the uploads edge name in mutations.
	EdgeUploads = "uploads"
	// Table holds the table name of the playlist in the database.
	Table = "playlists"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "playlists"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// UploadsTable is the table that holds the uploads relation/edge.
	UploadsTable = "uploads"
	// UploadsInverseTable is the table name for the Upload entity.
	// It exists in this package in order to avoid circular dependency with the "upload" package.
	UploadsInverseTable = "uploads"
	// UploadsColumn is the table column denoting the uploads relation/edge.
	UploadsColumn = "playlist_id"
)

// Columns holds all SQL columns for playlist fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldTitle,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
