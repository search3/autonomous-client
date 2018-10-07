package jobspec

type JobSpec struct {
	jobHash         string
	modelHash       string
	targetEmbedding string
}

type JobSubmission struct {
	photoPath string
	jobHash   string
}
