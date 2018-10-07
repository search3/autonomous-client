package resourceManager

import (
	"os"
	"net/http"
	ipfs "github.com/ipfs/go-ipfs-api"
	"github.com/space3/autonomous-client/jobmodels"
)


func GetJobResources(jobmodels.JobSpec job) {
	ipfsShell := ipfs.NewShell("localhost:5001")
	err := ipfsShell.Get(job.modelHash, "facenet_ipfs")
	if err != nil {
		panic(err)
	}
}

func UploadJobArtifiacts(JobSubmission submission) {
	ipfsShell := ipfs.NewShell("localhost:5001")
	cid, err := ipfsShell.Add(os.Open("job_submission"))
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", "http://search3upload.herokuapp.com/")
	q := req.URL.Query()
	q.Add("photo_hash", cid)
	q.Add("contact_info", "1 infinite loop")
	q.Add("job_id", submission.)	
	client := http.Client{}
	client.Do(req)
}



