package herokuApi

type TeamCollaborator struct {
	permissions []string
	email       string
	silent      bool
}
